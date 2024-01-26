package existence_conf

//func (c *ExistenceConf) itemBatchExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		bpID, product, plant, batch := getItemBatchExistenceConfKey(mapper, &item, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(bpID) || isZero(product) || isZero(plant) || isZero(batch) {
//				wg2.Done()
//				return
//			}
//			res, err := c.batchExistenceConfRequest(bpID, product, plant, batch, queueName, input, existenceMap, mtx, log)
//			if err != nil {
//				mtx.Lock()
//				*errs = append(*errs, err)
//				mtx.Unlock()
//			}
//			if res != "" {
//				*exconfErrMsg = res
//			}
//			wg2.Done()
//		}()
//	}
//	wg2.Wait()
//	if exReqTimes == 0 {
//		*existenceMap = append(*existenceMap, false)
//	}
//}
//
//func (c *ExistenceConf) batchExistenceConfRequest(bpID int, product string, plant string, batch string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"BusinessPartner": bpID,
//		"Product":         product,
//		"Plant":           plant,
//		"Batch":           batch,
//	})
//	exist := false
//	defer func() {
//		mtx.Lock()
//		*existenceMap = append(*existenceMap, exist)
//		mtx.Unlock()
//	}()
//
//	req, err := jsonTypeConversion[Returns](input)
//	if err != nil {
//		return "", xerrors.Errorf("request create error: %w", err)
//	}
//	req.BatchReturn.BusinessPartner = bpID
//	req.BatchReturn.Product = product
//	req.BatchReturn.Plant = plant
//	req.BatchReturn.Batch = batch
//
//	exist, err = c.exconfRequest(req, queueName, log)
//	if err != nil {
//		return "", err
//	}
//	if !exist {
//		return keys.fail(), nil
//	}
//	return "", nil
//}
//
//func getItemBatchExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, string, string, string) {
//	var bpID int
//	var product string
//	var plant string
//	var batch string
//
//	switch mapper.Field {
//	case "DeliverToPlantBatch":
//		if item.DeliverToParty == nil ||
//			item.Product == nil ||
//			item.DeliverToPlant == nil ||
//			item.DeliverToPlantBatch == nil {
//			bpID = 0
//			product = ""
//			plant = ""
//			batch = ""
//		} else {
//			bpID = *item.DeliverToParty
//			product = *item.Product
//			plant = *item.DeliverToPlant
//			batch = *item.DeliverToPlantBatch
//		}
//
//	case "DeliverFromPlantBatch":
//		if item.DeliverFromParty == nil ||
//			item.Product == nil ||
//			item.DeliverFromPlant == nil ||
//			item.DeliverFromPlantBatch == nil {
//			bpID = 0
//			product = ""
//			plant = ""
//			batch = ""
//		} else {
//			bpID = *item.DeliverFromParty
//			product = *item.Product
//			plant = *item.DeliverFromPlant
//			batch = *item.DeliverFromPlantBatch
//		}
//
//	case "StockConfirmationPlantBatch":
//		if item.StockConfirmationBusinessPartner == nil ||
//			item.Product == nil ||
//			item.StockConfirmationPlant == nil ||
//			item.StockConfirmationPlantBatch == nil {
//			bpID = 0
//			product = ""
//			plant = ""
//			batch = ""
//		} else {
//			bpID = *item.StockConfirmationBusinessPartner
//			product = *item.Product
//			plant = *item.StockConfirmationPlant
//			batch = *item.StockConfirmationPlantBatch
//		}
//
//	case "ProductionPlantBatch":
//		if item.ProductionPlantBusinessPartner == nil ||
//			item.Product == nil ||
//			item.ProductionPlant == nil ||
//			item.ProductionPlantBatch == nil {
//			bpID = 0
//			product = ""
//			plant = ""
//			batch = ""
//		} else {
//			bpID = *item.ProductionPlantBusinessPartner
//			product = *item.Product
//			plant = *item.ProductionPlant
//			batch = *item.ProductionPlantBatch
//		}
//	}
//
//	return bpID, product, plant, batch
//}
