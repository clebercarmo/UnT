package feirausecase


func (usecase usecase) Delete(id int) (error) {
	err := usecase.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}