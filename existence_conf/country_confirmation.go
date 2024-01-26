package existence_conf

//func (c *ExistenceConf) headerCountryExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	headers := make([]dpfm_api_input_reader.Header, 0, 1)
//	headers = append(headers, input.Header)
//	for _, header := range headers {
//		country := getHeaderCountryExistenceConfKey(mapper, &header, exconfErrMsg)
//		queueName, err := getQueueName(mapper)
//		if err != nil {
//			*errs = append(*errs, err)
//			return
//		}
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(country) {
//				wg2.Done()
//				return
//			}
//			res, err := c.countryExistenceConfRequest(country, queueName, input, existenceMap, mtx, log)
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
//func (c *ExistenceConf) countryExistenceConfRequest(country string, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"Country": country,
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
//	req.CountryReturn.Country = country
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
//func getHeaderCountryExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) string {
//	var country string
//
//	switch mapper.Field {
//	case "BillToCountry":
//		if header.BillToCountry == nil {
//			country = ""
//		} else {
//			country = *header.BillToCountry
//		}
//
//	case "BillFromCountry":
//		if header.BillFromCountry == nil {
//			country = ""
//		} else {
//			country = *header.BillFromCountry
//		}
//	}
//
//	return country
//}
