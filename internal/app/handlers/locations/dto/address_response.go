package dto

type AddressResponse struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	Uf          string `json:"uf,omitempty"`
	Estado      string `json:"estado,omitempty"`
	Regiao      string `json:"regiao,omitempty"`
	Ddd         string `json:"ddd,omitempty"`
}
