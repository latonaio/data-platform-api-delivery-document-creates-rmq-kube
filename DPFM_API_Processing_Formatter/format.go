package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		SupplyChainRelationshipDeliveryID:      *data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID: *data.SupplyChainRelationshipDeliveryPlantID,
		DeliverToParty:                         *data.DeliverToParty,
		DeliverFromParty:                       *data.DeliverFromParty,
		DeliverToPlant:                         *data.DeliverToPlant,
		DeliverFromPlant:                       *data.DeliverFromPlant,
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
	}
}

func ConvertToAddressUpdates(address dpfm_api_input_reader.Address) *AddressUpdates {
	data := address

	return &AddressUpdates{
		PostalCode:  data.PostalCode,
		LocalRegion: data.LocalRegion,
		Country:     data.Country,
		District:    data.District,
		StreetName:  data.StreetName,
		CityName:    data.CityName,
		Building:    data.Building,
		Floor:       data.Floor,
		Room:        data.Room,
	}
}

func ConvertToItemUpdates(item dpfm_api_input_reader.Item) *ItemUpdates {
	data := item

	return &ItemUpdates{
		SupplyChainRelationshipStockConfPlantID:      *data.SupplyChainRelationshipStockConfPlantID,
		SupplyChainRelationshipProductionPlantID:     *data.SupplyChainRelationshipProductionPlantID,
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
	}
}

func ConvertToPartnerUpdates(partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	data := partner

	return &PartnerUpdates{
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		ExternalDocumentID:      data.ExternalDocumentID,
	}
}