package feirarepository

import (
	"utest/core/domain"
	"utest/core/dto"
)

func (repository repository) Delete(feiraRequest *dto.DeleteFeiraRequest) error {
  retorno :=	repository.db.Delete(&domain.Feira{}, feiraRequest.Id) // find product with integer primary key
  if retorno.Error != nil {
	return retorno.Error
  } else {
	return nil
  }

}
