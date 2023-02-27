package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-delivery-document-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) *[]Header {
	var headers []Header

	for _, data := range *subfuncSDC.Message.Header {
		headers = append(headers, Header{
			DeliveryDocument:                       data.DeliveryDocument,
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipBillingID:       data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:       data.SupplyChainRelationshipPaymentID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			BillToParty:                            data.BillToParty,
			BillFromParty:                          data.BillFromParty,
			BillToCountry:                          data.BillToCountry,
			BillFromCountry:                        data.BillFromCountry,
			Payer:                                  data.Payer,
			Payee:                                  data.Payee,
			IsExportImport:                         data.IsExportImport,
			DeliverToPlantTimeZone:                 data.DeliverToPlantTimeZone,
			DeliverFromPlantTimeZone:               data.DeliverFromPlantTimeZone,
			ReferenceDocument:                      data.ReferenceDocument,
			ReferenceDocumentItem:                  data.ReferenceDocumentItem,
			OrderID:                                data.OrderID,
			OrderItem:                              data.OrderItem,
			ContractType:                           data.ContractType,
			OrderValidityStartDate:                 data.OrderValidityStartDate,
			OrderValidityEndDate:                   data.OrderValidityEndDate,
			DocumentDate:                           data.DocumentDate,
			PlannedGoodsIssueDate:                  data.PlannedGoodsIssueDate,
			PlannedGoodsIssueTime:                  data.PlannedGoodsIssueTime,
			PlannedGoodsReceiptDate:                data.PlannedGoodsReceiptDate,
			PlannedGoodsReceiptTime:                data.PlannedGoodsReceiptTime,
			InvoiceDocumentDate:                    data.InvoiceDocumentDate,
			HeaderCompleteDeliveryIsDefined:        data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:                   data.HeaderDeliveryStatus,
			CreationDate:                           data.CreationDate,
			CreationTime:                           data.CreationTime,
			LastChangeDate:                         data.LastChangeDate,
			LastChangeTime:                         data.LastChangeTime,
			GoodsIssueOrReceiptSlipNumber:          data.GoodsIssueOrReceiptSlipNumber,
			HeaderBillingStatus:                    data.HeaderBillingStatus,
			HeaderBillingConfStatus:                data.HeaderBillingConfStatus,
			HeaderBillingBlockStatus:               data.HeaderBillingBlockStatus,
			HeaderGrossWeight:                      data.HeaderGrossWeight,
			HeaderNetWeight:                        data.HeaderNetWeight,
			HeaderWeightUnit:                       data.HeaderWeightUnit,
			Incoterms:                              data.Incoterms,
			TransactionCurrency:                    data.TransactionCurrency,
			HeaderDeliveryBlockStatus:              data.HeaderDeliveryBlockStatus,
			HeaderIssuingBlockStatus:               data.HeaderIssuingBlockStatus,
			HeaderReceivingBlockStatus:             data.HeaderReceivingBlockStatus,
			IsCancelled:                            data.IsCancelled,
			IsMarkedForDeletion:                    data.IsMarkedForDeletion,
		})
	}

	return &headers
}

func ConvertToHeaderUpdates(headerUpdates *[]dpfm_api_processing_formatter.HeaderUpdates) *[]Header {
	var header []Header

	for _, data := range *headerUpdates {
		header = append(header, Header{
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			DocumentDate:                           data.DocumentDate,
			PlannedGoodsIssueDate:                  data.PlannedGoodsIssueDate,
			PlannedGoodsIssueTime:                  data.PlannedGoodsIssueTime,
			PlannedGoodsReceiptDate:                data.PlannedGoodsReceiptDate,
			PlannedGoodsReceiptTime:                data.PlannedGoodsReceiptTime,
			InvoiceDocumentDate:                    data.InvoiceDocumentDate,
			GoodsIssueOrReceiptSlipNumber:          data.GoodsIssueOrReceiptSlipNumber,
			HeaderBillingBlockStatus:               data.HeaderBillingBlockStatus,
			Incoterms:                              data.Incoterms,
			HeaderDeliveryBlockStatus:              data.HeaderDeliveryBlockStatus,
			HeaderIssuingBlockStatus:               data.HeaderIssuingBlockStatus,
			HeaderReceivingBlockStatus:             data.HeaderReceivingBlockStatus,
		})
	}

	return &header
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) *[]Partner {
	var partner []Partner

	for _, data := range *subfuncSDC.Message.Partner {
		partner = append(partner, Partner{
			DeliveryDocument:        data.DeliveryDocument,
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
		})
	}

	return &partner
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) *[]Partner {
	var partner []Partner

	for _, data := range *partnerUpdates {
		partner = append(partner, Partner{
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			ExternalDocumentID:      data.ExternalDocumentID,
		})
	}

	return &partner
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) *[]Item {
	var item []Item

	for _, data := range *subfuncSDC.Message.Item {
		item = append(item, Item{
			DeliveryDocument:                              data.DeliveryDocument,
			DeliveryDocumentItem:                          data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:                  data.DeliveryDocumentItemCategory,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			SupplyChainRelationshipBillingID:              data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:              data.SupplyChainRelationshipPaymentID,
			Buyer:                                         data.Buyer,
			Seller:                                        data.Seller,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverFromPlant:                              data.DeliverFromPlant,
			BillToParty:                                   data.BillToParty,
			BillFromParty:                                 data.BillFromParty,
			BillToCountry:                                 data.BillToCountry,
			BillFromCountry:                               data.BillFromCountry,
			Payer:                                         data.Payer,
			Payee:                                         data.Payee,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityStartTime:          data.DeliverToPlantBatchValidityStartTime,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverToPlantBatchValidityEndTime:            data.DeliverToPlantBatchValidityEndTime,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityStartTime:        data.DeliverFromPlantBatchValidityStartTime,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliverFromPlantBatchValidityEndTime:          data.DeliverFromPlantBatchValidityEndTime,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyInStockConfirmationPlant:       data.BatchMgmtPolicyInStockConfirmationPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityStartTime:  data.StockConfirmationPlantBatchValidityStartTime,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPlantBatchValidityEndTime:    data.StockConfirmationPlantBatchValidityEndTime,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityStartTime:         data.ProductionPlantBatchValidityStartTime,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			ProductionPlantBatchValidityEndTime:           data.ProductionPlantBatchValidityEndTime,
			DeliveryDocumentItemText:                      data.DeliveryDocumentItemText,
			DeliveryDocumentItemTextByBuyer:               data.DeliveryDocumentItemTextByBuyer,
			DeliveryDocumentItemTextBySeller:              data.DeliveryDocumentItemTextBySeller,
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			OriginalQuantityInBaseUnit:                    data.OriginalQuantityInBaseUnit,
			DeliveryUnit:                                  data.DeliveryUnit,
			PlannedGoodsIssueDate:                         data.PlannedGoodsIssueDate,
			PlannedGoodsIssueTime:                         data.PlannedGoodsIssueTime,
			PlannedGoodsReceiptDate:                       data.PlannedGoodsReceiptDate,
			PlannedGoodsReceiptTime:                       data.PlannedGoodsReceiptTime,
			PlannedGoodsIssueQuantity:                     data.PlannedGoodsIssueQuantity,
			PlannedGoodsIssueQtyInBaseUnit:                data.PlannedGoodsIssueQtyInBaseUnit,
			PlannedGoodsReceiptQuantity:                   data.PlannedGoodsReceiptQuantity,
			PlannedGoodsReceiptQtyInBaseUnit:              data.PlannedGoodsReceiptQtyInBaseUnit,
			ActualGoodsIssueDate:                          data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                          data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                        data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                        data.ActualGoodsReceiptTime,
			ActualGoodsIssueQuantity:                      data.ActualGoodsIssueQuantity,
			ActualGoodsIssueQtyInBaseUnit:                 data.ActualGoodsIssueQtyInBaseUnit,
			ActualGoodsReceiptQuantity:                    data.ActualGoodsReceiptQuantity,
			ActualGoodsReceiptQtyInBaseUnit:               data.ActualGoodsReceiptQtyInBaseUnit,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			ItemBillingStatus:                             data.ItemBillingStatus,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			ItemWeightUnit:                                data.ItemWeightUnit,
			InternalCapacityQuantity:                      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:                  data.InternalCapacityQuantityUnit,
			ItemIsBillingRelevant:                         data.ItemIsBillingRelevant,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			OrderType:                                     data.OrderType,
			ContractType:                                  data.ContractType,
			OrderValidityStartDate:                        data.OrderValidityStartDate,
			OrderValidityEndDate:                          data.OrderValidityEndDate,
			PaymentTerms:                                  data.PaymentTerms,
			DueCalculationBaseDate:                        data.DueCalculationBaseDate,
			PaymentDueDate:                                data.PaymentDueDate,
			NetPaymentDays:                                data.NetPaymentDays,
			PaymentMethod:                                 data.PaymentMethod,
			InvoicePeriodStartDate:                        data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                          data.InvoicePeriodEndDate,
			ConfirmedDeliveryDate:                         data.ConfirmedDeliveryDate,
			Project:                                       data.Project,
			ReferenceDocument:                             data.ReferenceDocument,
			ReferenceDocumentItem:                         data.ReferenceDocumentItem,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemIssuingBlockStatus:                        data.ItemIssuingBlockStatus,
			ItemReceivingBlockStatus:                      data.ItemReceivingBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
		})
	}

	return &item
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) *[]Item {
	var item []Item

	for _, data := range *itemUpdates {
		item = append(item, Item{
			SupplyChainRelationshipStockConfPlantID:      data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:     data.SupplyChainRelationshipProductionPlantID,
			DeliverToPlantStorageLocation:                data.DeliverToPlantStorageLocation,
			DeliverToPlantBatch:                          data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:         data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityEndDate:           data.DeliverToPlantBatchValidityEndDate,
			DeliverFromPlantStorageLocation:              data.DeliverFromPlantStorageLocation,
			DeliverFromPlantBatch:                        data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:       data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityEndDate:         data.DeliverFromPlantBatchValidityEndDate,
			StockConfirmationBusinessPartner:             data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                       data.StockConfirmationPlant,
			StockConfirmationPlantBatch:                  data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPolicy:                      data.StockConfirmationPolicy,
			ProductionPlantBusinessPartner:               data.ProductionPlantBusinessPartner,
			ProductionPlant:                              data.ProductionPlant,
			ProductionPlantStorageLocation:               data.ProductionPlantStorageLocation,
			ProductionPlantBatch:                         data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:        data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:          data.ProductionPlantBatchValidityEndDate,
			ActualGoodsIssueDate:                         data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                         data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                       data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                       data.ActualGoodsReceiptTime,
			ActualGoodsIssueQuantity:                     data.ActualGoodsIssueQuantity,
			ActualGoodsIssueQtyInBaseUnit:                data.ActualGoodsIssueQtyInBaseUnit,
			ActualGoodsReceiptQuantity:                   data.ActualGoodsReceiptQuantity,
			ActualGoodsReceiptQtyInBaseUnit:              data.ActualGoodsReceiptQtyInBaseUnit,
			ItemGrossWeight:                              data.ItemGrossWeight,
			ItemNetWeight:                                data.ItemNetWeight,
			ItemWeightUnit:                               data.ItemWeightUnit,
			ItemIsBillingRelevant:                        data.ItemIsBillingRelevant,
			DueCalculationBaseDate:                       data.DueCalculationBaseDate,
			PaymentDueDate:                               data.PaymentDueDate,
			NetPaymentDays:                               data.NetPaymentDays,
			ItemDeliveryBlockStatus:                      data.ItemDeliveryBlockStatus,
			ItemIssuingBlockStatus:                       data.ItemIssuingBlockStatus,
			ItemReceivingBlockStatus:                     data.ItemReceivingBlockStatus,
			ItemBillingBlockStatus:                       data.ItemBillingBlockStatus,
		})
	}

	return &item
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) *[]Address {
	var address []Address

	for _, data := range *subfuncSDC.Message.Address {
		address = append(address, Address{
			DeliveryDocument: data.DeliveryDocument,
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

	return &address
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) *[]Address {
	var address []Address

	for _, data := range *addressUpdates {
		address = append(address, Address{
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}

	return &address
}
