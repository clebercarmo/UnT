package feirarepository

import (
	"context"
	"errors"
)

func (repository repository) Delete(id int) error {


	commandTag, err := repository.db.Exec(context.Background(), "delete from feira where id=$1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("registro não encontrado")
	}
	return nil


}
