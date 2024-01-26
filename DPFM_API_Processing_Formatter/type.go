package dpfm_api_processing_formatter

type HeaderUpdates struct {
	DeliveryDocument                 int     `json:"DeliveryDocument"`
	SupplyChainRelationshipFreightID *int    `json:"SupplyChainRelationshipFreightID"`
	FreightPartner                   *int    `json:"FreightPartner"`
	PlannedGoodsIssueDate            *string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime            *string `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate          *string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime          *string `json:"PlannedGoodsReceiptTime"`
	FreightOrder                     *int    `json:"FreightOrder"`
	InvoiceDocumentDate              *string `json:"InvoiceDocumentDate"`
	HeaderBillingBlockStatus         *bool   `json:"HeaderBillingBlockStatus"`
	GoodsIssueOrReceiptSlipNumber    *string `json:"GoodsIssueOrReceiptSlipNumber"`
	Incoterms                        *string `json:"Incoterms"`
	HeaderDeliveryBlockStatus        *bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus         *bool   `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus       *bool   `json:"HeaderReceivingBlockStatus"`
	ExternalReferenceDocument        *string `json:"ExternalReferenceDocument"`
}

type ItemUpdates struct {
	DeliveryDocument               int      `json:"DeliveryDocument"`
	DeliveryDocumentItem           int      `json:"DeliveryDocumentItem"`
	PlannedGoodsIssueDate          *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime          *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate        *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime        *string  `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity      *float32 `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit *float32 `json:"PlannedGoodsIssueQtyInBaseUnit"`
}

type ItemPickingUpdates struct {
	DeliveryDocument                                 int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                             int      `json:"DeliverToPlantStorageBin"`
	DeliveryDocumentItemPickingID                    int      `json:"DeliverFromPlantStorageBin"`
	DeliverToPlantStorageBin                         *string  `json:"DeliverToPlantStorageBin"`
	DeliverToPlantKanbanContainer                    *int     `json:"DeliverToPlantKanbanContainer"`
	DeliverFromPlantStorageBin                       *string  `json:"DeliverFromPlantStorageBin"`
	DeliverFromPlantKanbanContainer                  *int     `json:"DeliverFromPlantKanbanContainer"`
	DeliverToPlantPlannedPickingQuantityInBaseUnit   *float32 `json:"DeliverToPlantPlannedPickingQuantityInBaseUnit"`
	DeliverFromPlantPlannedPickingQuantityInBaseUnit *float32 `json:"DeliverFromPlantPlannedPickingQuantityInBaseUnit"`
	DeliverToPlantPlannedPickingDate                 *string  `json:"DeliverToPlantPlannedPickingDate"`
	DeliverToPlantPlannedPickingTime                 *string  `json:"DeliverToPlantPlannedPickingTime"`
	DeliverFromPlantPlannedPickingDate               *string  `json:"DeliverFromPlantPlannedPickingDate"`
	DeliverFromPlantPlannedPickingTime               *string  `json:"DeliverFromPlantPlannedPickingTime"`
	DeliverToPlantActualPickingQuantityInBaseUnit    *float32 `json:"DeliverToPlantActualPickingQuantityInBaseUnit"`
	DeliverToPlantActualPickingDate                  *string  `json:"DeliverToPlantActualPickingDate"`
	DeliverToPlantActualPickingTime                  *string  `json:"DeliverToPlantActualPickingTime"`
	DeliverFromPlantActualPickingQuantityInBaseUnit  *float32 `json:"DeliverFromPlantActualPickingQuantityInBaseUnit"`
	DeliverFromPlantActualPickingDate                *string  `json:"DeliverFromPlantActualPickingDate"`
	DeliverFromPlantActualPickingTime                *string  `json:"DeliverFromPlantActualPickingTime"`
	ExternalReferenceDocument                        *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                    *string  `json:"ExternalReferenceDocumentItem"`
	ExternalReferenceDocumentItemPickingID           *string  `json:"ExternalReferenceDocumentItemPickingID"`
}

type PartnerUpdates struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type AddressUpdates struct {
	DeliveryDocument int     `json:"DeliveryDocument"`
	AddressID        int     `json:"AddressID"`
	PostalCode       *string `json:"PostalCode"`
	LocalRegion      *string `json:"LocalRegion"`
	Country          *string `json:"Country"`
	District         *string `json:"District"`
	StreetName       *string `json:"StreetName"`
	CityName         *string `json:"CityName"`
	Building         *string `json:"Building"`
	Floor            *int    `json:"Floor"`
	Room             *int    `json:"Room"`
}
