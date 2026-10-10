package dto

type BrasilApiStateResponse struct {
	Id      int                     `json:"id,omitempty"`
	Sigla   string                  `json:"sigla,omitempty"`
	Nome    string                  `json:"nome,omitempty"`
	Regiao  BrasilApiRegionResponse `json:"regiao"`
	Capital string                  `json:"capital,omitempty"`
}

type BrasilApiRegionResponse struct {
	Id    int    `json:"id,omitempty"`
	Sigla string `json:"sigla,omitempty"`
	Nome  string `json:"nome,omitempty"`
}
