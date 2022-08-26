package domain

import (
	"net/http"
	"utest/core/dto"
)

// Feira entidade da tabela
type Feira struct {
	ID		   int32   	`json:"id"`
	Long       string  	`json:"long"`
	Lat        string  	`json:"lat"`
	SetCens    string	`json:"setcens"`
	AreaP      string	`json:"arep"`
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

// FeiraService contrato da camada http
type FeiraService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	GetNome(response http.ResponseWriter, request *http.Request)
}

// FeiraUserCase é o contrato com a camada de negocio
type FeiraUseCase interface {
	Create(criatefeiraRequest *dto.CreateFeiraRequest) (*Feira, error)
	Update(updatefeiraRequest *dto.UpdateFeiraRequest) (*Feira, error)
	Delete(deletefeiraRequest *dto.DeleteFeiraRequest) error
	GetNome(getnomeRequest *dto.GetNomeRequest) (*Feira, error)
}

// FeiraRepository é o contrato da camada de banco de dados
type FeiraRepository interface {
	Create(criatefeiraRequest *dto.CreateFeiraRequest) (*Feira, error)
	Update(updatefeiraRequest *dto.UpdateFeiraRequest) (*Feira, error)
	Delete(deletefeiraRequest *dto.DeleteFeiraRequest) error
	GetNome(getnomeRequest	*dto.GetNomeRequest) (*Feira, error)
}