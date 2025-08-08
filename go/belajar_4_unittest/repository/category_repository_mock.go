package repository

import (
	"belajar_4_unittest/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (respository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := respository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entity.Category)
		return &category
		
	}
}