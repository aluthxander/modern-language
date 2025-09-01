package repository

import (
	"belajar-12-resful-api/helper"
	"belajar-12-resful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// Implementation for saving a category
	SQL := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicError(err)

	id, err :=result.LastInsertId()
	helper.PanicError(err)

	category.Id = int(id)
	return category
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// Implementation for updating a category
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicError(err)

	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	// Implementation for deleting a category
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicError(err)
}


func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
    SQL := "SELECT id, name FROM category WHERE id = ?"
    row := tx.QueryRowContext(ctx, SQL, categoryId)

    category := domain.Category{}
    err := row.Scan(&category.Id, &category.Name)

    if err != nil {
        if err == sql.ErrNoRows {
            return category, errors.New("category not found")
        }
        helper.PanicError(err)
    }

    return category, nil
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// Implementation for finding all category
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		categories = append(categories, category)
	}

	return categories
}