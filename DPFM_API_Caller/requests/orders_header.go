package requests

type OrdersHeader struct {
	OrderID                          int      `json:"OrderID"`
	OrderDate                        string   `json:"OrderDate"`
	OrderType                        string   `json:"OrderType"`
	OrderStatus                      string   `json:"OrderStatus"`
	SupplyChainRelationshipID        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int      `json:"Buyer"`
	Seller                           int      `json:"Seller"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	BillToCountry                    *string  `json:"BillToCountry"`
	BillFromCountry                  *string  `json:"BillFromCountry"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	ContractType                     *string  `json:"ContractType"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32  `json:"TotalNetAmount"`
	TotalTaxAmount                   float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32  `json:"TotalGrossAmount"`
	HeaderDeliveryStatus             string   `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus              string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus         string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency              string   `json:"TransactionCurrency"`
	PricingDate                      string   `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime            string   `json:"RequestedDeliveryTime"`
	HeaderCompleteDeliveryIsDefined  *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     string   `json:"PaymentTerms"`
	PaymentMethod                    string   `json:"PaymentMethod"`
	Contract		                 *int     `json:"Contract"`
	ContractItem	                 *int     `json:"ContractItem"`
	ProductionVersion				 *int     `json:"ProductionVersion"`
	ProductionVersionItem			 *int     `json:"ProductionVersionItem"`
	ProductionOrder					 *int     `json:"ProductionOrder"`
	ProductionOrderItem				 *int     `json:"ProductionOrderItem"`
	Operations						 *int     `json:"Operations"`
	OperationsItem					 *int     `json:"OperationsItem"`
	OperationID						 *int     `json:"OperationID"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string   `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus"`
	ExternalReferenceDocument        *string  `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain        *string  `json:"CertificateAuthorityChain"`
	UsageControlChain        		 *string  `json:"UsageControlChain"`
	CreationDate                     string   `json:"CreationDate"`
	CreationTime                     string   `json:"CreationTime"`
	LastChangeDate                   string   `json:"LastChangeDate"`
	LastChangeTime                   string   `json:"LastChangeTime"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}
