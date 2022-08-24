package feirausecase

import (
	"utest/core/dto"
)


func (usecase usecase) Delete(feiraDeleteRequest *dto.DeleteFeiraRequest) (error) {
	err := usecase.repository.Delete(feiraDeleteRequest)

	if err != nil {
		return err
	}

	return nil
}