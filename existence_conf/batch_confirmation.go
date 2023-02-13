package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) itemBatchExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		bpID, product, plant, batch, err := getItemBatchExistenceConfKey(mapper, &item, exconfErrMsg)
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
			res, err := c.batchExistenceConfRequest(bpID, product, plant, batch, queueName, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) batchExistenceConfRequest(bpID int, product string, plant string, batch string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Product":         product,
		"Plant":           plant,
		"Batch":           batch,
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
	req.BatchReturn.BusinessPartner = bpID
	req.BatchReturn.Product = product
	req.BatchReturn.Plant = plant
	req.BatchReturn.Batch = batch

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getItemBatchExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, string, string, string, error) {
	var bpID int
	var product string
	var plant string
	var batch string
	var err error

	switch mapper.Field {
	case "DeliverToPlantBatch":
		if item.DeliverToParty == nil ||
			item.Product == nil ||
			item.DeliverToPlant == nil ||
			item.DeliverToPlantBatch == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, "", "", "", err
		}
		if item.Product != nil ||
			item.DeliverToPlant != nil ||
			item.DeliverToPlantBatch != nil {
			if len(*item.Product) == 0 ||
				len(*item.DeliverToPlant) == 0 ||
				len(*item.DeliverToPlantBatch) == 0 {
				err = xerrors.Errorf("cannot specify null keys")
				return 0, "", "", "", err
			}
		}
		bpID = *item.DeliverToParty
		product = *item.Product
		plant = *item.DeliverToPlant
		batch = *item.DeliverToPlantBatch

	case "DeliverFromPlantBatch":
		if item.DeliverFromParty == nil ||
			item.Product == nil ||
			item.DeliverFromPlant == nil ||
			item.DeliverFromPlantBatch == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, "", "", "", err
		}
		if item.Product != nil ||
			item.DeliverFromPlant != nil ||
			item.DeliverFromPlantBatch != nil {
			if len(*item.Product) == 0 ||
				len(*item.DeliverFromPlant) == 0 ||
				len(*item.DeliverFromPlantBatch) == 0 {
				err = xerrors.Errorf("cannot specify null keys")
				return 0, "", "", "", err
			}
		}
		bpID = *item.DeliverFromParty
		product = *item.Product
		plant = *item.DeliverFromPlant
		batch = *item.DeliverFromPlantBatch

	case "StockConfirmationPlantBatch":
		if item.StockConfirmationBusinessPartner == nil ||
			item.Product == nil ||
			item.StockConfirmationPlant == nil ||
			item.StockConfirmationPlantBatch == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, "", "", "", err
		}
		if item.Product != nil ||
			item.StockConfirmationPlant != nil ||
			item.StockConfirmationPlantBatch != nil {
			if len(*item.Product) == 0 ||
				len(*item.StockConfirmationPlant) == 0 ||
				len(*item.StockConfirmationPlantBatch) == 0 {
				err = xerrors.Errorf("cannot specify null keys")
				return 0, "", "", "", err
			}
		}
		bpID = *item.StockConfirmationBusinessPartner
		product = *item.Product
		plant = *item.StockConfirmationPlant
		batch = *item.StockConfirmationPlantBatch

	case "ProductionPlantBatch":
		if item.ProductionPlantBusinessPartner == nil ||
			item.Product == nil ||
			item.ProductionPlant == nil ||
			item.ProductionPlantBatch == nil {
			err = xerrors.Errorf("cannot specify null keys")
			return 0, "", "", "", err
		}
		if item.Product != nil ||
			item.ProductionPlant != nil ||
			item.ProductionPlantBatch != nil {
			if len(*item.Product) == 0 ||
				len(*item.ProductionPlant) == 0 ||
				len(*item.ProductionPlantBatch) == 0 {
				err = xerrors.Errorf("cannot specify null keys")
				return 0, "", "", "", err
			}
		}
		bpID = *item.ProductionPlantBusinessPartner
		product = *item.Product
		plant = *item.ProductionPlant
		batch = *item.ProductionPlantBatch
	}

	return bpID, product, plant, batch, nil
}
