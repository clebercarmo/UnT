package feirarepository

import (
	"context"
	"errors"
	"utest/core/dto"
)

func (repository repository) Delete(feiraRequest *dto.DeleteFeiraRequest) error {


	commandTag, err := repository.db.Exec(context.Background(), "delete from feira where id=$1", feiraRequest.Id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}
	return nil


}
