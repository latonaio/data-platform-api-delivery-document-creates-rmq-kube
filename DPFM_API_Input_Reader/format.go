package dpfm_api_input_reader

import (
	"data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		DeliveryDocument:              data.DeliveryDocument,
		Buyer:                         data.Buyer,
		Seller:                        data.Seller,
		ReferenceDocument:             data.ReferenceDocument,
		ReferenceDocumentItem:         data.ReferenceDocumentItem,
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		ContractType:                  data.ContractType,
		OrderValidityStartDate:        data.OrderValidityStartDate,
		OrderValidityEndDate:          data.OrderValidityEndDate,
		InvoiceScheduleStartDate:      data.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:        data.InvoiceScheduleEndDate,
		IssuingPlantTimeZone:          data.IssuingPlantTimeZone,
		ReceivingPlantTimeZone:        data.ReceivingPlantTimeZone,
		DocumentDate:                  data.DocumentDate,
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:         data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:       data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:       data.PlannedGoodsReceiptTime,
		BillingDocumentDate:           data.BillingDocumentDate,
		CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
		OverallDeliveryStatus:         data.OverallDeliveryStatus,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		IssuingBlockReason:            data.IssuingBlockReason,
		ReceivingBlockReason:          data.ReceivingBlockReason,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		HeaderBillingStatus:           data.HeaderBillingStatus,
		HeaderBillingConfStatus:       data.HeaderBillingConfStatus,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		HeaderGrossWeight:             data.HeaderGrossWeight,
		HeaderNetWeight:               data.HeaderNetWeight,
		HeaderVolume:                  data.HeaderVolume,
		HeaderVolumeUnit:              data.HeaderVolumeUnit,
		HeaderWeightUnit:              data.HeaderWeightUnit,
		Incoterms:                     data.Incoterms,
		IsExportImportDelivery:        data.IsExportImportDelivery,
		LastChangeDate:                data.LastChangeDate,
		IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
		IssuingPlant:                  data.IssuingPlant,
		ReceivingPlant:                data.ReceivingPlant,
		ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
		DeliverToParty:                data.DeliverToParty,
		DeliverFromParty:              data.DeliverFromParty,
		TransactionCurrency:           data.TransactionCurrency,
		OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
	}
}
