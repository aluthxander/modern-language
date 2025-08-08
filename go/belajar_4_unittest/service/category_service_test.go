package service

import (
	"belajar_4_unittest/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}



func TestCategoryService_Get(t *testing.T){

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, error := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, error)
}