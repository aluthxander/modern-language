package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(SimpleRepository *SimpleRepository) (*SimpleService, error) {
	if SimpleRepository.Error {
		return nil, errors.New("error creating repository")
	}else{
		return &SimpleService{
			SimpleRepository: SimpleRepository,
		}, nil
	}
}
