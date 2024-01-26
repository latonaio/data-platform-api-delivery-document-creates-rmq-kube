package existence_conf

//func (c *ExistenceConf) supplyChainRelationshipProductionPlantRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		supplyChainRelationshipID, supplyChainRelationshipProductionPlantID, buyer, seller, productionPlantBusinessPartner, productionPlant := getSupplyChainRelationshipProductionPlantRelationExistenceConfKey(mapper, &item, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(supplyChainRelationshipID) || isZero(supplyChainRelationshipProductionPlantID) ||
//				isZero(buyer) || isZero(seller) || isZero(productionPlantBusinessPartner) || isZero(productionPlant) {
//				wg2.Done()
//				return
//			}
//			res, err := c.supplyChainRelationshipProductionPlantRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipProductionPlantID, buyer, seller, productionPlantBusinessPartner, productionPlant, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) supplyChainRelationshipProductionPlantRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipProductionPlantID int, buyer int, seller int, productionPlantBusinessPartner int, productionPlant string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"SupplyChainRelationshipID":                supplyChainRelationshipID,
//		"SupplyChainRelationshipProductionPlantID": supplyChainRelationshipProductionPlantID,
//		"Buyer":                          buyer,
//		"Seller":                         seller,
//		"ProductionPlantBusinessPartner": productionPlantBusinessPartner,
//		"ProductionPlant":                productionPlant,
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
//
//	data := &SupplyChainRelationshipProductionPlantRelationReturn{
//		SupplyChainRelationshipID:                supplyChainRelationshipID,
//		SupplyChainRelationshipProductionPlantID: supplyChainRelationshipProductionPlantID,
//		Buyer:                                    buyer,
//		Seller:                                   seller,
//		ProductionPlantBusinessPartner:           productionPlantBusinessPartner,
//		ProductionPlant:                          productionPlant,
//	}
//	req.SupplyChainRelationshipProductionPlantRelationReturn = data
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
//func getSupplyChainRelationshipProductionPlantRelationExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, int, int, int, int, string) {
//	var supplyChainRelationshipID, supplyChainRelationshipProductionPlantID, buyer, seller, productionPlantBusinessPartner int
//	var productionPlant string
//
//	if item.SupplyChainRelationshipID == nil ||
//		item.SupplyChainRelationshipProductionPlantID == nil ||
//		item.Buyer == nil ||
//		item.Seller == nil ||
//		item.ProductionPlantBusinessPartner == nil ||
//		item.ProductionPlant == nil {
//		supplyChainRelationshipID = 0
//		supplyChainRelationshipProductionPlantID = 0
//		buyer = 0
//		seller = 0
//		productionPlantBusinessPartner = 0
//		productionPlant = ""
//	} else {
//		supplyChainRelationshipID = *item.SupplyChainRelationshipID
//		supplyChainRelationshipProductionPlantID = *item.SupplyChainRelationshipProductionPlantID
//		buyer = *item.Buyer
//		seller = *item.Seller
//		productionPlantBusinessPartner = *item.ProductionPlantBusinessPartner
//		productionPlant = *item.ProductionPlant
//	}
//
//	return supplyChainRelationshipID, supplyChainRelationshipProductionPlantID, buyer, seller, productionPlantBusinessPartner, productionPlant
//}
