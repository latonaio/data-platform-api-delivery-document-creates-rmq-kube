package existence_conf

//func (c *ExistenceConf) headerPlantGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	headers := make([]dpfm_api_input_reader.Header, 0, 1)
//	headers = append(headers, input.Header)
//	for _, header := range headers {
//		plant, bpID := getHeaderPlantGeneralExistenceConfKey(mapper, &header, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(plant) || isZero(bpID) {
//				wg2.Done()
//				return
//			}
//			res, err := c.plantGeneralExistenceConfRequest(plant, bpID, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) itemPlantGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		plant, bpID := getItemPlantGeneralExistenceConfKey(mapper, &item, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(plant) || isZero(bpID) {
//				wg2.Done()
//				return
//			}
//			res, err := c.plantGeneralExistenceConfRequest(plant, bpID, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) plantGeneralExistenceConfRequest(plant string, bpID int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"BusinessPartner": bpID,
//		"Plant":           plant,
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
//	req.PlantGeneralReturn.Plant = plant
//	req.PlantGeneralReturn.BusinessPartner = bpID
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
//func getHeaderPlantGeneralExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (string, int) {
//	var plant string
//	var bpID int
//
//	switch mapper.Field {
//	case "DeliverToPlant":
//		if header.DeliverToPlant == nil || header.DeliverToParty == nil {
//			plant = ""
//			bpID = 0
//		} else {
//			plant = *header.DeliverToPlant
//			bpID = *header.DeliverToParty
//		}
//	case "DeliverFromPlant":
//		if header.DeliverFromPlant == nil || header.DeliverFromParty == nil {
//			plant = ""
//			bpID = 0
//		} else {
//			plant = *header.DeliverFromPlant
//			bpID = *header.DeliverFromParty
//		}
//	}
//
//	return plant, bpID
//}
//
//func getItemPlantGeneralExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (string, int) {
//	var plant string
//	var bpID int
//
//	switch mapper.Field {
//	case "StockConfirmationPlant":
//		if item.StockConfirmationPlant == nil || item.StockConfirmationBusinessPartner == nil {
//			plant = ""
//			bpID = 0
//		} else {
//			plant = *item.StockConfirmationPlant
//			bpID = *item.StockConfirmationBusinessPartner
//		}
//	case "ProductionPlant":
//		if item.ProductionPlant == nil || item.ProductionPlantBusinessPartner == nil {
//			plant = ""
//			bpID = 0
//		} else {
//			plant = *item.ProductionPlant
//			bpID = *item.ProductionPlantBusinessPartner
//		}
//	}
//
//	return plant, bpID
//}
