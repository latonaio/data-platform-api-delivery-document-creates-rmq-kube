package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		DeliveryDocument:                       data.DeliveryDocument,
		DeliveryDocumentDate:      				data.DeliveryDocumentDate,
		PlannedGoodsIssueDate:                  data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:                  data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:                data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:                data.PlannedGoodsReceiptTime,
		InvoiceDocumentDate:                    data.InvoiceDocumentDate,
		HeaderBillingBlockStatus:               data.HeaderBillingBlockStatus,
		GoodsIssueOrReceiptSlipNumber:          data.GoodsIssueOrReceiptSlipNumber,
		Incoterms:                              data.Incoterms,
		HeaderDeliveryBlockStatus:              data.HeaderDeliveryBlockStatus,
		HeaderIssuingBlockStatus:               data.HeaderIssuingBlockStatus,
		HeaderReceivingBlockStatus:             data.HeaderReceivingBlockStatus,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &ItemUpdates{
		DeliveryDocument:					dataHeader.DeliveryDocument,
		DeliveryDocumentItem:				data.DeliveryDocumentItem,
		DeliverToPlantStorageLocation:		data.DeliverToPlantStorageLocation,
		DeliverFromPlantStorageLocation:	data.DeliverFromPlantStorageLocation,
		PlannedGoodsIssueDate:				data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:				data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:			data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:			data.PlannedGoodsReceiptTime,
		PlannedGoodsIssueQuantity:			data.PlannedGoodsIssueQuantity,
		PlannedGoodsIssueQtyInBaseUnit:		data.PlannedGoodsIssueQtyInBaseUnit,
		PlannedGoodsReceiptQuantity:		data.PlannedGoodsReceiptQuantity,
		PlannedGoodsReceiptQtyInBaseUnit:	data.PlannedGoodsReceiptQtyInBaseUnit,
		ActualGoodsIssueDate:				data.ActualGoodsIssueDate,
		ActualGoodsIssueTime:				data.ActualGoodsIssueTime,
		ActualGoodsReceiptDate:				data.ActualGoodsReceiptDate,
		ActualGoodsReceiptTime:				data.ActualGoodsReceiptTime,
		ActualGoodsIssueQuantity:			data.ActualGoodsIssueQuantity,
		ActualGoodsIssueQtyInBaseUnit:		data.ActualGoodsIssueQtyInBaseUnit,
		ActualGoodsReceiptQuantity:			data.ActualGoodsReceiptQuantity,
		ActualGoodsReceiptQtyInBaseUnit:	data.ActualGoodsReceiptQtyInBaseUnit,
		ItemGrossWeight:              data.ItemGrossWeight,
		ItemNetWeight:                data.ItemNetWeight,
		ItemWeightUnit:               data.ItemWeightUnit,
		DueCalculationBaseDate:       data.DueCalculationBaseDate,
		PaymentDueDate:               data.PaymentDueDate,
		NetPaymentDays:               data.NetPaymentDays,
		ItemDeliveryBlockStatus:      data.ItemDeliveryBlockStatus,
		ItemIssuingBlockStatus:       data.ItemIssuingBlockStatus,
		ItemReceivingBlockStatus:     data.ItemReceivingBlockStatus,
		ItemBillingBlockStatus:       data.ItemBillingBlockStatus,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		DeliveryDocument:        dataHeader.DeliveryDocument,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		ExternalDocumentID:      data.ExternalDocumentID,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		DeliveryDocument: dataHeader.DeliveryDocument,
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
	}
}
