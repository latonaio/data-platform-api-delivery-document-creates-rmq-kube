package existence_conf

//func (c *ExistenceConf) supplyChainRelationshipStockConfPlantRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		supplyChainRelationshipID, supplyChainRelationshipStockConfPlantID, buyer, seller, stockConfirmationBusinessPartner, stockConfirmationPlant := getSupplyChainRelationshipStockConfPlantRelationExistenceConfKey(mapper, &item, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(supplyChainRelationshipID) || isZero(supplyChainRelationshipStockConfPlantID) ||
//				isZero(buyer) || isZero(seller) || isZero(stockConfirmationBusinessPartner) || isZero(stockConfirmationPlant) {
//				wg2.Done()
//				return
//			}
//			res, err := c.supplyChainRelationshipStockConfPlantRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipStockConfPlantID, buyer, seller, stockConfirmationBusinessPartner, stockConfirmationPlant, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) supplyChainRelationshipStockConfPlantRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipStockConfPlantID int, buyer int, seller int, stockConfirmationBusinessPartner int, stockConfirmationPlant string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"SupplyChainRelationshipID":               supplyChainRelationshipID,
//		"SupplyChainRelationshipStockConfPlantID": supplyChainRelationshipStockConfPlantID,
//		"Buyer":                            buyer,
//		"Seller":                           seller,
//		"StockConfirmationBusinessPartner": stockConfirmationBusinessPartner,
//		"StockConfirmationPlant":           stockConfirmationPlant,
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
//	data := &SupplyChainRelationshipStockConfPlantRelationReturn{
//		SupplyChainRelationshipID:               supplyChainRelationshipID,
//		SupplyChainRelationshipStockConfPlantID: supplyChainRelationshipStockConfPlantID,
//		Buyer:                                   buyer,
//		Seller:                                  seller,
//		StockConfirmationBusinessPartner:        stockConfirmationBusinessPartner,
//		StockConfirmationPlant:                  stockConfirmationPlant,
//	}
//	req.SupplyChainRelationshipStockConfPlantRelationReturn = data
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
//func getSupplyChainRelationshipStockConfPlantRelationExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, int, int, int, int, string) {
//	var supplyChainRelationshipID, supplyChainRelationshipStockConfPlantID, buyer, seller, stockConfirmationBusinessPartner int
//	var stockConfirmationPlant string
//
//	if item.SupplyChainRelationshipID == nil ||
//		item.SupplyChainRelationshipStockConfPlantID == nil ||
//		item.Buyer == nil ||
//		item.Seller == nil ||
//		item.StockConfirmationBusinessPartner == nil ||
//		item.StockConfirmationPlant == nil {
//		supplyChainRelationshipID = 0
//		supplyChainRelationshipStockConfPlantID = 0
//		buyer = 0
//		seller = 0
//		stockConfirmationBusinessPartner = 0
//		stockConfirmationPlant = ""
//	} else {
//		supplyChainRelationshipID = *item.SupplyChainRelationshipID
//		supplyChainRelationshipStockConfPlantID = *item.SupplyChainRelationshipStockConfPlantID
//		buyer = *item.Buyer
//		seller = *item.Seller
//		stockConfirmationBusinessPartner = *item.StockConfirmationBusinessPartner
//		stockConfirmationPlant = *item.StockConfirmationPlant
//	}
//
//	return supplyChainRelationshipID, supplyChainRelationshipStockConfPlantID, buyer, seller, stockConfirmationBusinessPartner, stockConfirmationPlant
//}
