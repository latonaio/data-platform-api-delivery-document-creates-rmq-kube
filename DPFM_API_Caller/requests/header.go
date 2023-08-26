package requests

type Header struct {
	DeliveryDocument                       int       `json:"DeliveryDocument"`
	SupplyChainRelationshipID              int       `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int       `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int       `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int      `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                  int       `json:"Buyer"`
	Seller                                 int       `json:"Seller"`
	DeliverToParty                         int       `json:"DeliverToParty"`
	DeliverFromParty                       int       `json:"DeliverFromParty"`
	DeliverToPlant                         string    `json:"DeliverToPlant"`
	DeliverFromPlant                       string    `json:"DeliverFromPlant"`
	BillToParty                            *int      `json:"BillToParty"`
	BillFromParty                          *int      `json:"BillFromParty"`
	BillToCountry                          *string   `json:"BillToCountry"`
	BillFromCountry                        *string   `json:"BillFromCountry"`
	Payer                                  *int      `json:"Payer"`
	Payee                                  *int      `json:"Payee"`
	IsExportImport                         *bool     `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string   `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string   `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int      `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int      `json:"ReferenceDocumentItem"`
	OrderID                                *int      `json:"OrderID"`
	OrderItem                              *int      `json:"OrderItem"`
	ProductionOrder                        *int      `json:"ProductionOrder"`
	ProductionOrderItem                    *int      `json:"ProductionOrderItem"`
	Operations                        	   *int      `json:"Operations"`
	OperationsItem                    	   *int      `json:"OperationsItem"`
	BillOfMaterial                    	   *int      `json:"BillOfMaterial"`
	BillOfMaterialItem                 	   *int      `json:"BillOfMaterialItem"`
	ContractType                           *string   `json:"ContractType"`
	OrderValidityStartDate                 *string   `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string   `json:"OrderValidityEndDate"`
	DeliveryDocumentDate                   string    `json:"DeliveryDocumentDate"`
	PlannedGoodsIssueDate                  string    `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  string    `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                string    `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                string    `json:"PlannedGoodsReceiptTime"`
	InvoiceDocumentDate                    *string   `json:"InvoiceDocumentDate"`
	HeaderCompleteDeliveryIsDefined        *bool     `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus                   *string   `json:"HeaderDeliveryStatus"`
	GoodsIssueOrReceiptSlipNumber          *string   `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus                    *string   `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus                *string   `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockStatus               *bool     `json:"HeaderBillingBlockStatus"`
	HeaderGrossWeight                      *float32  `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32  `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string   `json:"HeaderWeightUnit"`
	Incoterms                              *string   `json:"Incoterms"`
	TransactionCurrency                    *string   `json:"TransactionCurrency"`
	HeaderDeliveryBlockStatus              *bool     `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *bool     `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *bool     `json:"HeaderReceivingBlockStatus"`
	CreationDate                           string    `json:"CreationDate"`
	CreationTime                           string    `json:"CreationTime"`
	LastChangeDate                         string    `json:"LastChangeDate"`
	LastChangeTime                         string    `json:"LastChangeTime"`
	IsCancelled                            *bool     `json:"IsCancelled"`
	IsMarkedForDeletion                    *bool     `json:"IsMarkedForDeletion"`
}
