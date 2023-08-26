package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string          `json:"connection_key"`
	Result           bool            `json:"result"`
	RedisKey         string          `json:"redis_key"`
	Filepath         string          `json:"filepath"`
	APIStatusCode    int             `json:"api_status_code"`
	RuntimeSessionID string          `json:"runtime_session_id"`
	BusinessPartner  *int            `json:"business_partner"`
	ServiceLabel     string          `json:"service_label"`
	APIType          string          `json:"api_type"`
	InputParameters  InputParameters `json:"InputParameters"`
	Header           Header          `json:"DeliveryDocument"`
	APISchema        string          `json:"api_schema"`
	Accepter         []string        `json:"accepter"`
	Deleted          bool            `json:"deleted"`
}

type InputParameters struct {
	DeliverToParty            *[]*int    `json:"DeliverToParty"`
	DeliverToPartyTo          *int       `json:"DeliverToPartyTo"`
	DeliverToPartyFrom        *int       `json:"DeliverToPartyFrom"`
	DeliverFromParty          *[]*int    `json:"DeliverFromParty"`
	DeliverFromPartyTo        *int       `json:"DeliverFromPartyTo"`
	DeliverFromPartyFrom      *int       `json:"DeliverFromPartyFrom"`
	DeliverToPlant            *[]*string `json:"DeliverToPlant"`
	DeliverToPlantTo          *string    `json:"DeliverToPlantTo"`
	DeliverToPlantFrom        *string    `json:"DeliverToPlantFrom"`
	DeliverFromPlant          *[]*string `json:"DeliverFromPlant"`
	DeliverFromPlantTo        *string    `json:"DeliverFromPlantTo"`
	DeliverFromPlantFrom      *string    `json:"DeliverFromPlantFrom"`
	ConfirmedDeliveryDate     *[]*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateTo   *string    `json:"ConfirmedDeliveryDateTo"`
	ConfirmedDeliveryDateFrom *string    `json:"ConfirmedDeliveryDateFrom"`
}

type Header struct {
	SupplyChainRelationshipID              *int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      *int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID *int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int      `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                  *int      `json:"Buyer"`
	Seller                                 *int      `json:"Seller"`
	DeliverToParty                         *int      `json:"DeliverToParty"`
	DeliverFromParty                       *int      `json:"DeliverFromParty"`
	DeliverToPlant                         *string   `json:"DeliverToPlant"`
	DeliverFromPlant                       *string   `json:"DeliverFromPlant"`
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
	DeliveryDocumentDate                   *string   `json:"DeliveryDocumentDate"`
	PlannedGoodsIssueDate                  *string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string   `json:"PlannedGoodsReceiptTime"`
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
	CreationDate                           *string   `json:"CreationDate"`
	CreationTime                           *string   `json:"CreationTime"`
	LastChangeDate                         *string   `json:"LastChangeDate"`
	LastChangeTime                         *string   `json:"LastChangeTime"`
	IsCancelled                            *bool     `json:"IsCancelled"`
	IsMarkedForDeletion                    *bool     `json:"IsMarkedForDeletion"`
	Item                                   []Item    `json:"Item"`
	Partner                                []Partner `json:"Partner"`
	Address                                []Address `json:"Address"`
}

type Partner struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
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

type Item struct {
	DeliveryDocument                              int           `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int           `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  *string       `json:"DeliveryDocumentItemCategory"`
	SupplyChainRelationshipID                     *int          `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int          `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int          `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int          `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int          `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipBillingID              *int          `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID              *int          `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                         *int          `json:"Buyer"`
	Seller                                        *int          `json:"Seller"`
	DeliverToParty                                *int          `json:"DeliverToParty"`
	DeliverFromParty                              *int          `json:"DeliverFromParty"`
	DeliverToPlant                                *string       `json:"DeliverToPlant"`
	DeliverFromPlant                              *string       `json:"DeliverFromPlant"`
	BillToParty                                   *int          `json:"BillToParty"`
	BillFromParty                                 *int          `json:"BillFromParty"`
	BillToCountry                                 *string       `json:"BillToCountry"`
	BillFromCountry                               *string       `json:"BillFromCountry"`
	Payer                                         *int          `json:"Payer"`
	Payee                                         *int          `json:"Payee"`
	Product                                       *string       `json:"Product"`
	ProductStandardID                             *string       `json:"ProductStandardID"`
	ProductGroup                                  *string       `json:"ProductGroup"`
	BaseUnit                                      *string       `json:"BaseUnit"`
	DeliveryUnit                                  *string       `json:"DeliveryUnit"`
	OriginalQuantityInBaseUnit                    *float32      `json:"OriginalQuantityInBaseUnit"`
	OriginalQuantityInDeliveryUnit                *float32      `json:"OriginalQuantityInDeliveryUnit"`
	DeliverToPlantStorageLocation                 *string       `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool         `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string       `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string       `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string       `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string       `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string       `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string       `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlantStorageLocation               *string       `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool         `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string       `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string       `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string       `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string       `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string       `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string       `json:"DeliverFromPlantBatchValidityEndTime"`
	StockConfirmationBusinessPartner              *int          `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string       `json:"StockConfirmationPlant"`
	ProductIsBatchManagedInStockConfirmationPlant *bool         `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string       `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string       `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string       `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string       `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string       `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string       `json:"StockConfirmationPlantBatchValidityEndTime"`
	StockConfirmationPolicy                       *string       `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string       `json:"StockConfirmationStatus"`
	ProductionPlantBusinessPartner                *int          `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string       `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string       `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool         `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string       `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string       `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string       `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime         *string       `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate           *string       `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime           *string       `json:"ProductionPlantBatchValidityEndTime"`
	InspectionPlan								  *int	   		`json:"InspectionPlan"`
	InspectionPlant								  *string  		`json:"InspectionPlant"`
	InspectionOrder								  *int	   		`json:"InspectionOrder"`
	DeliveryDocumentItemText                      *string       `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               *string       `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              *string       `json:"DeliveryDocumentItemTextBySeller"`
	PlannedGoodsIssueDate                         *string       `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                         *string       `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                       *string       `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                       *string       `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity                     *float32      `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit                *float32      `json:"PlannedGoodsIssueQtyInBaseUnit"`
	PlannedGoodsReceiptQuantity                   *float32      `json:"PlannedGoodsReceiptQuantity"`
	PlannedGoodsReceiptQtyInBaseUnit              *float32      `json:"PlannedGoodsReceiptQtyInBaseUnit"`
	ActualGoodsIssueDate                          *string       `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                          *string       `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                        *string       `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                        *string       `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity                      *float32      `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit                 *float32      `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                    *float32      `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit               *float32      `json:"ActualGoodsReceiptQtyInBaseUnit"`
	QuantityPerPackage                            *float32      `json:"QuantityPerPackage"`
	ItemBillingStatus                             *string       `json:"ItemBillingStatus"`
	ItemCompleteDeliveryIsDefined                 *bool         `json:"ItemCompleteDeliveryIsDefined"`
	ItemWeightUnit                                *string       `json:"ItemWeightUnit"`
	ItemNetWeight                                 *float32      `json:"ItemNetWeight"`
	ItemGrossWeight                               *float32      `json:"ItemGrossWeight"`
	InternalCapacityQuantity                      *float32      `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string       `json:"InternalCapacityQuantityUnit"`
	ItemIsBillingRelevant                         *bool         `json:"ItemIsBillingRelevant"`
	NetAmount                                     *float32      `json:"NetAmount"`
	TaxAmount                                     *float32      `json:"TaxAmount"`
	GrossAmount                                   *float32      `json:"GrossAmount"`
	OrderID                                       *int          `json:"OrderID"`
	OrderItem                                     *int          `json:"OrderItem"`
	ProductionOrder                               *int          `json:"ProductionOrder"`
	ProductionOrderItem                           *int          `json:"ProductionOrderItem"`
	BillOfMaterial                                *int          `json:"BillOfMaterial"`
	BillOfMaterialItem                            *int          `json:"BillOfMaterialItem"`
	Operations                                    *int          `json:"Operations"`
	OperationsItem                                *int          `json:"OperationsItem"`
	OrderType                                     *string       `json:"OrderType"`
	ContractType                                  *string       `json:"ContractType"`
	OrderValidityStartDate                        *string       `json:"OrderValidityStartDate"`
	OrderValidityEndDate                          *string       `json:"OrderValidityEndDate"`
	PaymentTerms                                  *string       `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string       `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string       `json:"PaymentDueDate"`
	NetPaymentDays                                *int          `json:"NetPaymentDays"`
	PaymentMethod                                 *string       `json:"PaymentMethod"`
	InvoicePeriodStartDate                        *string       `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                          *string       `json:"InvoicePeriodEndDate"`
	ConfirmedDeliveryDate                         *string       `json:"ConfirmedDeliveryDate"`
	Project                                       *int          `json:"Project"`
	WBSElement                                    *int          `json:"WBSElement"`
	ReferenceDocument                             *int          `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int          `json:"ReferenceDocumentItem"`
	TransactionTaxClassification                  *string       `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string       `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string       `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string       `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string       `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string       `json:"ProductAccountAssignmentGroup"`
	TaxCode                                       *string       `json:"TaxCode"`
	TaxRate                                       *float32      `json:"TaxRate"`
	CountryOfOrigin                               *string       `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string       `json:"CountryOfOriginLanguage"`
	Equipment                                     *int          `json:"Equipment"`
	ItemDeliveryBlockStatus                       *bool         `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                        *bool         `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                      *bool         `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                        *bool         `json:"ItemBillingBlockStatus"`
	CreationDate                                  *string       `json:"CreationDate"`
	CreationTime                                  *string       `json:"CreationTime"`
	LastChangeDate                                *string       `json:"LastChangeDate"`
	LastChangeTime                                *string       `json:"LastChangeTime"`
	IsCancelled                                   *bool         `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool         `json:"IsMarkedForDeletion"`
}
