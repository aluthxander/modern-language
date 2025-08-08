package repository

import "belajar_4_unittest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}