package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-delivery-document-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) *HeaderCreates {
	data := subfuncSDC.Message.Header

	headerCreate := &HeaderCreates{
		DeliveryDocument:                data.DeliveryDocument,
		Buyer:                           data.Buyer,
		Seller:                          data.Seller,
		ReferenceDocument:               data.ReferenceDocument,
		ReferenceDocumentItem:           data.ReferenceDocumentItem,
		OrderID:                         data.OrderID,
		OrderItem:                       data.OrderItem,
		ContractType:                    data.ContractType,
		OrderValidityStartDate:          data.OrderValidityStartDate,
		OrderValidityEndDate:            data.OrderValidityEndDate,
		IssuingPlantTimeZone:            data.IssuingPlantTimeZone,
		ReceivingPlantTimeZone:          data.ReceivingPlantTimeZone,
		DocumentDate:                    data.DocumentDate,
		PlannedGoodsIssueDate:           data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:           data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:         data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:         data.PlannedGoodsReceiptTime,
		InvoiceDocumentDate:             data.InvoiceDocumentDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		CreationDate:                    data.CreationDate,
		CreationTime:                    data.CreationTime,
		HeaderDeliveryBlockStatus:       data.HeaderDeliveryBlockStatus,
		HeaderIssuingBlockStatus:        data.HeaderIssuingBlockStatus,
		HeaderReceivingBlockStatus:      data.HeaderReceivingBlockStatus,
		GoodsIssueOrReceiptSlipNumber:   data.GoodsIssueOrReceiptSlipNumber,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingConfStatus:         data.HeaderBillingConfStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		HeaderGrossWeight:               data.HeaderGrossWeight,
		HeaderNetWeight:                 data.HeaderNetWeight,
		HeaderWeightUnit:                data.HeaderWeightUnit,
		Incoterms:                       data.Incoterms,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		BillFromCountry:                 data.BillFromCountry,
		BillToCountry:                   data.BillToCountry,
		Payer:                           data.Payer,
		Payee:                           data.Payee,
		IsExportImportDelivery:          data.IsExportImportDelivery,
		LastChangeDate:                  data.LastChangeDate,
		IssuingPlantBusinessPartner:     data.IssuingPlantBusinessPartner,
		IssuingPlant:                    data.IssuingPlant,
		ReceivingPlantBusinessPartner:   data.ReceivingPlantBusinessPartner,
		ReceivingPlant:                  data.ReceivingPlant,
		DeliverToParty:                  data.DeliverToParty,
		DeliverFromParty:                data.DeliverFromParty,
		TransactionCurrency:             data.TransactionCurrency,
		StockIsFullyConfirmed:           data.StockIsFullyConfirmed,
	}

	return headerCreate
}

func ConvertToHeaderUpdates(headerUpdates dpfm_api_input_reader.HeaderUpdates) *HeaderUpdates {
	data := headerUpdates

	return &HeaderUpdates{
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:         data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:       data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:       data.PlannedGoodsReceiptTime,
		BillingDocumentDate:           data.BillingDocumentDate,
		HeaderIssuingBlockStatus:      data.HeaderIssuingBlockStatus,
		HeaderReceivingBlockStatus:    data.HeaderReceivingBlockStatus,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		HeaderBillingBlockStatus:      data.HeaderBillingBlockStatus,
		Incoterms:                     data.Incoterms,
		IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
		IssuingPlant:                  data.IssuingPlant,
		ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
		ReceivingPlant:                data.ReceivingPlant,
		DeliverFromParty:              data.DeliverFromParty,
		DeliverToParty:                data.DeliverToParty,
		StockIsFullyConfirmed:         data.StockIsFullyConfirmed,
	}
}

func ConvertToHeaderPartner(subfuncSDC *sub_func_complementer.SDC) []HeaderPartner {
	var headerPartner []HeaderPartner

	for _, data := range subfuncSDC.Message.HeaderPartner {
		headerPartner = append(headerPartner, HeaderPartner{
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

	return headerPartner
}

func ConvertToHeaderPartnerPlant(subfuncSDC *sub_func_complementer.SDC) []HeaderPartnerPlant {
	var headerPartnerPlant []HeaderPartnerPlant

	for _, data := range subfuncSDC.Message.HeaderPartnerPlant {
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			DeliveryDocument: data.DeliveryDocument,
			PartnerFunction:  data.PartnerFunction,
			BusinessPartner:  data.BusinessPartner,
			Plant:            data.Plant,
		})
	}

	return headerPartnerPlant
}

func ConvertToItem(subfuncSDC *sub_func_complementer.SDC) []Item {
	var item []Item

	for _, data := range subfuncSDC.Message.Item {
		item = append(item, Item{
			DeliveryDocument:                             data.DeliveryDocument,
			DeliveryDocumentItem:                         data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:                 data.DeliveryDocumentItemCategory,
			DeliveryDocumentItemText:                     data.DeliveryDocumentItemText,
			DeliveryDocumentItemTextByBuyer:              data.DeliveryDocumentItemTextByBuyer,
			DeliveryDocumentItemTextBySeller:             data.DeliveryDocumentItemTextBySeller,
			Product:                                      data.Product,
			ProductStandardID:                            data.ProductStandardID,
			ProductGroup:                                 data.ProductGroup,
			BaseUnit:                                     data.BaseUnit,
			OriginalQuantityInBaseUnit:                   data.OriginalQuantityInBaseUnit,
			IssuingUnit:                                  data.IssuingUnit,
			ReceivingUnit:                                data.ReceivingUnit,
			ActualGoodsIssueDate:                         data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                         data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                       data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                       data.ActualGoodsReceiptTime,
			ActualGoodsIssueQtyInBaseUnit:                data.ActualGoodsIssueQtyInBaseUnit,
			ActualGoodsIssueQuantity:                     data.ActualGoodsIssueQuantity,
			ActualGoodsReceiptQtyInBaseUnit:              data.ActualGoodsReceiptQtyInBaseUnit,
			ActualGoodsReceiptQuantity:                   data.ActualGoodsReceiptQuantity,
			CompleteItemDeliveryIsDefined:                data.CompleteItemDeliveryIsDefined,
			StockConfirmationPartnerFunction:             data.StockConfirmationPartnerFunction,
			StockConfirmationBusinessPartner:             data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                       data.StockConfirmationPlant,
			StockConfirmationPlantBatch:                  data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPolicy:                      data.StockConfirmationPolicy,
			StockConfirmationStatus:                      data.StockConfirmationStatus,
			ProductionPlantPartnerFunction:               data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner:               data.ProductionPlantBusinessPartner,
			ProductionPlant:                              data.ProductionPlant,
			ProductionPlantStorageLocation:               data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:                  data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:                  data.IssuingPlantBusinessPartner,
			IssuingPlant:                                 data.IssuingPlant,
			IssuingPlantStorageLocation:                  data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:                data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:                data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                               data.ReceivingPlant,
			ReceivingPlantStorageLocation:                data.ReceivingPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:       data.ProductIsBatchManagedInProductionPlant,
			ProductIsBatchManagedInIssuingPlant:          data.ProductIsBatchManagedInIssuingPlant,
			ProductIsBatchManagedInReceivingPlant:        data.ProductIsBatchManagedInReceivingPlant,
			BatchMgmtPolicyInProductionPlant:             data.BatchMgmtPolicyInProductionPlant,
			BatchMgmtPolicyInIssuingPlant:                data.BatchMgmtPolicyInIssuingPlant,
			BatchMgmtPolicyInReceivingPlant:              data.BatchMgmtPolicyInReceivingPlant,
			ProductionPlantBatch:                         data.ProductionPlantBatch,
			IssuingPlantBatch:                            data.IssuingPlantBatch,
			ReceivingPlantBatch:                          data.ReceivingPlantBatch,
			ProductionPlantBatchValidityStartDate:        data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:          data.ProductionPlantBatchValidityEndDate,
			IssuingPlantBatchValidityStartDate:           data.IssuingPlantBatchValidityStartDate,
			IssuingPlantBatchValidityEndDate:             data.IssuingPlantBatchValidityEndDate,
			ReceivingPlantBatchValidityStartDate:         data.ReceivingPlantBatchValidityStartDate,
			ReceivingPlantBatchValidityEndDate:           data.ReceivingPlantBatchValidityEndDate,
			CreationDate:                                 data.CreationDate,
			CreationTime:                                 data.CreationTime,
			ItemBillingStatus:                            data.ItemBillingStatus,
			ItemBillingConfStatus:                        data.ItemBillingConfStatus,
			SalesCostGLAccount:                           data.SalesCostGLAccount,
			ReceivingGLAccount:                           data.ReceivingGLAccount,
			IssuingGoodsMovementType:                     data.IssuingGoodsMovementType,
			ReceivingGoodsMovementType:                   data.ReceivingGoodsMovementType,
			ItemDeliveryBlockStatus:                      data.ItemDeliveryBlockStatus,
			ItemIssuingBlockStatus:                       data.ItemIssuingBlockStatus,
			ItemReceivingBlockStatus:                     data.ItemReceivingBlockStatus,
			ItemBillingBlockStatus:                       data.ItemBillingBlockStatus,
			ItemCompleteDeliveryIsDefined:                data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryIncompletionStatus:               data.ItemDeliveryIncompletionStatus,
			ItemGrossWeight:                              data.ItemGrossWeight,
			ItemNetWeight:                                data.ItemNetWeight,
			ItemWeightUnit:                               data.ItemWeightUnit,
			ItemIsBillingRelevant:                        data.ItemIsBillingRelevant,
			LastChangeDate:                               data.LastChangeDate,
			OrderID:                                      data.OrderID,
			OrderItem:                                    data.OrderItem,
			OrderType:                                    data.OrderType,
			ContractType:                                 data.ContractType,
			OrderValidityStartDate:                       data.OrderValidityStartDate,
			OrderValidityEndDate:                         data.OrderValidityEndDate,
			PaymentTerms:                                 data.PaymentTerms,
			DueCalculationBaseDate:                       data.DueCalculationBaseDate,
			PaymentDueDate:                               data.PaymentDueDate,
			NetPaymentDays:                               data.NetPaymentDays,
			PaymentMethod:                                data.PaymentMethod,
			InvoicePeriodStartDate:                       data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                         data.InvoicePeriodEndDate,
			ProductAvailabilityDate:                      data.ProductAvailabilityDate,
			Project:                                      data.Project,
			ReferenceDocument:                            data.ReferenceDocument,
			ReferenceDocumentItem:                        data.ReferenceDocumentItem,
			BPTaxClassification:                          data.BPTaxClassification,
			ProductTaxClassification:                     data.ProductTaxClassification,
			BPAccountAssignmentGroup:                     data.BPAccountAssignmentGroup,
			ProductAccountAssignmentGroup:                data.ProductAccountAssignmentGroup,
			TaxCode:                                      data.TaxCode,
			TaxRate:                                      data.TaxRate,
			CountryOfOrigin:                              data.CountryOfOrigin,
			StockIsFullyConfirmed:                        data.StockIsFullyConfirmed,
		})
	}

	return item
}
