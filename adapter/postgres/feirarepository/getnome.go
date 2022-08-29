package feirarepository

import (
	"context"
	"utest/core/domain"
)

func (repository repository) GetNome(nome string)  (*domain.Feira, error){

	var feira domain.Feira

	rows, err := repository.db.Query(context.Background(), "SELECT * FROM feira WHERE nomefreira=$1", nome)
	if err != nil {
		return nil, err
	}

	for rows.Next() {		

		rows.Scan(
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

	}	

	return &feira, nil

}
