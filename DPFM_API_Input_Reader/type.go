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
	DeliveryDocument                       int       `json:"DeliveryDocument"`
	DeliveryDocumentDate                   string    `json:"DeliveryDocumentDate"`
	SupplyChainRelationshipID              *int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      *int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID *int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int      `json:"SupplyChainRelationshipPaymentID"`
	SupplyChainRelationshipFreightID       *int      `json:"SupplyChainRelationshipFreightID"`
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
	FreightPartner                         *int      `json:"FreightPartner"`
	IsExportImport                         *bool     `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string   `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string   `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int      `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int      `json:"ReferenceDocumentItem"`
	OrderID                                *int      `json:"OrderID"`
	OrderItem                              *int      `json:"OrderItem"`
	Contract                               *int      `json:"Contract"`
	ContractItem                           *int      `json:"ContractItem"`
	ProductionVersion                      *int      `json:"ProductionVersion"`
	ProductionVersionItem                  *int      `json:"ProductionVersionItem"`
	ProductionOrder                        *int      `json:"ProductionOrder"`
	ProductionOrderItem                    *int      `json:"ProductionOrderItem"`
	Operations                             *int      `json:"Operations"`
	OperationsItem                         *int      `json:"OperationsItem"`
	OperationID                            *int      `json:"OperationID"`
	BillOfMaterial                         *int      `json:"BillOfMaterial"`
	BillOfMaterialItem                     *int      `json:"BillOfMaterialItem"`
	ContractType                           *string   `json:"ContractType"`
	OrderValidityStartDate                 *string   `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string   `json:"OrderValidityEndDate"`
	PlannedGoodsIssueDate                  *string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string   `json:"PlannedGoodsReceiptTime"`
	FreightOrder                           *int      `json:"FreightOrder"`
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
	ExternalReferenceDocument              *string   `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain              *string   `json:"CertificateAuthorityChain"`
	UsageControlChain                      *string   `json:"UsageControlChain"`
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

type Item struct {
	DeliveryDocument               int           `json:"DeliveryDocument"`
	DeliveryDocumentItem           int           `json:"DeliveryDocumentItem"`
	PlannedGoodsIssueDate          *string       `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime          *string       `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate        *string       `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime        *string       `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity      *float32      `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit *float32      `json:"PlannedGoodsIssueQtyInBaseUnit"`
	ItemPicking                    []ItemPicking `json:"ItemPicking"`
}

type ItemPicking struct {
	DeliveryDocument                                 int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                             int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemPickingID                    int      `json:"DeliveryDocumentItemPickingID"`
	SupplyChainRelationshipID                        *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID                *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID           *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                            *int     `json:"Buyer"`
	Seller                                           *int     `json:"Seller"`
	Product                                          *string  `json:"Product"`
	DeliverToParty                                   *int     `json:"DeliverToParty"`
	DeliverToPlant                                   *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                    *string  `json:"DeliverToPlantStorageLocation"`
	DeliverToPlantStorageBin                         *string  `json:"DeliverToPlantStorageBin"`
	DeliverToPlantKanbanContainer                    *int     `json:"DeliverToPlantKanbanContainer"`
	DeliverFromParty                                 *int     `json:"DeliverFromParty"`
	DeliverFromPlant                                 *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation                  *string  `json:"DeliverFromPlantStorageLocation"`
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
	CreationDate                                     *string  `json:"CreationDate"`
	CreationTime                                     *string  `json:"CreationTime"`
	LastChangeDate                                   *string  `json:"LastChangeDate"`
	LastChangeTime                                   *string  `json:"LastChangeTime"`
	IsCancelled                                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                              *bool    `json:"IsMarkedForDeletion"`
}
