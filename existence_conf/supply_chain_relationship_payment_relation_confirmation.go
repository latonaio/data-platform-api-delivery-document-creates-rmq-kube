package existence_conf

//func (c *ExistenceConf) supplyChainRelationshipPaymentRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//	headers := make([]dpfm_api_input_reader.Header, 0, 1)
//	headers = append(headers, input.Header)
//	for _, header := range headers {
//		supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee := getSupplyChainRelationshipPaymentRelationExistenceConfKey(mapper, &header, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(supplyChainRelationshipID) || isZero(supplyChainRelationshipBillingID) || isZero(supplyChainRelationshipPaymentID) ||
//				isZero(buyer) || isZero(seller) || isZero(billToParty) || isZero(billFromParty) || isZero(payer) || isZero(payee) {
//				wg2.Done()
//				return
//			}
//			res, err := c.supplyChainRelationshipPaymentRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) supplyChainRelationshipPaymentRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipBillingID int, supplyChainRelationshipPaymentID int, buyer int, seller int, billToParty int, billFromParty int, payer int, payee int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"SupplyChainRelationshipID":        supplyChainRelationshipID,
//		"SupplyChainRelationshipBillingID": supplyChainRelationshipBillingID,
//		"Buyer":                            buyer,
//		"Seller":                           seller,
//		"BillToParty":                      billToParty,
//		"BillFromParty":                    billFromParty,
//		"Payer":                            payer,
//		"Payee":                            payee,
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
//	data := &SupplyChainRelationshipPaymentRelationReturn{
//		SupplyChainRelationshipID:        supplyChainRelationshipID,
//		SupplyChainRelationshipBillingID: supplyChainRelationshipBillingID,
//		SupplyChainRelationshipPaymentID: supplyChainRelationshipPaymentID,
//		Buyer:                            buyer,
//		Seller:                           seller,
//		BillToParty:                      billToParty,
//		BillFromParty:                    billFromParty,
//		Payer:                            payer,
//		Payee:                            payee,
//	}
//	req.SupplyChainRelationshipPaymentRelationReturn = data
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
//func getSupplyChainRelationshipPaymentRelationExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (int, int, int, int, int, int, int, int, int) {
//	var supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee int
//
//	if header.SupplyChainRelationshipID == nil ||
//		header.SupplyChainRelationshipBillingID == nil ||
//		header.SupplyChainRelationshipPaymentID == nil ||
//		header.Buyer == nil ||
//		header.Seller == nil ||
//		header.BillToParty == nil ||
//		header.BillFromParty == nil ||
//		header.Payer == nil ||
//		header.Payee == nil {
//		supplyChainRelationshipID = 0
//		supplyChainRelationshipBillingID = 0
//		supplyChainRelationshipPaymentID = 0
//		buyer = 0
//		seller = 0
//		billToParty = 0
//		billFromParty = 0
//		payer = 0
//		payee = 0
//	} else {
//		supplyChainRelationshipID = *header.SupplyChainRelationshipID
//		supplyChainRelationshipBillingID = *header.SupplyChainRelationshipBillingID
//		supplyChainRelationshipPaymentID = *header.SupplyChainRelationshipPaymentID
//		buyer = *header.Buyer
//		seller = *header.Seller
//		billToParty = *header.BillToParty
//		billFromParty = *header.BillFromParty
//		payer = *header.Payer
//		payee = *header.Payee
//	}
//
//	return supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee
//}
