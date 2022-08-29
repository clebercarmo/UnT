package shared

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"utest/adapter/postgres"
	"utest/core/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CriaFeira(dado *domain.Feira, conn *pgxpool.Pool) (*domain.Feira, error) {
	ctx := context.Background()
	var feira = domain.Feira{}

	

	err := conn.QueryRow(
		ctx,
		"INSERT INTO feira (long,lat,setcens,areap,coddist,distrito,codsubpref,subprere,regiao5,regiao8,nomefreira,registo,logradouro,numero,bairro,referencia) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) returning *",
		dado.Long,
		dado.Lat,
		dado.SetCens,
		dado.AreaP,
		dado.CodDist,
		dado.Distrito,
		dado.CodSubPref,
		dado.SubPrere,
		dado.Regiao5,
		dado.Regiao8,
		dado.NomeFreira,
		dado.Registo,
		dado.Logradouro,
		dado.Numero,
		dado.Bairro,
		dado.Referencia,
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

func ImportarDados() {
	conn := postgres.GetConnection(context.Background())
	file, err := os.Open("DEINFO_AB_FEIRASLIVRES_2014.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.Comma = ','
	for {

		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		dado := domain.Feira{Long: rec[1],
			Lat:        rec[2],
			SetCens:    rec[3],
			AreaP:      rec[4],
			CodDist:    rec[5],
			Distrito:   rec[6],
			CodSubPref: rec[7],
			SubPrere:   rec[8],
			Regiao5:    rec[9],
			Regiao8:    rec[10],
			NomeFreira: rec[11],
			Registo:    rec[12],
			Logradouro: rec[13],
			Numero:     rec[14],
			Bairro:     rec[15],
			Referencia: rec[16]}
		_, err = CriaFeira(&dado, conn)

	}

}
