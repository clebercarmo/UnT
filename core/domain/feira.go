package domain

import (
	"gorm.io/gorm"
	"net/http"
	"utest/core/dto"
)

// Feira entidade da tabela
type Feira struct {
	gorm.Model
	Id         int16
	Long       string
	Lat        string
	SetCens    string
	AreaP      string
	CodDist    string
	Distrito   string
	CodSubPref string
	SubPrere   string
	Regiao5    string
	Regiao8    string
	NomeFreira string
	Registo    string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
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
	Create(criatefeiraRequest *dto.CreateFeiraRequest, request *http.Request)
	Update(updatefeiraRequest *dto.UpdateFeiraRequest, request *http.Request)
	Delete(deletefeiraRequest *dto.DeleteFeiraRequest, request *http.Request)
	GetNome(getnomeRequest *dto.GetNomeRequest, request *http.Request)
}

// FeiraRepository é o contrato da camada de banco de dados
type FeiraRepository interface {
	Create(criatefeiraRequest *dto.CreateFeiraRequest) (*Feira, error)
	Update(updatefeiraRequest *dto.UpdateFeiraRequest) (*Feira, int16, error)
	Delete(deletefeiraRequest *dto.DeleteFeiraRequest) (int16, error)
	GetNome(getnomeRequest	*dto.GetNomeRequest) (string, error)
}