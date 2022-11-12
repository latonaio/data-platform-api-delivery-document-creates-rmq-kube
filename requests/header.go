package requests

type Header struct {
	DeliveryDocumentstring              string `json:"DeliveryDocument"`
	Buyerstring                         string `json:"Buyer"`
	Sellerstring                        string `json:"Seller"`
	ContractTypestring                  string `json:"ContractType"`
	OrderValidityStartDatestring        string `json:"OrderValidityStartDate"`
	OrderValidityEndDatestring          string `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDatestring      string `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDatestring        string `json:"InvoiceScheduleEndDate"`
	IssuingLocationTimeZonestring       string `json:"IssuingLocationTimeZone"`
	ReceivingLocationTimeZonestring     string `json:"ReceivingLocationTimeZone"`
	DocumentDatestring                  string `json:"DocumentDate"`
	PlannedGoodsIssueDatestring         string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTimestring         string `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDatestring       string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTimestring       string `json:"PlannedGoodsReceiptTime"`
	BillingDocumentDatestring           string `json:"BillingDocumentDate"`
	CompleteDeliveryIsDefinedstring     string `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatusstring         string `json:"OverallDeliveryStatus"`
	CreationDatestring                  string `json:"CreationDate"`
	CreationTimestring                  string `json:"CreationTime"`
	IssuingBlockReasonstring            string `json:"IssuingBlockReason"`
	ReceivingBlockReasonstring          string `json:"ReceivingBlockReason"`
	GoodsIssueOrReceiptSlipNumberstring string `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatusstring           string `json:"HeaderBillingStatus"`
	HeaderBillingConfStatusstring       string `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockReasonstring      string `json:"HeaderBillingBlockReason"`
	HeaderGrossWeightstring             string `json:"HeaderGrossWeight"`
	HeaderNetWeightstring               string `json:"HeaderNetWeight"`
	HeaderVolumeUnitstring              string `json:"HeaderVolumeUnit"`
	HeaderWeightUnitstring              string `json:"HeaderWeightUnit"`
	Incotermsstring                     string `json:"Incoterms"`
	IsExportImportDeliverystring        string `json:"IsExportImportDelivery"`
	LastChangeDatestring                string `json:"LastChangeDate"`
	IssuingPlantBusinessPartnerstring   string `json:"IssuingPlantBusinessPartner"`
	IssuingPlantstring                  string `json:"IssuingPlant"`
	ReceivingPlantBusinessPartnerstring string `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlantstring                string `json:"ReceivingPlant"`
	DeliverToPartystring                string `json:"DeliverToParty"`
	DeliverFromPartystring              string `json:"DeliverFromParty"`
	TransactionCurrencystring           string `json:"TransactionCurrency"`
	OverallDelivReltdBillgStatusstring  string `json:"OverallDelivReltdBillgStatus"`
}
