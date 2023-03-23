package sub_func_complementer

import (
	"context"
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-delivery-document-creates-rmq-kube/config"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SubFuncComplementer struct {
	ctx context.Context
	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
	db  *database.Mysql
}

func NewSubFuncComplementer(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *SubFuncComplementer {
	return &SubFuncComplementer{
		ctx: ctx,
		c:   c,
		rmq: rmq,
		db:  db,
	}
}

func (c *SubFuncComplementer) ComplementHeader(input *dpfm_api_input_reader.SDC, subfuncSDC *SDC, l *logger.Logger) error {
	s := &SDC{}
	numRange, err := c.ComplementDeliveryDocument(input, l)
	if err != nil {
		return xerrors.Errorf("complement deliveryDocument error: %w", err)
	}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Headers"], input)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}

	err = c.IncrementLatestNumber(numRange, l)
	if err != nil {
		return xerrors.Errorf("increment latest number error: %w")
	}
	subfuncSDC.SubfuncResult = s.SubfuncResult
	subfuncSDC.SubfuncError = s.SubfuncError

	subfuncSDC.Message.Header = s.Message.Header
	return nil
}

func (c *SubFuncComplementer) ComplementItem(input *dpfm_api_input_reader.SDC, subfuncSDC *SDC, l *logger.Logger) error {
	s := &SDC{}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Items"], input)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(s.Message)
	msg := &Message{}
	err = json.Unmarshal(b, msg)
	if err != nil {
		return err
	}
	subfuncSDC.SubfuncResult = s.SubfuncResult
	subfuncSDC.SubfuncError = s.SubfuncError

	subfuncSDC.Message.Item = msg.Item
	subfuncSDC.Message.Partner = msg.Partner
	subfuncSDC.Message.Address = msg.Address

	return err
}

func getBoolPtr(b bool) *bool {
	return &b
}

func (c *SubFuncComplementer) ComplementDeliveryDocument(input *dpfm_api_input_reader.SDC, l *logger.Logger) (*NumberRange, error) {
	rows, err := c.db.Query(
		`SELECT NumberRangeID, ServiceLabel, FieldNameWithNumberRange, LatestNumber
			FROM DataPlatformCommonSettingsMysqlKube.data_platform_number_range_latest_number_data
			WHERE (ServiceLabel, FieldNameWithNumberRange) = ( (?, 'DeliveryDocument') );`, input.ServiceLabel,
	)
	if err != nil {
		return nil, xerrors.Errorf("DB Query error: %w", err)
	}
	nr := NumberRange{}
	if !rows.Next() {
		return nil, xerrors.Errorf("number range does not exist")
	}
	err = rows.Scan(
		&nr.NumberRangeID,
		&nr.ServiceLabel,
		&nr.FieldNameWithNumberRange,
		&nr.LatestNumber,
	)
	if err != nil {
		return nil, xerrors.Errorf("DB Scan error: %w", err)
	}
	nr.LatestNumber++
	input.Header.DeliveryDocument = nr.LatestNumber
	return &nr, nil
}

func (c *SubFuncComplementer) IncrementLatestNumber(nr *NumberRange, l *logger.Logger) error {
	_, err := c.db.Query(
		`UPDATE DataPlatformCommonSettingsMysqlKube.data_platform_number_range_latest_number_data
			SET LatestNumber = ?
			WHERE  (NumberRangeID, ServiceLabel, FieldNameWithNumberRange) = ((?,?,?));`, nr.LatestNumber, nr.NumberRangeID, nr.ServiceLabel, nr.FieldNameWithNumberRange,
	)
	if err != nil {
		return xerrors.Errorf("DB Query error: %w", err)
	}

	return nil
}
