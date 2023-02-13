package existence_conf

type ExConfMapper struct {
	ServiceLabel          string  `json:"ServiceLabel"`
	APIType               string  `json:"APIType"`
	APIName               string  `json:"APIName"`
	Field                 string  `json:"Field"`
	ExConfAPIServiceName  *string `json:"ExConfAPIServiceName"`
	ExConfAPIName         *string `json:"ExConfAPIName"`
	ExConfAPIQueueName    *string `json:"ExConfAPIQueueName"`
	ExConfProgramFileName *string `json:"ExConfProgramFileName"`
	Tabletag              *string `json:"Tabletag"`
	TableConfirmed        *string `json:"TableConfirmed"`
	ExConfAPITyp          *string `json:"ExConfAPITyp"`
}
