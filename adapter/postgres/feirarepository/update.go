package feirarepository

import (
	"context"
	"utest/core/domain"
	"utest/core/dto"
)

func (repository repository) Update(feiraRequest *dto.UpdateFeiraRequest) (*domain.Feira, error) {

	ctx := context.Background()
	feira := domain.Feira{}

	err := repository.db.QueryRow(ctx,
		"update freira set long = $2 ,lat= $3 ,setcens= $4 ,areap= $5 ,coddist= $6 ,distrito= $7 ,codsubpref= $8 ,subprere= $9 ,regiao5= $10 ,regiao8= $11 ,nomefreira= $12 ,registo= $13 ,logradouro= $14 ,numero= $15 ,bairro=  $16,referencia  = $17 where id=$1",
		feiraRequest.Id,
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
		feiraRequest.Referencia).Scan(
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
