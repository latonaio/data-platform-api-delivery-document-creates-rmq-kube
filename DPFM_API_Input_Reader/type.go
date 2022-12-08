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
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"DeliveryDocument"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	DeliveryDocument              *int     `json:"DeliveryDocument"`
	Buyer                         *int     `json:"Buyer"`
	Seller                        *int     `json:"Seller"`
	ReferenceDocument             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem         *int     `json:"ReferenceDocumentItem"`
	OrderID                       *int     `json:"OrderID"`
	OrderItem                     *int     `json:"OrderItem"`
	ContractType                  *string  `json:"ContractType"`
	OrderValidityStartDate        *string  `json:"OrderVaridityStartDate"`
	OrderValidityEndDate          *string  `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate      *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate        *string  `json:"InvoiceScheduleEndDate"`
	IssuingPlantTimeZone          *string  `json:"IssuingPlantTimeZone"`
	ReceivingPlantTimeZone        *string  `json:"ReceivingPlantTimeZone"`
	DocumentDate                  *string  `json:"DocumentDate"`
	PlannedGoodsIssueDate         *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime         *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate       *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime       *string  `json:"PlannedGoodsReceiptTime"`
	BillingDocumentDate           *string  `json:"BillingDocumentDate"`
	CompleteDeliveryIsDefined     *bool    `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus         *string  `json:"OverallDeliveryStatus"`
	CreationDate                  *string  `json:"CreationDate"`
	CreationTime                  *string  `json:"CreationTime"`
	IssuingBlockReason            *bool    `json:"IssuingBlockReason"`
	ReceivingBlockReason          *bool    `json:"ReceivingBlockReason"`
	GoodsIssueOrReceiptSlipNumber *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus           *string  `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus       *string  `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockReason      *bool    `json:"HeaderBillingBlockReason"`
	HeaderGrossWeight             *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight               *float32 `json:"HeaderNetWeight"`
	HeaderVolume                  *string  `json:"HeaderVolume"`
	HeaderVolumeUnit              *string  `json:"HeaderVolumeUnit"`
	HeaderWeightUnit              *string  `json:"HeaderWeightUnit"`
	Incoterms                     *string  `json:"Incoterms"`
	IsExportImportDelivery        *bool    `json:"IsExportImportDelivery"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	IssuingPlantBusinessPartner   *string  `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                  *string  `json:"IssuingPlant"`
	ReceivingPlant                *string  `json:"ReceivingPlant"`
	ReceivingPlantBusinessPartner *string  `json:"ReceivingPlantBusinessPartner"`
	DeliverToParty                *int     `json:"DeliverToParty"`
	DeliverFromParty              *int     `json:"DeliverFromParty"`
	TransactionCurrency           *string  `json:"TransactionCurrency"`
	OverallDelivReltdBillgStatus  *string  `json:"OverallDelivReltdBillgStatus"`
}

type HeaderUpdates struct {
	PlannedGoodsIssueDate         *string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime         *string `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate       *string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime       *string `json:"PlannedGoodsReceiptTime"`
	BillingDocumentDate           *string `json:"BillingDocumentDate"`
	HeaderIssuingBlockStatus      *bool   `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus    *bool   `json:"HeaderReceivingBlockStatus"`
	GoodsIssueOrReceiptSlipNumber *string `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingBlockStatus      *bool   `json:"HeaderBillingBlockStatus"`
	Incoterms                     *string `json:"Incoterms"`
	IssuingPlantBusinessPartner   *string `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                  *string `json:"IssuingPlant"`
	ReceivingPlantBusinessPartner *string `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                *string `json:"ReceivingPlant"`
	DeliverFromParty              *int    `json:"DeliverFromParty"`
	DeliverToParty                *int    `json:"DeliverToParty"`
	StockIsFullyConfirmed         *bool   `json:"StockIsFullyConfirmed"`
}

type HeaderPartner struct {
	DeliveryDocument        *int                   `json:"DeliveryDocument"`
	PartnerFunction         string                 `json:"PartnerFunction"`
	BusinessPartner         *int                   `json:"BusinessPartner"`
	BusinessPartnerFullName *string                `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string                `json:"BusinessPartnerName"`
	Organization            *string                `json:"Organization"`
	Country                 *string                `json:"Country"`
	Language                *string                `json:"Language"`
	Currency                *string                `json:"Currency"`
	ExternalDocumentID      *string                `json:"ExternalDocumentID"`
	AddressID               *int                   `json:"AddressID"`
	HeaderPartnerContact    []HeaderPartnerContact `json:"HeaderPartnerContact"`
	HeaderPartnerPlant      []HeaderPartnerPlant   `json:"HeaderPartnerPlant"`
}

type HeaderPartnerUpdates struct {
	BusinessPartner         *int    `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}

type HeaderPartnerContact struct {
	DeliveryDocument  *int   `json:"DeliveryDocument"`
	PartnerFunction   string `json:"PartnerFunction"`
	ContactID         *int   `json:"ContactID"`
	BusinessPartner   *int   `json:"BusinessPartner"`
	ContactPersonName string `json:"ContactPersonName"`
	EmailAddress      string `json:"EmailAddress"`
	PhoneNumber       string `json:"PhoneNumber"`
	MobilePhoneNumber string `json:"MobilePhoneNumber"`
	FaxNumber         string `json:"FaxNumber"`
	ContactTag1       string `json:"ContactTag1"`
	ContactTag2       string `json:"ContactTag2"`
	ContactTag3       string `json:"ContactTag3"`
	ContactTag4       string `json:"ContactTag4"`
}

type HeaderPartnerContactUpdates struct {
	ContactID         *int   `json:"ContactID"`
	ContactPersonName string `json:"ContactPersonName"`
	EmailAddress      string `json:"EmailAddress"`
	PhoneNumber       string `json:"PhoneNumber"`
	MobilePhoneNumber string `json:"MobilePhoneNumber"`
	FaxNumber         string `json:"FaxNumber"`
	ContactTag1       string `json:"ContactTag1"`
	ContactTag2       string `json:"ContactTag2"`
	ContactTag3       string `json:"ContactTag3"`
	ContactTag4       string `json:"ContactTag4"`
}

type HeaderPartnerPlant struct {
	DeliveryDocument *int   `json:"DeliveryDocument"`
	PartnerFunction  string `json:"PartnerFunction"`
	BusinessPartner  *int   `json:"BusinessPartner"`
	Plant            string `json:"Plant"`
}

type HeaderPartnerPlantUpdates struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type HeaderPDF struct {
	DeliveryDocument         *int   `json:"DeliveryDocument"`
	DocType                  string `json:"DocType"`
	DocVersionID             *int   `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner *int   `json:"DocIssuerBusinessPartner"`
	FileName                 string `json:"FileName"`
}

type HeaderPDFUpdates struct {
	DocIssuerBusinessPartner *int   `json:"DocIssuerBusinessPartner"`
	FileName                 string `json:"FileName"`
}

type Address struct {
	DeliveryDocument *int   `json:"DeliveryDocument"`
	AddressID        *int   `json:"AddressID"`
	PostalCode       string `json:"PostalCode"`
	LocalRegion      string `json:"LocalRegion"`
	Country          string `json:"Country"`
	District         string `json:"District"`
	StreetName       string `json:"StreetName"`
	CityName         string `json:"CityName"`
	Building         string `json:"Building"`
	Floor            *int   `json:"Floor"`
	Room             *int   `json:"Room"`
}

type AddressUpdates struct {
	PostalCode  string `json:"PostalCode"`
	LocalRegion string `json:"LocalRegion"`
	Country     string `json:"Country"`
	District    string `json:"District"`
	StreetName  string `json:"StreetName"`
	CityName    string `json:"CityName"`
	Building    string `json:"Building"`
	Floor       *int   `json:"Floor"`
	Room        *int   `json:"Room"`
}

type Item struct {
	DeliveryDocument                             *int          `json:"DeliveryDocument"`
	DeliveryDocumentItem                         *int          `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                 string        `json:"DeliveryDocumentItemCategory"`
	DeliveryDocumentItemText                     string        `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer              string        `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller             string        `json:"DeliveryDocumentItemTextBySeller"`
	Product                                      string        `json:"Product"`
	ProductStandardID                            string        `json:"ProductStandardID"`
	ProductGroup                                 string        `json:"ProductGroup"`
	BaseUnit                                     string        `json:"BaseUnit"`
	OriginalQuantityInBaseUnit                   *float32      `json:"ActualGoodsIssueQtyInBaseUnit"`
	IssuingUnit                                  string        `json:"IssuingUnit"`
	ReceivingUnit                                string        `json:"ReceivingUnit"`
	ActualGoodsIssueDate                         *string       `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                         *string       `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                       *string       `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                       *string       `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQtyInBaseUnit                *float32      `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsIssueQuantity                     *float32      `json:"ActualGoodsIssueQuantity"`
	ActualGoodsReceiptQtyInBaseUnit              *float32      `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                   *float32      `json:"ActualGoodsReceiptQuantity"`
	CompleteItemDeliveryIsDefined                *bool         `json:"CompleteItemDeliveryIsDefined"`
	StockConfirmationPartnerFunction             string        `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner             *int          `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string        `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  string        `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string       `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string       `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPolicy                      string        `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                      string        `json:"StockConfirmationStatus"`
	ProductionPlantPartnerFunction               string        `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner               string        `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                              string        `json:"ProductionPlant"`
	ProductionPlantStorageLocation               string        `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction                  string        `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner                  string        `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                                 string        `json:"IssuingPlant"`
	IssuingPlantStorageLocation                  string        `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction                string        `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner                string        `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                               string        `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation                string        `json:"ReceivingPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant       *bool         `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsBatchManagedInIssuingPlant          *bool         `json:"ProductIsBatchManagedInIssuingPlant"`
	ProductIsBatchManagedInReceivingPlant        *bool         `json:"ProductIsBatchManagedInReceivingPlant"`
	BatchMgmtPolicyInProductionPlant             string        `json:"BatchMgmtPolicyInProductionPlant"`
	BatchMgmtPolicyInIssuingPlant                string        `json:"BatchMgmtPolicyInIssuingPlant"`
	BatchMgmtPolicyInReceivingPlant              string        `json:"BatchMgmtPolicyInReceivingPlant"`
	ProductionPlantBatch                         string        `json:"ProductionPlantBatch"`
	IssuingPlantBatch                            string        `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                          string        `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate        *string       `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate          *string       `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate           *string       `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate             *string       `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate         *string       `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate           *string       `json:"ReceivingPlantBatchValidityEndDate"`
	CreationDate                                 *string       `json:"CreationDate"`
	CreationTime                                 *string       `json:"CreationTime"`
	ItemBillingStatus                            string        `json:"ItemBillingStatus"`
	ItemBillingConfStatus                        string        `json:"ItemBillingConfStatus"`
	SalesCostGLAccount                           string        `json:"SalesCostGLAccount"`
	ReceivingGLAccount                           string        `json:"ReceivingGLAccount"`
	IssuingGoodsMovementType                     string        `json:"IssuingGoodsMovementType"`
	ReceivingGoodsMovementType                   string        `json:"ReceivingGoodsMovementType"`
	ItemDeliveryBlockStatus                      *bool         `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                       *bool         `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                     *bool         `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                       *bool         `json:"ItemBillingBlockStatus"`
	ItemCompleteDeliveryIsDefined                *bool         `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryIncompletionStatus               string        `json:"ItemDeliveryIncompletionStatus"`
	ItemGrossWeight                              *float32      `json:"ItemGrossWeight"`
	ItemNetWeight                                *float32      `json:"ItemNetWeight"`
	ItemWeightUnit                               string        `json:"ItemWeightUnit"`
	ItemIsBillingRelevant                        *bool         `json:"ItemIsBillingRelevant"`
	LastChangeDate                               *string       `json:"LastChangeDate"`
	OrderID                                      *int          `json:"OrderID"`
	OrderItem                                    *int          `json:"OrderItem"`
	OrderType                                    string        `json:"OrderType"`
	ContractType                                 string        `json:"ContractType"`
	OrderValidityStartDate                       *string       `json:"OrderValidityStartDate"`
	OrderValidityEndDate                         *string       `json:"OrderValidityEndDate"`
	PaymentTerms                                 string        `json:"PaymentTerms"`
	PaymentMethod                                string        `json:"PaymentMethod"`
	InvoiceScheduleStartDate                     *string       `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate                       *string       `json:"InvoiceScheduleEndDate"`
	ConfirmedDeliveryDate                        *string       `json:"ConfirmedDeliveryDate"`
	Project                                      string        `json:"Project"`
	ProfitCenter                                 string        `json:"ProfitCenter"`
	ReferenceDocument                            *int          `json:"ReferenceDocument"`
	ReferenceDocumentItem                        *int          `json:"ReferenceDocumentItem"`
	BPTaxClassification                          string        `json:"BPTaxClassification"`
	ProductTaxClassification                     string        `json:"ProductTaxClassification"`
	BPAccountAssignmentGroup                     string        `json:"BPAccountAssignmentGroup"`
	ProductAccountAssignmentGroup                string        `json:"ProductAccountAssignmentGroup"`
	TaxCode                                      string        `json:"TaxCode"`
	TaxRate                                      *float32      `json:"TaxRate"`
	CountryOfOrigin                              string        `json:"CountryOfOrigin"`
	StockIsFullyConfirmed                        *bool         `json:"StockIsFullyConfirmed"`
	ItemPartner                                  []ItemPartner `json:"ItemPartner"`
}

type ItemUpdates struct {
	ActualGoodsIssueDate                         *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                         *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                       *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                       *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQtyInBaseUnit                *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsIssueQuantity                     *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsReceiptQtyInBaseUnit              *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                   *float32 `json:"ActualGoodsReceiptQuantity"`
	StockConfirmationPartnerFunction             string   `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner             *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPolicy                      string   `json:"StockConfirmationPolicy"`
	ProductionPlantStorageLocation               string   `json:"ProductionPlantStorageLocation"`
	IssuingPlantStorageLocation                  string   `json:"IssuingPlantStorageLocation"`
	ReceivingPlantStorageLocation                string   `json:"ReceivingPlantStorageLocation"`
	ProductionPlantBatch                         string   `json:"ProductionPlantBatch"`
	IssuingPlantBatch                            string   `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                          string   `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate        *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate          *string  `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate           *string  `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate             *string  `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate         *string  `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate           *string  `json:"ReceivingPlantBatchValidityEndDate"`
	ItemReceivingBlockStatus                     *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                       *bool    `json:"ItemBillingBlockStatus"`
	ItemGrossWeight                              *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                                *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                               string   `json:"ItemWeightUnit"`
	ItemIsBillingRelevant                        *bool    `json:"ItemIsBillingRelevant"`
}

type ItemPartner struct {
	DeliveryDocument     *int               `json:"DeliveryDocument"`
	DeliveryDocumentItem *int               `json:"DeliveryDocumentItem"`
	PartnerFunction      string             `json:"PartnerFunction"`
	BusinessPartner      *int               `json:"BusinessPartner"`
	ItemPartnerPlant     []ItemPartnerPlant `json:"ItemPartnerPlant"`
}

type ItemPartnerUpdates struct {
	BusinessPartner *int `json:"BusinessPartner"`
}

type ItemPartnerPlant struct {
	DeliveryDocument     *int   `json:"DeliveryDocument"`
	DeliveryDocumentItem *int   `json:"DeliveryDocumentItem"`
	PartnerFunction      string `json:"PartnerFunction"`
	BusinessPartner      *int   `json:"BusinessPartner"`
	Plant                string `json:"Plant"`
}

type ItemPartnerPlantUpdates struct {
	Plant string `json:"Plant"`
}
