package dto

import (
	"encoding/json"
	"io"
)

// CreateFeiraRequest uma representação do request para criação de um novo registro
type CreateFeiraRequest struct{
	Id         int16	`json:"id"`
	Long       string   `json:"long"`
	Lat        string   `json:"lat"`
	SetCens    string   `json:"setcens"`
	AreaP      string	`json:"areap"`
	CodDist    string	`json:"coddist"`
	Distrito   string	`json:"distrito"`
	CodSubPref string	`json:"codsubpref"`
	SubPrere   string	`json:"subprere"`
	Regiao5    string	`json:"regiao5"`
	Regiao8    string	`json:"regiao8"`
	NomeFreira string	`json:"nomefreira"`
	Registo    string	`json:"registro"`
	Logradouro string	`json:"logradouro"`
	Numero     string	`json:"numero"`
	Bairro     string	`json:"bairro"`
	Referencia string	`json:"referencia"`

}


// FromJSONCreateFeiraRequest converte json para um CreateFeiraRequest struct
func FromJSONCreateFeiraRequest(body io.Reader) (*CreateFeiraRequest, error) {
	CreateFeiraRequest := CreateFeiraRequest{}
	if err := json.NewDecoder(body).Decode(&CreateFeiraRequest); err != nil {
		return nil, err
	}

	return &CreateFeiraRequest, nil
}
