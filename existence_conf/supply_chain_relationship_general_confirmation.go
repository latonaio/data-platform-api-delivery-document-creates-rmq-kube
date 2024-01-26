package existence_conf

//func (c *ExistenceConf) supplyChainRelationshipGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//	headers := make([]dpfm_api_input_reader.Header, 0, 1)
//	headers = append(headers, input.Header)
//	for _, header := range headers {
//		supplyChainRelationshipID, buyer, seller := getSupplyChainRelationshipGeneralExistenceConfKey(mapper, &header, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(supplyChainRelationshipID) || isZero(buyer) || isZero(seller) {
//				wg2.Done()
//				return
//			}
//			res, err := c.supplyChainRelationshipGeneralExistenceConfRequest(supplyChainRelationshipID, buyer, seller, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) supplyChainRelationshipGeneralExistenceConfRequest(supplyChainRelationshipID int, buyer int, seller int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"SupplyChainRelationshipID": supplyChainRelationshipID,
//		"Buyer":                     buyer,
//		"Seller":                    seller,
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
//	data := &SupplyChainRelationshipGeneralReturn{
//		SupplyChainRelationshipID: supplyChainRelationshipID,
//		Buyer:                     buyer,
//		Seller:                    seller,
//	}
//	req.SupplyChainRelationshipGeneralReturn = data
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
//func getSupplyChainRelationshipGeneralExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (int, int, int) {
//	var supplyChainRelationshipID, buyer, seller int
//
//	if header.SupplyChainRelationshipID == nil ||
//		header.Buyer == nil ||
//		header.Seller == nil {
//		supplyChainRelationshipID = 0
//		buyer = 0
//		seller = 0
//	} else {
//		supplyChainRelationshipID = *header.SupplyChainRelationshipID
//		buyer = *header.Buyer
//		seller = *header.Seller
//	}
//
//	return supplyChainRelationshipID, buyer, seller
//}
