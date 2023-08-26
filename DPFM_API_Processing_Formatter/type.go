package dpfm_api_processing_formatter

type HeaderUpdates struct {
	DeliveryDocument                       int     `json:"DeliveryDocument"`
	DeliveryDocumentDate     			   *string `json:"DeliveryDocumentDate"`
	PlannedGoodsIssueDate                  *string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string `json:"PlannedGoodsReceiptTime"`
	InvoiceDocumentDate                    *string `json:"InvoiceDocumentDate"`
	HeaderBillingBlockStatus			   *string `json:"HeaderBillingBlockStatus"`
	GoodsIssueOrReceiptSlipNumber		   *string `json:"GoodsIssueOrReceiptSlipNumber"`
	Incoterms							   *string `json:"Incoterms"`
	HeaderDeliveryBlockStatus			   *string `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus			   *string `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus			   *string `json:"HeaderReceivingBlockStatus"`
	}

type ItemUpdates struct {
	DeliveryDocument                        int     `json:"DeliveryDocument"`
	DeliveryDocumentItem                    int     `json:"DeliveryDocumentItem"`
	DeliverToPlantStorageLocation           *string `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation         *string `json:"DeliverFromPlantStorageLocation"`
	PlannedGoodsIssueDate            		*string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime            		*string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate          		*string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime          		*string  `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity        		*float32 `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit   		*float32 `json:"PlannedGoodsIssueQtyInBaseUnit"`
	PlannedGoodsReceiptQuantity      		*float32 `json:"PlannedGoodsReceiptQuantity"`
	PlannedGoodsReceiptQtyInBaseUnit 		*float32 `json:"PlannedGoodsReceiptQtyInBaseUnit"`
	ActualGoodsIssueDate            		*string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime            		*string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate          		*string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime          		*string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity        		*float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit   		*float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity      		*float32 `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit 		*float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ItemGrossWeight              			*float32 `json:"ItemGrossWeight"`
	ItemNetWeight                			*float32 `json:"ItemNetWeight"`
	ItemWeightUnit               			*string  `json:"ItemWeightUnit"`
	DueCalculationBaseDate       			*string  `json:"DueCalculationBaseDate"`
	PaymentDueDate               			*string  `json:"PaymentDueDate"`
	NetPaymentDays               			*int     `json:"NetPaymentDays"`
	ItemDeliveryBlockStatus      			*bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus       			*bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus     			*bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus       			*bool    `json:"ItemBillingBlockStatus"`
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
