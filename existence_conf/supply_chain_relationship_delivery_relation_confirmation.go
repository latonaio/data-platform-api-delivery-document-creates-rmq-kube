package existence_conf

//func (c *ExistenceConf) supplyChainRelationshipDeliveryRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//	headers := make([]dpfm_api_input_reader.Header, 0, 1)
//	headers = append(headers, input.Header)
//	for _, header := range headers {
//		supplyChainRelationshipID, supplyChainRelationshipDeliveryID, buyer, seller, deliverToParty, deliverFromParty := getSupplyChainRelationshipDeliveryRelationExistenceConfKey(mapper, &header, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(supplyChainRelationshipID) || isZero(supplyChainRelationshipDeliveryID) ||
//				isZero(buyer) || isZero(seller) || isZero(deliverToParty) || isZero(deliverFromParty) {
//				wg2.Done()
//				return
//			}
//			res, err := c.supplyChainRelationshipDeliveryRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipDeliveryID, buyer, seller, deliverToParty, deliverFromParty, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) supplyChainRelationshipDeliveryRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipDeliveryID int, buyer int, seller int, deliverToParty int, deliverFromParty int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"SupplyChainRelationshipID":         supplyChainRelationshipID,
//		"SupplyChainRelationshipDeliveryID": supplyChainRelationshipDeliveryID,
//		"Buyer":                             buyer,
//		"Seller":                            seller,
//		"DeliverToParty":                    deliverToParty,
//		"DeliverFromParty":                  deliverFromParty,
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
//	data := &SupplyChainRelationshipDeliveryRelationReturn{
//		SupplyChainRelationshipID:         supplyChainRelationshipID,
//		SupplyChainRelationshipDeliveryID: supplyChainRelationshipDeliveryID,
//		Buyer:                             buyer,
//		Seller:                            seller,
//		DeliverToParty:                    deliverToParty,
//		DeliverFromParty:                  deliverFromParty,
//	}
//	req.SupplyChainRelationshipDeliveryRelationReturn = data
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
//func getSupplyChainRelationshipDeliveryRelationExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (int, int, int, int, int, int) {
//	var supplyChainRelationshipID, supplyChainRelationshipDeliveryID, buyer, seller, deliverToParty, deliverFromParty int
//
//	if header.SupplyChainRelationshipID == nil ||
//		header.SupplyChainRelationshipDeliveryID == nil ||
//		header.Buyer == nil ||
//		header.Seller == nil ||
//		header.DeliverToParty == nil ||
//		header.DeliverFromParty == nil {
//		supplyChainRelationshipID = 0
//		supplyChainRelationshipDeliveryID = 0
//		buyer = 0
//		seller = 0
//		deliverToParty = 0
//		deliverFromParty = 0
//	} else {
//		supplyChainRelationshipID = *header.SupplyChainRelationshipID
//		supplyChainRelationshipDeliveryID = *header.SupplyChainRelationshipDeliveryID
//		buyer = *header.Buyer
//		seller = *header.Seller
//		deliverToParty = *header.DeliverToParty
//		deliverFromParty = *header.DeliverFromParty
//	}
//
//	return supplyChainRelationshipID, supplyChainRelationshipDeliveryID, buyer, seller, deliverToParty, deliverFromParty
//}
