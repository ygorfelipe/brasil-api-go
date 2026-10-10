package dto

type StatesResponse struct {
	Acronym string         `json:"sigla,omitempty"`
	Name    string         `json:"nome,omitempty"`
	Capital string         `json:"capital,omitempty"`
	Region  RegionResponse `json:"regiao"`
}
