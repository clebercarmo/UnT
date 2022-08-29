package feirarepository

import (
	"context"
	"errors"
	"utest/core/domain"
	"utest/core/dto"
)

func (repository repository) Update(id int, feiraRequest *dto.UpdateFeiraRequest) (*domain.Feira, error) {

	ctx := context.Background()
	feira := domain.Feira{}

	commandTag, err := repository.db.Exec(ctx,
		"update feira set long = $2 ,lat= $3 ,setcens= $4 ,areap= $5 ,coddist= $6 ,distrito= $7 ,codsubpref= $8 ,subprere= $9 ,regiao5= $10 ,regiao8= $11 ,nomefreira= $12 ,registo= $13 ,logradouro= $14 ,numero= $15 ,bairro=  $16,referencia  = $17 where id=$1",
		id,
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
		feiraRequest.Referencia)

	if err != nil {
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		return nil, errors.New("registro não encontrado")
	}

	feira.Long = feiraRequest.Long
	feira.Lat = feiraRequest.Lat
	feira.SetCens = feiraRequest.SetCens
	feira.AreaP = feiraRequest.AreaP
	feira.CodDist = feiraRequest.CodDist
	feira.Distrito = feiraRequest.Distrito
	feira.CodSubPref = feiraRequest.CodSubPref
	feira.SubPrere = feiraRequest.SubPrere
	feira.Regiao5 = feiraRequest.Regiao5
	feira.Regiao8 = feiraRequest.Regiao8
	feira.NomeFreira = feiraRequest.NomeFreira
	feira.Registo = feiraRequest.Registo
	feira.Logradouro = feiraRequest.Logradouro
	feira.Numero = feiraRequest.Numero
	feira.Bairro = feiraRequest.Bairro
	feira.Referencia = feiraRequest.Referencia

	return &feira, nil

}
