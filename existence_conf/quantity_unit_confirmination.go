package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) quantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		quantityUnit, err := getQuantityUnitExistenceConfKey(mapper, &item, exconfErrMsg)
		if err != nil {
			*errs = append(*errs, err)
			return
		}
		queueName, err := getQueueName(mapper)
		if err != nil {
			*errs = append(*errs, err)
			return
		}
		wg2.Add(1)
		exReqTimes++
		go func() {
			res, err := c.quantityUnitExistenceConfRequest(quantityUnit, queueName, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) quantityUnitExistenceConfRequest(quantityUnit string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"QuantityUnit": quantityUnit,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.QuantityUnitReturn.QuantityUnit = quantityUnit

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getQuantityUnitExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (string, error) {
	var quantityUnit string
	var err error

	switch mapper.Field {
	case "ItemWeightUnit":
		if item.ItemWeightUnit == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return "", err
		}
		quantityUnit = *item.ItemWeightUnit
	}

	return quantityUnit, nil
}
