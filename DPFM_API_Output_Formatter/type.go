package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header      *[]Header      `json:"Header"`
	Item        *[]Item        `json:"Item"`
	Partner     *[]Partner     `json:"Partner"`
	Address     *[]Address     `json:"Address"`
	ItemPicking *[]ItemPicking `json:"ItemPicking"`
}

type Header struct {
	DeliveryDocument                       int      `json:"DeliveryDocument"`
	DeliveryDocumentDate                   string   `json:"DeliveryDocumentDate"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int     `json:"SupplyChainRelationshipPaymentID"`
	SupplyChainRelationshipFreightID       *int     `json:"SupplyChainRelationshipFreightID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	BillToParty                            *int     `json:"BillToParty"`
	BillFromParty                          *int     `json:"BillFromParty"`
	BillToCountry                          *string  `json:"BillToCountry"`
	BillFromCountry                        *string  `json:"BillFromCountry"`
	Payer                                  *int     `json:"Payer"`
	Payee                                  *int     `json:"Payee"`
	FreightPartner                         *int     `json:"FreightPartner"`
	IsExportImport                         *bool    `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string  `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string  `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int     `json:"ReferenceDocumentItem"`
	OrderID                                *int     `json:"OrderID"`
	OrderItem                              *int     `json:"OrderItem"`
	Contract                               *int     `json:"Contract"`
	ContractItem                           *int     `json:"ContractItem"`
	ProductionVersion                      *int     `json:"ProductionVersion"`
	ProductionVersionItem                  *int     `json:"ProductionVersionItem"`
	ProductionOrder                        *int     `json:"ProductionOrder"`
	ProductionOrderItem                    *int     `json:"ProductionOrderItem"`
	Operations                             *int     `json:"Operations"`
	OperationsItem                         *int     `json:"OperationsItem"`
	OperationID                            *int     `json:"OperationID"`
	BillOfMaterial                         *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                     *int     `json:"BillOfMaterialItem"`
	ContractType                           *string  `json:"ContractType"`
	OrderValidityStartDate                 *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string  `json:"OrderValidityEndDate"`
	PlannedGoodsIssueDate                  string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                string   `json:"PlannedGoodsReceiptTime"`
	FreightOrder                           *int     `json:"FreightOrder"`
	InvoiceDocumentDate                    *string  `json:"InvoiceDocumentDate"`
	HeaderCompleteDeliveryIsDefined        *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus                   *string  `json:"HeaderDeliveryStatus"`
	GoodsIssueOrReceiptSlipNumber          *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus                    *string  `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus                *string  `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockStatus               *bool    `json:"HeaderBillingBlockStatus"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit"`
	Incoterms                              *string  `json:"Incoterms"`
	TransactionCurrency                    *string  `json:"TransactionCurrency"`
	HeaderDeliveryBlockStatus              *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *bool    `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *bool    `json:"HeaderReceivingBlockStatus"`
	ExternalReferenceDocument              *string  `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain              *string  `json:"CertificateAuthorityChain"`
	UsageControlChain                      *string  `json:"UsageControlChain"`
	CreationDate                           string   `json:"CreationDate"`
	CreationTime                           string   `json:"CreationTime"`
	LastChangeDate                         string   `json:"LastChangeDate"`
	LastChangeTime                         string   `json:"LastChangeTime"`
	IsCancelled                            *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                    *bool    `json:"IsMarkedForDeletion"`
}

type Item struct {
	DeliveryDocument               int      `json:"DeliveryDocument"`
	DeliveryDocumentItem           int      `json:"DeliveryDocumentItem"`
	PlannedGoodsIssueDate          *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime          *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate        *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime        *string  `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity      *float32 `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit *float32 `json:"PlannedGoodsIssueQtyInBaseUnit"`
}

type ItemPicking struct {
	DeliveryDocument                                 int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                             int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemPickingID                    int      `json:"DeliveryDocumentItemPickingID"`
	SupplyChainRelationshipID                        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID                int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID           int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                            int      `json:"Buyer"`
	Seller                                           int      `json:"Seller"`
	Product                                          string   `json:"Product"`
	DeliverToParty                                   int      `json:"DeliverToParty"`
	DeliverToPlant                                   string   `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                    string   `json:"DeliverToPlantStorageLocation"`
	DeliverToPlantStorageBin                         *string  `json:"DeliverToPlantStorageBin"`
	DeliverToPlantKanbanContainer                    *int     `json:"DeliverToPlantKanbanContainer"`
	DeliverFromParty                                 int      `json:"DeliverFromParty"`
	DeliverFromPlant                                 string   `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation                  string   `json:"DeliverFromPlantStorageLocation"`
	DeliverFromPlantStorageBin                       *string  `json:"DeliverFromPlantStorageBin"`
	DeliverFromPlantKanbanContainer                  *int     `json:"DeliverFromPlantKanbanContainer"`
	DeliverToPlantPlannedPickingQuantityInBaseUnit   float32  `json:"DeliverToPlantPlannedPickingQuantityInBaseUnit"`
	DeliverFromPlantPlannedPickingQuantityInBaseUnit float32  `json:"DeliverFromPlantPlannedPickingQuantityInBaseUnit"`
	DeliverToPlantPlannedPickingDate                 string   `json:"DeliverToPlantPlannedPickingDate"`
	DeliverToPlantPlannedPickingTime                 string   `json:"DeliverToPlantPlannedPickingTime"`
	DeliverFromPlantPlannedPickingDate               string   `json:"DeliverFromPlantPlannedPickingDate"`
	DeliverFromPlantPlannedPickingTime               string   `json:"DeliverFromPlantPlannedPickingTime"`
	DeliverToPlantActualPickingQuantityInBaseUnit    *float32 `json:"DeliverToPlantActualPickingQuantityInBaseUnit"`
	DeliverToPlantActualPickingDate                  *string  `json:"DeliverToPlantActualPickingDate"`
	DeliverToPlantActualPickingTime                  *string  `json:"DeliverToPlantActualPickingTime"`
	DeliverFromPlantActualPickingQuantityInBaseUnit  *float32 `json:"DeliverFromPlantActualPickingQuantityInBaseUnit"`
	DeliverFromPlantActualPickingDate                *string  `json:"DeliverFromPlantActualPickingDate"`
	DeliverFromPlantActualPickingTime                *string  `json:"DeliverFromPlantActualPickingTime"`
	ExternalReferenceDocument                        *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                    *string  `json:"ExternalReferenceDocumentItem"`
	ExternalReferenceDocumentItemPickingID           *string  `json:"ExternalReferenceDocumentItemPickingID"`
	CreationDate                                     string   `json:"CreationDate"`
	CreationTime                                     string   `json:"CreationTime"`
	LastChangeDate                                   string   `json:"LastChangeDate"`
	LastChangeTime                                   string   `json:"LastChangeTime"`
	IsCancelled                                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                              *bool    `json:"IsMarkedForDeletion"`
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
	EmailAddress            *string `json:"EmailAddress"`
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
