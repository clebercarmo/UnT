package dto

import (
	"encoding/json"
	"io"
)

// CreateFeiraRequest uma representação do request para criação de um novo registro
type CreateFeiraRequest struct {
	Id         int16  `json:"id"`
	Long       string `json:"long"`
	Lat        string `json:"lat"`
	SetCens    string `json:"setcens"`
	AreaP      string `json:"areap"`
	CodDist    string `json:"coddist"`
	Distrito   string `json:"distrito"`
	CodSubPref string `json:"codsubpref"`
	SubPrere   string `json:"subprere"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	NomeFreira string `json:"nomefreira"`
	Registo    string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}

type UpdateFeiraRequest struct {
	Long       string `json:"long"`
	Lat        string `json:"lat"`
	SetCens    string `json:"setcens"`
	AreaP      string `json:"areap"`
	CodDist    string `json:"coddist"`
	Distrito   string `json:"distrito"`
	CodSubPref string `json:"codsubpref"`
	SubPrere   string `json:"subprere"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	NomeFreira string `json:"nomefreira"`
	Registo    string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}

type DeleteFeiraRequest struct {
	Id int16 `json:"id"`
}


type GetNomeRequest struct {
	NomeFreira string `json:"nomefreira"`
}

// FromJSONCreateFeiraRequest converte json para um CreateFeiraRequest struct
func FromJSONCreateFeiraRequest(body io.Reader) (*CreateFeiraRequest, error) {
	CreateFeiraRequest := CreateFeiraRequest{}
	if err := json.NewDecoder(body).Decode(&CreateFeiraRequest); err != nil {
		return nil, err
	}

	return &CreateFeiraRequest, nil
}

// FromJSONUpdateFeiraRequest converte json para um UpdateFeiraRequest struct
func FromJSONUpdateFeiraRequest(body io.Reader) (*UpdateFeiraRequest, error) {
	UpdateFeiraRequest := UpdateFeiraRequest{}
	if err := json.NewDecoder(body).Decode(&UpdateFeiraRequest); err != nil {
		return nil, err
	}

	return &UpdateFeiraRequest, nil
}


// FromJSONUpdateFeiraRequest converte json para um DeleteFeiraRequest struct
func FromJSONDeleteFeiraRequest(body io.Reader) (*DeleteFeiraRequest, error) {
	DeleteFeiraRequest := DeleteFeiraRequest{}
	if err := json.NewDecoder(body).Decode(&DeleteFeiraRequest); err != nil {
		return nil, err
	}

	return &DeleteFeiraRequest, nil
}


func FromJSONGetNameRequest(body io.Reader) (*GetNomeRequest, error) {
	GetNomeRequest := GetNomeRequest{}
	if err := json.NewDecoder(body).Decode(&GetNomeRequest); err != nil {
		return nil, err
	}

	return &GetNomeRequest, nil
}