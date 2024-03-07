package dpfm_api_output_formatter

import (
	"data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-delivery-document-creates-rmq-kube/sub_func_complementer"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *subfuncSDC.Message.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToItemPickingCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemPicking, error) {
	itemPickings := make([]ItemPicking, 0)

	for _, data := range *subfuncSDC.Message.ItemPicking {
		itemPicking, err := TypeConverter[*ItemPicking](data)
		if err != nil {
			return nil, err
		}

		itemPickings = append(itemPickings, *itemPicking)
	}

	return &itemPickings, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToItemPickingUpdates(itemPickingUpdates *[]dpfm_api_processing_formatter.ItemPickingUpdates) (*[]ItemPicking, error) {
	itemPickings := make([]ItemPicking, 0)

	for _, data := range *itemPickingUpdates {
		itemPicking, err := TypeConverter[*ItemPicking](data)
		if err != nil {
			return nil, err
		}

		itemPickings = append(itemPickings, *itemPicking)
	}

	return &itemPickings, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func ConvertToHeaderFromOrders(
	rows *sql.Rows,
	deliveryDocumentIssuedID int,
) (*[]sub_func_complementer.Header, error) {
	defer rows.Close()
	header := make([]sub_func_complementer.Header, 0)

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")
	isCancelled := false
	isMarkedForDeletion := false

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.ProductSpecification,
			&pm.MarkingOfMaterial,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityStartTime,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverToPlantBatchValidityEndTime,
			&pm.DeliverFromPlantTimeZone,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityStartTime,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliverFromPlantBatchValidityEndTime,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityStartTime,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.StockConfirmationPlantBatchValidityEndTime,
			&pm.ServicesRenderingDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.QuantityPerPackage,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ProductWeightUnit,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.InvoiceDocumentDate,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityStartTime,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.ProductionPlantBatchValidityEndTime,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.InspectionPlan,
			&pm.InspectionLot,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.Contract,
			&pm.ContractItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.IssuingStatus,
			&pm.ReceivingStatus,
			&pm.ItemBillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.Equipment,
			&pm.FreightAgreement,
			&pm.FreightAgreementItem,
			&pm.ItemBlockStatus,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemBillingBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		headerCompleteDeliveryIsDefined := false
		headerDeliveryStatus := "NP"

		var orderID *int

		dataOrderID := data.OrderID
		if &dataOrderID != nil {
			orderID = &dataOrderID
		}

		var orderItem *int

		dataOrderItem := data.OrderItem
		if &dataOrderItem != nil {
			orderItem = &dataOrderItem
		}

		header = append(header, sub_func_complementer.Header{
			DeliveryDocument:                       deliveryDocumentIssuedID,
			DeliveryDocumentDate:                   formattedDate,
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      *data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: *data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         *data.DeliverToParty,
			DeliverFromParty:                       *data.DeliverFromParty,
			DeliverToPlant:                         *data.DeliverToPlant,
			DeliverFromPlant:                       *data.DeliverFromPlant,
			OrderID:                                orderID,
			OrderItem:                              orderItem,
			Contract:                               data.Contract,
			ContractItem:                           data.ContractItem,
			Project:                                data.Project,
			WBSElement:                             data.WBSElement,
			PlannedGoodsIssueDate:                  data.RequestedDeliveryDate,
			PlannedGoodsIssueTime:                  "00:00:00",
			PlannedGoodsReceiptDate:                data.RequestedDeliveryDate,
			PlannedGoodsReceiptTime:                data.RequestedDeliveryTime,
			HeaderCompleteDeliveryIsDefined:        &headerCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:                   &headerDeliveryStatus,
			Incoterms:                              data.Incoterms,
			HeaderGrossWeight:                      data.ItemGrossWeight,
			HeaderNetWeight:                        data.ItemNetWeight,
			HeaderWeightUnit:                       data.ProductWeightUnit,
			CreationDate:                           formattedDate,
			CreationTime:                           formattedTime,
			LastChangeDate:                         formattedDate,
			LastChangeTime:                         formattedTime,
			IsCancelled:                            &isCancelled,
			IsMarkedForDeletion:                    &isMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil
}

func ConvertToItemFromOrders(
	rows *sql.Rows,
	deliveryDocumentIssuedID int,
) (*[]sub_func_complementer.Item, error) {
	defer rows.Close()
	item := make([]sub_func_complementer.Item, 0)

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")
	isCancelled := false
	isMarkedForDeletion := false

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.ProductSpecification,
			&pm.MarkingOfMaterial,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityStartTime,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverToPlantBatchValidityEndTime,
			&pm.DeliverFromPlantTimeZone,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityStartTime,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliverFromPlantBatchValidityEndTime,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityStartTime,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.StockConfirmationPlantBatchValidityEndTime,
			&pm.ServicesRenderingDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.QuantityPerPackage,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ProductWeightUnit,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.InvoiceDocumentDate,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityStartTime,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.ProductionPlantBatchValidityEndTime,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.InspectionPlan,
			&pm.InspectionLot,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.Contract,
			&pm.ContractItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.IssuingStatus,
			&pm.ReceivingStatus,
			&pm.ItemBillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.Equipment,
			&pm.FreightAgreement,
			&pm.FreightAgreementItem,
			&pm.ItemBlockStatus,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemBillingBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemCompleteDeliveryIsDefined := false

		var orderID *int

		dataOrderID := data.OrderID
		if &dataOrderID != nil {
			orderID = &dataOrderID
		}

		var orderItem *int

		dataOrderItem := data.OrderItem
		if &dataOrderItem != nil {
			orderItem = &dataOrderItem
		}

		itemDeliveryBlockStatus := false

		item = append(item, sub_func_complementer.Item{
			DeliveryDocument:                        deliveryDocumentIssuedID,
			DeliveryDocumentItem:                    data.OrderItem,
			DeliveryDocumentItemCategory:            data.OrderItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:       *data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:  *data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID: *data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			DeliverToParty:                          *data.DeliverToParty,
			DeliverFromParty:                        *data.DeliverFromParty,
			DeliverToPlant:                          *data.DeliverToPlant,
			DeliverFromPlant:                        *data.DeliverFromPlant,
			Product:                                 data.Product,
			SizeOrDimensionText:                     data.SizeOrDimensionText,
			ProductSpecification:                    data.ProductSpecification,
			MarkingOfMaterial:                       data.MarkingOfMaterial,
			BaseUnit:                                data.BaseUnit,
			DeliveryUnit:                            data.DeliveryUnit,
			OriginalQuantityInBaseUnit:              data.OrderQuantityInBaseUnit,
			OriginalQuantityInDeliveryUnit:          data.OrderQuantityInDeliveryUnit,
			StockConfirmationBusinessPartner:        *data.DeliverFromParty,
			StockConfirmationPlant:                  *data.DeliverFromPlant,
			InspectionPlantBusinessPartner:          data.InspectionPlantBusinessPartner,
			InspectionPlant:                         data.InspectionPlant,
			InspectionPlan:                          data.InspectionPlan,
			InspectionLot:                           data.InspectionLot,
			DeliveryDocumentItemText:                data.OrderItemText,
			DeliveryDocumentItemTextByBuyer:         data.OrderItemTextByBuyer,
			DeliveryDocumentItemTextBySeller:        data.OrderItemTextBySeller,
			PlannedGoodsIssueDate:                   data.RequestedDeliveryDate,
			PlannedGoodsIssueTime:                   "00:00:00",
			PlannedGoodsReceiptDate:                 data.RequestedDeliveryDate,
			PlannedGoodsReceiptTime:                 data.RequestedDeliveryTime,
			PlannedGoodsIssueQuantity:               data.OrderQuantityInDeliveryUnit,
			PlannedGoodsIssueQtyInBaseUnit:          data.OrderQuantityInBaseUnit,
			PlannedGoodsReceiptQuantity:             data.OrderQuantityInDeliveryUnit,
			PlannedGoodsReceiptQtyInBaseUnit:        data.OrderQuantityInBaseUnit,
			ItemCompleteDeliveryIsDefined:           &itemCompleteDeliveryIsDefined,
			ProductWeightUnit:                       data.ProductWeightUnit,
			ProductNetWeight:                        data.ProductNetWeight,
			OrderID:                                 orderID,
			OrderItem:                               orderItem,
			Contract:                                data.Contract,
			ContractItem:                            data.ContractItem,
			PaymentTerms:                            &data.PaymentTerms,
			Project:                                 data.Project,
			WBSElement:                              data.WBSElement,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			AccountAssignmentGroup:                  data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:           data.ProductAccountAssignmentGroup,
			ItemDeliveryBlockStatus:                 &itemDeliveryBlockStatus,
			CreationDate:                            formattedDate,
			CreationTime:                            formattedTime,
			LastChangeDate:                          formattedDate,
			LastChangeTime:                          formattedTime,
			IsCancelled:                             &isCancelled,
			IsMarkedForDeletion:                     &isMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &item, nil
}

func ConvertToPartnerFromOrders(
	rows *sql.Rows,
	deliveryDocumentIssuedID int,
) (*[]sub_func_complementer.Partner, error) {
	defer rows.Close()
	partner := make([]sub_func_complementer.Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersPartner{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
			&pm.EmailAddress,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm

		partner = append(partner, sub_func_complementer.Partner{
			DeliveryDocument:        deliveryDocumentIssuedID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
			EmailAddress:            data.EmailAddress,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &partner, nil
}

func ConvertToAddressFromOrders(
	rows *sql.Rows,
	deliveryDocumentIssuedID int,
) (*[]sub_func_complementer.Address, error) {
	defer rows.Close()
	address := make([]sub_func_complementer.Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersAddress{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm

		address = append(address, sub_func_complementer.Address{
			DeliveryDocument: deliveryDocumentIssuedID,
			AddressID:        data.AddressID,
			PostalCode:       data.PostalCode,
			LocalRegion:      data.LocalRegion,
			Country:          data.Country,
			District:         data.District,
			StreetName:       data.StreetName,
			CityName:         data.CityName,
			Building:         data.Building,
			Floor:            data.Floor,
			Room:             data.Room,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &address, nil
}
