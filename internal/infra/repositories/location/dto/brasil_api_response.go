package dto

type BrasilApiStateResponse struct {
	// temos anotação de json, toJson from Json
	Sigla string `json:"sigla,omitempty"`
	Nome  string `json:"nome,omitempty"`
}
