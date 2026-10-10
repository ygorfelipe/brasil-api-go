package dto

type RegionResponse struct {
	Id      int    `json:"id,omitempty"`
	Acronym string `json:"sigla,omitempty"`
	Name    string `json:"nome,omitempty"`
}
