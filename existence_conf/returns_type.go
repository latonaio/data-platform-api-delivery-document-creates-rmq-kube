package existence_conf

type Returns struct {
	ConnectionKey                                        string                                                `json:"connection_key"`
	Result                                               bool                                                  `json:"result"`
	RedisKey                                             string                                                `json:"redis_key"`
	RuntimeSessionID                                     string                                                `json:"runtime_session_id"`
	BusinessPartnerID                                    *int                                                  `json:"business_partner"`
	Filepath                                             string                                                `json:"filepath"`
	ServiceLabel                                         string                                                `json:"service_label"`
	BPGeneralReturn                                      BPGeneralReturn                                       `json:"BusinessPartnerGeneral"`
	PlantGeneralReturn                                   PlantGeneralReturn                                    `json:"PlantGeneral"`
	SupplyChainRelationshipGeneralReturn                 *SupplyChainRelationshipGeneralReturn                 `json:"SupplyChainRelationshipGeneral,omitempty"`
	SupplyChainRelationshipBillingRelationReturn         *SupplyChainRelationshipBillingRelationReturn         `json:"SupplyChainRelationshipBillingRelation,omitempty"`
	SupplyChainRelationshipPaymentRelationReturn         *SupplyChainRelationshipPaymentRelationReturn         `json:"SupplyChainRelationshipPaymentRelation,omitempty"`
	SupplyChainRelationshipDeliveryRelationReturn        *SupplyChainRelationshipDeliveryRelationReturn        `json:"SupplyChainRelationshipDeliveryRelation,omitempty"`
	SupplyChainRelationshipDeliveryPlantRelationReturn   *SupplyChainRelationshipDeliveryPlantRelationReturn   `json:"SupplyChainRelationshipDeliveryPlantRelation,omitempty"`
	SupplyChainRelationshipProductionPlantRelationReturn *SupplyChainRelationshipProductionPlantRelationReturn `json:"SupplyChainRelationshipProductionPlantRelation,omitempty"`
	SupplyChainRelationshipStockConfPlantRelationReturn  *SupplyChainRelationshipStockConfPlantRelationReturn  `json:"SupplyChainRelationshipStockConfPlantRelation,omitempty"`
	IncotermsReturn                                      IncotermsReturn                                       `json:"Incoterms"`
	CountryReturn                                        CountryReturn                                         `json:"Country"`
	AddressReturn                                        AddressReturn                                         `json:"Address"`
	StorageLocationReturn                                StorageLocationReturn                                 `json:"StorageLocation"`
	BatchReturn                                          BatchReturn                                           `json:"Batch"`
	QuantityUnitReturn                                   QuantityUnitReturn                                    `json:"QuantityUnit"`
	APISchema                                            string                                                `json:"api_schema"`
	Accepter                                             []string                                              `json:"accepter"`
	Deleted                                              bool                                                  `json:"deleted"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type PlantGeneralReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type SupplyChainRelationshipGeneralReturn struct {
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`
	Buyer                     int `json:"Buyer"`
	Seller                    int `json:"Seller"`
}

type SupplyChainRelationshipBillingRelationReturn struct {
	SupplyChainRelationshipID        int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int `json:"Buyer"`
	Seller                           int `json:"Seller"`
	BillToParty                      int `json:"BillToParty"`
	BillFromParty                    int `json:"BillFromParty"`
}

type SupplyChainRelationshipPaymentRelationReturn struct {
	SupplyChainRelationshipID        int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int `json:"Buyer"`
	Seller                           int `json:"Seller"`
	BillToParty                      int `json:"BillToParty"`
	BillFromParty                    int `json:"BillFromParty"`
	Payer                            int `json:"Payer"`
	Payee                            int `json:"Payee"`
}

type SupplyChainRelationshipDeliveryRelationReturn struct {
	SupplyChainRelationshipID         int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int `json:"Buyer"`
	Seller                            int `json:"Seller"`
	DeliverToParty                    int `json:"DeliverToParty"`
	DeliverFromParty                  int `json:"DeliverFromParty"`
}

type SupplyChainRelationshipDeliveryPlantRelationReturn struct {
	SupplyChainRelationshipID              int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int    `json:"Buyer"`
	Seller                                 int    `json:"Seller"`
	DeliverToParty                         int    `json:"DeliverToParty"`
	DeliverFromParty                       int    `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
}

type SupplyChainRelationshipProductionPlantRelationReturn struct {
	SupplyChainRelationshipID                int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int    `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int    `json:"Buyer"`
	Seller                                   int    `json:"Seller"`
	ProductionPlantBusinessPartner           int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string `json:"ProductionPlant"`
}

type SupplyChainRelationshipStockConfPlantRelationReturn struct {
	SupplyChainRelationshipID               int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int    `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int    `json:"Buyer"`
	Seller                                  int    `json:"Seller"`
	StockConfirmationBusinessPartner        int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string `json:"StockConfirmationPlant"`
}

type CountryReturn struct {
	Country string `json:"Country"`
}

type IncotermsReturn struct {
	Incoterms string `json:"Incoterms"`
}

type AddressReturn struct {
	AddressID       int    `json:"AddressID"`
	ValidityEndDate string `json:"ValidityEndDate"`
}

type BatchReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Product         string `json:"Product"`
	Plant           string `json:"Plant"`
	Batch           string `json:"Batch"`
}

type StorageLocationReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
}

type QuantityUnitReturn struct {
	QuantityUnit string `json:"QuantityUnit"`
}
