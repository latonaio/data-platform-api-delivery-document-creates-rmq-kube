package existence_conf

//func (c *ExistenceConf) itemQuantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		quantityUnit := getItemQuantityUnitExistenceConfKey(mapper, &item, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(quantityUnit) {
//				wg2.Done()
//				return
//			}
//			res, err := c.quantityUnitExistenceConfRequest(quantityUnit, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) quantityUnitExistenceConfRequest(quantityUnit string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"QuantityUnit": quantityUnit,
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
//	req.QuantityUnitReturn.QuantityUnit = quantityUnit
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
//func getItemQuantityUnitExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) string {
//	var quantityUnit string
//
//	switch mapper.Field {
//	case "ItemWeightUnit":
//		if item.ItemWeightUnit == nil {
//			quantityUnit = ""
//		} else {
//			quantityUnit = *item.ItemWeightUnit
//		}
//	case "InternalCapacityQuantityUnit":
//		if item.InternalCapacityQuantityUnit == nil {
//			quantityUnit = ""
//		} else {
//			quantityUnit = *item.InternalCapacityQuantityUnit
//		}
//	}
//
//	return quantityUnit
//}
