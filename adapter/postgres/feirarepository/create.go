package feirarepository

import (
	"context"
	"utest/core/domain"
	"utest/core/dto"
)


func (repository repository) Create(
	feiraRequest *dto.CreateFeiraRequest,
) (*domain.Feira, error) {
	ctx := context.Background()
	feira := domain.Feira{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO feira (long,lat,setcens,areap,coddist,distrito,codsubpref,subprere,regiao5,regiao8,nomefreira,registo,logradouro,numero,bairro,referencia) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) returning *",
		feiraRequest.Long,
		feiraRequest.Lat,
		feiraRequest.SetCens,
		feiraRequest.AreaP,
		feiraRequest.CodDist,
		feiraRequest.Distrito,
		feiraRequest.CodSubPref,
		feiraRequest.SubPrere,
		feiraRequest.Regiao5,
		feiraRequest.Regiao8,
		feiraRequest.NomeFreira,
		feiraRequest.Registo,
		feiraRequest.Logradouro,
		feiraRequest.Numero,
		feiraRequest.Bairro,
		feiraRequest.Referencia,
	).Scan(
		&feira.ID,
		&feira.Long, 
		&feira.Lat, 
		&feira.SetCens, 
		&feira.AreaP, 
		&feira.CodDist, 
		&feira.Distrito, 
		&feira.CodSubPref, 
		&feira.SubPrere,
		&feira.Regiao5, 
		&feira.Regiao8, 
		&feira.NomeFreira, 
		&feira.Registo, 
		&feira.Logradouro, 
		&feira.Numero, 
		&feira.Bairro, 
		&feira.Referencia, 
	)

	if err != nil {
		return nil, err
	}

	return &feira, nil
}
