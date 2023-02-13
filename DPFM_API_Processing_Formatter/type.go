package dpfm_api_processing_formatter

type HeaderUpdates struct {
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	DocumentDate                           *string `json:"DocumentDate"`
	PlannedGoodsIssueDate                  *string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string `json:"PlannedGoodsReceiptTime"`
	InvoiceDocumentDate                    *string `json:"InvoiceDocumentDate"`
	GoodsIssueOrReceiptSlipNumber          *string `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingBlockStatus               *bool   `json:"HeaderBillingBlockStatus"`
	Incoterms                              *string `json:"Incoterms"`
	HeaderDeliveryBlockStatus              *bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *bool   `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *bool   `json:"HeaderReceivingBlockStatus"`
}

type PartnerUpdates struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}

type AddressUpdates struct {
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type ItemUpdates struct {
	SupplyChainRelationshipStockConfPlantID      int      `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID     int      `json:"SupplyChainRelationshipProductionPlantID"`
	DeliverToPlantStorageLocation                *string  `json:"DeliverToPlantStorageLocation"`
	DeliverToPlantBatch                          *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate         *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityEndDate           *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverFromPlantStorageLocation              *string  `json:"DeliverFromPlantStorageLocation"`
	DeliverFromPlantBatch                        *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate       *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityEndDate         *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	StockConfirmationBusinessPartner             *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPolicy                      *string  `json:"StockConfirmationPolicy"`
	ProductionPlantBusinessPartner               *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                              *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation               *string  `json:"ProductionPlantStorageLocation"`
	ProductionPlantBatch                         *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate        *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate          *string  `json:"ProductionPlantBatchValidityEndDate"`
	ActualGoodsIssueDate                         *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                         *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                       *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                       *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity                     *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit                *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                   *float32 `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit              *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ItemGrossWeight                              *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                                *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                               *string  `json:"ItemWeightUnit"`
	ItemIsBillingRelevant                        *bool    `json:"ItemIsBillingRelevant"`
	DueCalculationBaseDate                       *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                               *string  `json:"PaymentDueDate"`
	NetPaymentDays                               *int     `json:"NetPaymentDays"`
	ItemDeliveryBlockStatus                      *bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                       *bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                     *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                       *bool    `json:"ItemBillingBlockStatus"`
}
