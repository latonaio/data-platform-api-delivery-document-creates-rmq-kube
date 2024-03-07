package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-delivery-document-creates-rmq-kube/sub_func_complementer"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var partner *[]dpfm_api_output_formatter.Partner
	var address *[]dpfm_api_output_formatter.Address
	var itemPicking *[]dpfm_api_output_formatter.ItemPicking

	handleAccepter := referenceDocumentHandle(
		input,
		subfuncSDC,
		accepter,
	)

	var calculateDeliveryDocumentQueryGets *sub_func_complementer.CalculateDeliveryDocumentQueryGets
	var deliveryDocumentIssuedID int

	if len(handleAccepter) > 0 {
		calculateDeliveryDocumentQueryGets = c.CalculateDeliveryDocument(errs)

		if calculateDeliveryDocumentQueryGets == nil {
			err := xerrors.Errorf("calculateDeliveryDocumentQueryGets is nil")
			*errs = append(*errs, err)
			return nil
		}

		deliveryDocumentIssuedID = calculateDeliveryDocumentQueryGets.DeliveryDocumentLatestNumber + 1

		subfuncSDCHeader := c.getOrdersHeader(
			input,
			errs,
			deliveryDocumentIssuedID,
		)

		if len(*subfuncSDCHeader) > 0 {
			subfuncSDC.Message.Header = &(*subfuncSDCHeader)[0]
		}

		subfuncSDCItem := c.getOrdersItem(
			input,
			errs,
			deliveryDocumentIssuedID,
		)

		if len(*subfuncSDCItem) > 0 {
			subfuncSDC.Message.Item = subfuncSDCItem
		}

		subfuncSDCPartner := c.getOrdersPartner(
			input,
			errs,
			deliveryDocumentIssuedID,
		)

		if len(*subfuncSDCPartner) > 0 {
			subfuncSDC.Message.Partner = subfuncSDCPartner
		}

	}

	for _, fn := range handleAccepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Item":
			item = c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Partner":
			partner = c.partnerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Address":
			address = c.addressCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemPicking":
			itemPicking = c.itemPickingCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:

		}
	}

	if calculateDeliveryDocumentQueryGets != nil {
		err := c.UpdateLatestNumber(errs, deliveryDocumentIssuedID)
		if err != nil {
			*errs = append(*errs, err)
			return nil
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:      header,
		Item:        item,
		Partner:     partner,
		Address:     address,
		ItemPicking: itemPicking,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var partner *[]dpfm_api_output_formatter.Partner
	var address *[]dpfm_api_output_formatter.Address
	//var ItemPicking *[]dpfm_api_output_formatter.ItemPicking
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Item":
			item = c.itemUpdateSql(mtx, input, output, errs, log)
		case "Partner":
			partner = c.partnerUpdateSql(mtx, input, output, errs, log)
		case "Address":
			address = c.addressUpdateSql(mtx, input, output, errs, log)
		//case "ItemPicking":
		//	itemPicking = c.itemPickingUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:  header,
		Item:    item,
		Partner: partner,
		Address: address,
		//ItemPicking: itemPicking,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_delivery_document_header_dataの更新
	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "DeliveryDocumentHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		*errs = append(*errs, err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_delivery_document_item_dataの更新
	for _, itemData := range *subfuncSDC.Message.Item {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "DeliveryDocumentItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_delivery_document_partner_dataの更新
	for _, partnerData := range *subfuncSDC.Message.Partner {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "DeliveryDocumentPartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Partner Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_delivery_document_address_dataの更新
	for _, addressData := range *subfuncSDC.Message.Address {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "DeliveryDocumentAddress", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Address Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemPickingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPicking {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_delivery_document_item_picking_dataの更新
	for _, itemPickingData := range *subfuncSDC.Message.ItemPicking {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemPickingData, "function": "DeliveryDocumentItemPicking", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ItemPicking Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPickingCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "DeliveryDocumentHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	req := make([]dpfm_api_processing_formatter.ItemUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		itemData := *dpfm_api_processing_formatter.ConvertToItemUpdates(header, item)

		if itemIsUpdate(&itemData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "DeliveryDocumentItem", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Item Data cannot insert"
				return nil
			}
		}
		req = append(req, itemData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	req := make([]dpfm_api_processing_formatter.PartnerUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, partner := range header.Partner {
		partnerData := *dpfm_api_processing_formatter.ConvertToPartnerUpdates(header, partner)

		if partnerIsUpdate(&partnerData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "DeliveryDocumentPartner", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Partner Data cannot insert"
				return nil
			}
		}
		req = append(req, partnerData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	req := make([]dpfm_api_processing_formatter.AddressUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, address := range header.Address {
		addressData := *dpfm_api_processing_formatter.ConvertToAddressUpdates(header, address)

		if addressIsUpdate(&addressData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "DeliveryDocumentAddress", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Address Data cannot insert"
				return nil
			}
		}
		req = append(req, addressData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemPickingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPicking {
	req := make([]dpfm_api_processing_formatter.ItemPickingUpdates, 0)
	//sessionID := input.RuntimeSessionID

	//header := input.Header
	//for _, itemPicking := range header.ItemPicking {
	//	itemPickingData := *dpfm_api_processing_formatter.ConvertToItemPickingUpdates(header, itemPicking)
	//
	//	if itemPickingIsUpdate(&itemPickingData) {
	//		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemPickingData, "function": "DeliveryDocumentItemPicking", "runtime_session_id": sessionID})
	//		if err != nil {
	//			err = xerrors.Errorf("rmq error: %w", err)
	//			*errs = append(*errs, err)
	//			return nil
	//		}
	//		res.Success()
	//		if !checkResult(res) {
	//			output.SQLUpdateResult = getBoolPtr(false)
	//			output.SQLUpdateError = "ItemPicking Data cannot insert"
	//			return nil
	//		}
	//	}
	//	req = append(req, itemPickingData)
	//}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPickingUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	deliveryDocument := header.DeliveryDocument

	return !(deliveryDocument == 0)
}

func itemIsUpdate(item *dpfm_api_processing_formatter.ItemUpdates) bool {
	deliveryDocument := item.DeliveryDocument
	deliveryDocumentItem := item.DeliveryDocumentItem

	return !(deliveryDocument == 0 || deliveryDocumentItem == 0)
}

func partnerIsUpdate(partner *dpfm_api_processing_formatter.PartnerUpdates) bool {
	deliveryDocument := partner.DeliveryDocument
	partnerFunction := partner.PartnerFunction
	businessPartner := partner.BusinessPartner

	return !(deliveryDocument == 0 || partnerFunction == "" || businessPartner == 0)
}

func addressIsUpdate(address *dpfm_api_processing_formatter.AddressUpdates) bool {
	deliveryDocument := address.DeliveryDocument
	addressID := address.AddressID

	return !(deliveryDocument == 0 || addressID == 0)
}

func itemPickingIsUpdate(itemPicking *dpfm_api_processing_formatter.ItemPickingUpdates) bool {
	deliveryDocument := itemPicking.DeliveryDocument
	DeliveryDocumentItem := itemPicking.DeliveryDocumentItem
	//DeliveryDocumentItemPickingID := itemPicking.DeliveryDocumentItemPickingID

	return !(deliveryDocument == 0 || DeliveryDocumentItem == 0)
}

func referenceDocumentHandle(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
) []string {
	var handleAccepter []string

	if input.InputParameters.ReferenceDocument != nil {
		handleAccepter = append(handleAccepter, "Header")
		handleAccepter = append(handleAccepter, "Item")
		handleAccepter = append(handleAccepter, "Partner")
	} else {
		handleAccepter = accepter
	}

	return handleAccepter
}

func (c *DPFMAPICaller) getOrdersHeader(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	deliveryDocumentIssuedID int,
) *[]sub_func_complementer.Header {
	orderID := input.InputParameters.ReferenceDocument
	orderItem := 1

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem) = (?, ?);`, orderID, orderItem,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeaderFromOrders(rows, deliveryDocumentIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) getOrdersItem(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	deliveryDocumentIssuedID int,
) *[]sub_func_complementer.Item {
	orderID := input.InputParameters.ReferenceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE OrderID = ?;`, orderID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemFromOrders(rows, deliveryDocumentIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) getOrdersPartner(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	deliveryDocumentIssuedID int,
) *[]sub_func_complementer.Partner {
	orderID := input.InputParameters.ReferenceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data
		WHERE OrderID = ?;`, orderID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartnerFromOrders(rows, deliveryDocumentIssuedID)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) CalculateDeliveryDocument(
	errs *[]error,
) *sub_func_complementer.CalculateDeliveryDocumentQueryGets {
	pm := &sub_func_complementer.CalculateDeliveryDocumentQueryGets{}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, "DELIVERY_DOCUMENT", "DeliveryDocument",
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				*errs = append(*errs, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。"))
				return nil
			} else {
				break
			}
		}
		err = rows.Scan(
			&pm.NumberRangeID,
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.DeliveryDocumentLatestNumber,
		)
		if err != nil {
			*errs = append(*errs, err)
			return nil
		}
	}

	return pm
}

func (c *DPFMAPICaller) UpdateLatestNumber(
	errs *[]error,
	deliveryDocumentIssuedID int,
) error {
	//rows, err := c.db.Query(
	//	`SELECT *
	//	FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
	//	WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, "DELIVERY_DOCUMENT", "DeliveryDocument",
	//)

	_, err := c.db.Exec(`
			UPDATE data_platform_number_range_latest_number_data SET LatestNumber=(?)
			WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`,
		deliveryDocumentIssuedID,
		"DELIVERY_DOCUMENT",
		"DeliveryDocument",
	)
	if err != nil {
		return xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルの更新に失敗しました。")
	}

	return nil
}
