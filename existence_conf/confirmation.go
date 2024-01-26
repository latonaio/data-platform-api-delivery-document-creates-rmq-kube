package existence_conf

//type ExistenceConf struct {
//	ctx context.Context
//
//	c   *config.Conf
//	rmq *rabbitmq.RabbitmqClient
//	db  *database.Mysql
//}
//
//func NewExistenceConf(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *ExistenceConf {
//	return &ExistenceConf{
//		ctx: ctx,
//		c:   c,
//		rmq: rmq,
//		db:  db,
//	}
//}
//
//// Confirm returns existenceMap, allExist, err
//func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, ssdc *dpfm_api_output_formatter.SDC, accepter []string, l *logger.Logger) (allExist bool, errs []error) {
//	var resMsg string
//	existenceMap := make([]bool, 0, 12)
//	wg := sync.WaitGroup{}
//	mtx := &sync.Mutex{}
//
//	serviceLabel := data.ServiceLabel
//	exConfMapper, err := c.getExConfMapper(serviceLabel)
//	if err != nil {
//		errs = append(errs, err)
//		return false, errs
//	}
//
//	for _, v := range *exConfMapper {
//		if v.Tabletag == nil {
//			continue
//		}
//		tabletag := *v.Tabletag
//		apiName := v.APIName
//		if !contains(accepter, apiName) {
//			continue
//		}
//		switch tabletag {
//		case "BusinessPartnerGeneral":
//			switch apiName {
//			case "Header":
//				wg.Add(1)
//				go c.headerBPGeneralExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//			case "Item":
//				wg.Add(1)
//				go c.itemBPGeneralExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//			}
//		case "PlantGeneral":
//			switch apiName {
//			case "Header":
//				wg.Add(1)
//				go c.headerPlantGeneralExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//			case "Item":
//				wg.Add(1)
//				go c.itemPlantGeneralExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//			}
//		case "Incoterms":
//			wg.Add(1)
//			go c.headerIncotermsExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "Country":
//			wg.Add(1)
//			go c.headerCountryExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "QuantityUnit":
//			wg.Add(1)
//			go c.itemQuantityUnitExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipGeneral":
//			wg.Add(1)
//			go c.supplyChainRelationshipGeneralExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipBillingRelation":
//			wg.Add(1)
//			go c.supplyChainRelationshipBillingRelationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipPaymentRelation":
//			wg.Add(1)
//			go c.supplyChainRelationshipPaymentRelationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipDeliveryRelation":
//			wg.Add(1)
//			go c.supplyChainRelationshipDeliveryRelationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipDeliveryPlantRelation":
//			wg.Add(1)
//			go c.supplyChainRelationshipDeliveryPlantRelationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipProductionPlantRelation":
//			wg.Add(1)
//			go c.supplyChainRelationshipProductionPlantRelationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "SupplyChainRelationshipStockConfPlantRelation":
//			wg.Add(1)
//			go c.supplyChainRelationshipStockConfPlantRelationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "Batch":
//			wg.Add(1)
//			go c.itemBatchExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		case "StorageLocation":
//			wg.Add(1)
//			go c.itemStorageLocationExistenceConf(v, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
//		}
//	}
//
//	wg.Wait()
//
//	if len(errs) != 0 {
//		return false, errs
//	}
//
//	ssdc.ExconfResult = getBoolPtr(true)
//
//	if len(resMsg) != 0 {
//		ssdc.ExconfResult = getBoolPtr(false)
//		ssdc.ExconfError = resMsg
//		return false, nil
//	}
//
//	return true, nil
//}
//
//func (c *ExistenceConf) getExConfMapper(serviceLabel string) (*[]ExConfMapper, error) {
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformCommonSettingsMysqlKube.data_platform_ex_conf_api_mapper_data
//		WHERE ServiceLabel = ?;`, serviceLabel,
//	)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	res := make([]ExConfMapper, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		data := ExConfMapper{}
//
//		err := rows.Scan(
//			&data.ServiceLabel,
//			&data.APIType,
//			&data.APIName,
//			&data.Field,
//			&data.ExConfAPIServiceName,
//			&data.ExConfAPIName,
//			&data.ExConfAPIQueueName,
//			&data.ExConfProgramFileName,
//			&data.Tabletag,
//			&data.TableConfirmed,
//			&data.ExConfAPITyp,
//		)
//		if err != nil {
//			return nil, err
//		}
//
//		res = append(res, data)
//	}
//	if i == 0 {
//		return nil, fmt.Errorf("'data_platform_ex_conf_api_mapper_data'テーブルに対象のレコードが存在しません。")
//	}
//
//	return &res, nil
//}
//
//func getBoolPtr(b bool) *bool {
//	return &b
//}
//
//func jsonTypeConversion[T any](data interface{}) (T, error) {
//	var dist T
//	b, err := json.Marshal(data)
//	if err != nil {
//		return dist, xerrors.Errorf("Marshal error: %w", err)
//	}
//	err = json.Unmarshal(b, &dist)
//	if err != nil {
//		return dist, xerrors.Errorf("Unmarshal error: %w", err)
//	}
//	return dist, nil
//}
//func confKeyExistence(res map[string]interface{}) bool {
//	if res == nil {
//		return false
//	}
//	raw, ok := res["ExistenceConf"]
//	exist := fmt.Sprintf("%v", raw)
//	if ok {
//		return strings.ToLower(exist) == "true"
//	}
//
//	return false
//}
//func (c *ExistenceConf) exconfRequest(req interface{}, queueName string, log *logger.Logger) (bool, error) {
//	res, err := c.rmq.SessionKeepRequest(nil, queueName, req)
//	if err != nil {
//		return false, xerrors.Errorf("response error: %w", err)
//	}
//	res.Success()
//	exist := confKeyExistence(res.Data())
//	log.Info(res.Data())
//	return exist, nil
//}
//
//type result struct {
//	keys map[string]interface{}
//}
//
//func newResult(keys map[string]interface{}) *result {
//	return &result{
//		keys: keys,
//	}
//}
//
//func (r *result) fail() string {
//	txt := ""
//	for k, v := range r.keys {
//		txt = fmt.Sprintf("%s%s:%v, ", k, v)
//	}
//	txt = fmt.Sprintf("%s does not exist", txt)
//	return txt
//}
//
//func getQueueName(mapper ExConfMapper) (string, error) {
//	var err error
//
//	if mapper.ExConfAPIQueueName == nil {
//		err = xerrors.Errorf("cannot specify null queue name")
//		return "", err
//	}
//	queueName := *mapper.ExConfAPIQueueName
//
//	return queueName, nil
//}
//
//func contains(slice []string, target string) bool {
//	for _, v := range slice {
//		if v == target {
//			return true
//		}
//	}
//	return false
//}
//
//func isZero[T comparable](obj T) bool {
//	var zero T
//	return obj == zero
//}
