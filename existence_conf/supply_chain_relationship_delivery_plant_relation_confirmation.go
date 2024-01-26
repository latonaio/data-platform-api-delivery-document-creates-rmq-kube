package existence_conf

//func (c *ExistenceConf) supplyChainRelationshipDeliveryPlantRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//	headers := make([]dpfm_api_input_reader.Header, 0, 1)
//	headers = append(headers, input.Header)
//	for _, header := range headers {
//		supplyChainRelationshipID, supplyChainRelationshipDeliveryID, supplyChainRelationshipDeliveryPlantID, buyer, seller, deliverToParty, deliverFromParty, deliverToPlant, deliverFromPlant := getSupplyChainRelationshipDeliveryPlantRelationExistenceConfKey(mapper, &header, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(supplyChainRelationshipID) || isZero(supplyChainRelationshipDeliveryID) ||
//				isZero(supplyChainRelationshipDeliveryPlantID) || isZero(buyer) || isZero(seller) ||
//				isZero(deliverToParty) || isZero(deliverFromParty) || isZero(deliverToParty) || isZero(deliverFromParty) {
//				wg2.Done()
//				return
//			}
//			res, err := c.supplyChainRelationshipDeliveryPlantRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipDeliveryID, supplyChainRelationshipDeliveryPlantID, buyer, seller, deliverToParty, deliverFromParty, deliverToPlant, deliverFromPlant, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) supplyChainRelationshipDeliveryPlantRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipDeliveryID int, supplyChainRelationshipDeliveryPlantID int, buyer int, seller int, deliverToParty int, deliverFromParty int, deliverToPlant string, deliverFromPlant string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"SupplyChainRelationshipID":              supplyChainRelationshipID,
//		"SupplyChainRelationshipDeliveryID":      supplyChainRelationshipDeliveryID,
//		"SupplyChainRelationshipDeliveryPlantID": supplyChainRelationshipDeliveryPlantID,
//		"Buyer":                                  buyer,
//		"Seller":                                 seller,
//		"DeliverToParty":                         deliverToParty,
//		"DeliverFromParty":                       deliverFromParty,
//		"DeliverToPlant":                         deliverToPlant,
//		"DeliverFromPlant":                       deliverFromPlant,
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
//	data := &SupplyChainRelationshipDeliveryPlantRelationReturn{
//		SupplyChainRelationshipID:              supplyChainRelationshipID,
//		SupplyChainRelationshipDeliveryID:      supplyChainRelationshipDeliveryID,
//		SupplyChainRelationshipDeliveryPlantID: supplyChainRelationshipDeliveryPlantID,
//		Buyer:                                  buyer,
//		Seller:                                 seller,
//		DeliverToParty:                         deliverToParty,
//		DeliverFromParty:                       deliverFromParty,
//		DeliverToPlant:                         deliverToPlant,
//		DeliverFromPlant:                       deliverFromPlant,
//	}
//	req.SupplyChainRelationshipDeliveryPlantRelationReturn = data
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
//func getSupplyChainRelationshipDeliveryPlantRelationExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (int, int, int, int, int, int, int, string, string) {
//	var supplyChainRelationshipID, supplyChainRelationshipDeliveryID, supplyChainRelationshipDeliveryPlantID, buyer, seller, deliverToParty, deliverFromParty int
//	var deliverToPlant, deliverFromPlant string
//
//	if header.SupplyChainRelationshipID == nil ||
//		header.SupplyChainRelationshipDeliveryID == nil ||
//		header.SupplyChainRelationshipDeliveryPlantID == nil ||
//		header.Buyer == nil ||
//		header.Seller == nil ||
//		header.DeliverToParty == nil ||
//		header.DeliverFromParty == nil ||
//		header.DeliverToPlant == nil ||
//		header.DeliverFromPlant == nil {
//		supplyChainRelationshipID = 0
//		supplyChainRelationshipDeliveryID = 0
//		supplyChainRelationshipDeliveryPlantID = 0
//		buyer = 0
//		seller = 0
//		deliverToParty = 0
//		deliverFromParty = 0
//		deliverToPlant = ""
//		deliverFromPlant = ""
//	} else {
//		supplyChainRelationshipID = *header.SupplyChainRelationshipID
//		supplyChainRelationshipDeliveryID = *header.SupplyChainRelationshipDeliveryID
//		supplyChainRelationshipDeliveryPlantID = *header.SupplyChainRelationshipDeliveryPlantID
//		buyer = *header.Buyer
//		seller = *header.Seller
//		deliverToParty = *header.DeliverToParty
//		deliverFromParty = *header.DeliverFromParty
//		deliverToPlant = *header.DeliverToPlant
//		deliverFromPlant = *header.DeliverFromPlant
//	}
//
//	return supplyChainRelationshipID, supplyChainRelationshipDeliveryID, supplyChainRelationshipDeliveryPlantID, buyer, seller, deliverToParty, deliverFromParty, deliverToPlant, deliverFromPlant
//}
