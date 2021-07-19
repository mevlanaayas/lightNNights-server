package score

import (
	"fmt"
	"lightNNights/errors"
)

type CommandService struct {
	repository CommandRepository
}

func NewCommandService(repository CommandRepository) CommandService {
	return CommandService{repository: repository}
}

func (receiver CommandService) Save(score Score) errors.Error {
	err := receiver.repository.Save(score)
	if err != nil {
		return errors.New("error while saving score", 100, fmt.Errorf("error while saving score \n\t%v", err))
	}
	return nil
}

func (receiver CommandService) Update(score Score) errors.Error {
	affectedRowCount, err := receiver.repository.Update(score)
	if err != nil {
		return errors.New("error while updating score", 100, fmt.Errorf("error while updating score \n\t%v", err))
	}
	if affectedRowCount != 1 {
		return errors.New("error while updating score unexpected row count", 100, fmt.Errorf("error while updating score unexpected row count %d", affectedRowCount))
	}
	return nil
}
