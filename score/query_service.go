package score

import (
	"fmt"
	"lightNNights/errors"
)

type QueryService struct {
	repository QueryRepository
}

func NewQueryService(repository QueryRepository) QueryService {
	return QueryService{repository: repository}
}

func (receiver QueryService) Get(id string) errors.Error {
	err := receiver.repository.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("error while getting score with id %s", id), 100, fmt.Errorf("error while getting score with id %s \n\t%v", id, err))
	}
	return nil
}

func (receiver QueryService) List() ([]Score, errors.Error) {
	scores, err := receiver.repository.List()
	if err != nil {
		return nil, errors.New("error while listing scores", 100, fmt.Errorf("error while listing scores \n\t%v", err))
	}
	return scores, nil
}
