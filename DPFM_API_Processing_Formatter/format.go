package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-delivery-document-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		DeliveryDocument:                 data.DeliveryDocument,
		SupplyChainRelationshipFreightID: data.SupplyChainRelationshipFreightID,
		FreightPartner:                   data.FreightPartner,
		PlannedGoodsIssueDate:            data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:            data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:          data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:          data.PlannedGoodsReceiptTime,
		FreightOrder:                     data.FreightOrder,
		InvoiceDocumentDate:              data.InvoiceDocumentDate,
		HeaderBillingBlockStatus:         data.HeaderBillingBlockStatus,
		GoodsIssueOrReceiptSlipNumber:    data.GoodsIssueOrReceiptSlipNumber,
		Incoterms:                        data.Incoterms,
		HeaderDeliveryBlockStatus:        data.HeaderDeliveryBlockStatus,
		HeaderIssuingBlockStatus:         data.HeaderIssuingBlockStatus,
		HeaderReceivingBlockStatus:       data.HeaderReceivingBlockStatus,
		ExternalReferenceDocument:        data.ExternalReferenceDocument,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &ItemUpdates{
		DeliveryDocument:               dataHeader.DeliveryDocument,
		DeliveryDocumentItem:           data.DeliveryDocumentItem,
		PlannedGoodsIssueDate:          data.PlannedGoodsIssueDate,
		PlannedGoodsIssueTime:          data.PlannedGoodsIssueTime,
		PlannedGoodsReceiptDate:        data.PlannedGoodsReceiptDate,
		PlannedGoodsReceiptTime:        data.PlannedGoodsReceiptTime,
		PlannedGoodsIssueQuantity:      data.PlannedGoodsIssueQuantity,
		PlannedGoodsIssueQtyInBaseUnit: data.PlannedGoodsIssueQtyInBaseUnit,
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
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
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

//func ConvertToItemPickingUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.ItemPicking) *ItemPickingUpdates {
//	dataHeader := header
//	data := itemPicking
//
//	return &ItemPickingUpdates{
//		DeliveryDocument:                                 dataHeader.DeliveryDocument,
//		DeliveryDocumentItem:                             data.DeliveryDocumentItem,
//		DeliveryDocumentItemPickingID:                    data.DeliveryDocumentItemPickingID,
//		DeliverToPlantStorageBin:                         data.DeliverToPlantStorageBin,
//		DeliverToPlantKanbanContainer:                    data.DeliverToPlantKanbanContainer,
//		DeliverFromPlantStorageBin:                       data.DeliverFromPlantStorageBin,
//		DeliverFromPlantKanbanContainer:                  data.DeliverFromPlantKanbanContainer,
//		DeliverToPlantPlannedPickingQuantityInBaseUnit:   data.DeliverToPlantPlannedPickingQuantityInBaseUnit,
//		DeliverFromPlantPlannedPickingQuantityInBaseUnit: data.DeliverFromPlantPlannedPickingQuantityInBaseUnit,
//		DeliverToPlantPlannedPickingDate:                 data.DeliverToPlantPlannedPickingDate,
//		DeliverToPlantPlannedPickingTime:                 data.DeliverToPlantPlannedPickingTime,
//		DeliverFromPlantPlannedPickingDate:               data.DeliverFromPlantPlannedPickingDate,
//		DeliverFromPlantPlannedPickingTime:               data.DeliverFromPlantPlannedPickingTime,
//		DeliverToPlantActualPickingQuantityInBaseUnit:    data.DeliverToPlantActualPickingQuantityInBaseUnit,
//		DeliverToPlantActualPickingDate:                  data.DeliverToPlantActualPickingDate,
//		DeliverToPlantActualPickingTime:                  data.DeliverToPlantActualPickingTime,
//		DeliverFromPlantActualPickingQuantityInBaseUnit:  data.DeliverFromPlantActualPickingQuantityInBaseUnit,
//		DeliverFromPlantActualPickingDate:                data.DeliverFromPlantActualPickingDate,
//		DeliverFromPlantActualPickingTime:                data.DeliverFromPlantActualPickingTime,
//		ExternalReferenceDocument:        				  data.ExternalReferenceDocument,
//		ExternalReferenceDocumentItem:    				  data.ExternalReferenceDocumentItem,
//		ExternalReferenceDocumentItemPickingID:			  data.ExternalReferenceDocumentItemPickingID,
//	}
//}
