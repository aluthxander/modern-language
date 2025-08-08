package repository

import (
	"belajar-7-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentsRepositoryImpl struct {
	DB *sql.DB
}

func (repository *commentsRepositoryImpl) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error){
	script := "INSERT INTO comments(email, comment) values(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int(id)
	return comment, nil
}

func (repository *commentsRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comments, error){
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	repository.DB.QueryContext(ctx, script, id)
	
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comments{}
	if err != nil {
		return entity.Comments{}, err
	}
	
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		// tidak ada
		return entity.Comments{}, errors.New("comments "+ strconv.Itoa(int(id))+" is not found")
	}
}

func NewCommentsRepository(db *sql.DB) CommentsRepository {
	return &commentsRepositoryImpl{DB: db}
}

func (repository *commentsRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error){
	script := "SELECT id, email, comment FROM comments"
	repository.DB.QueryContext(ctx, script)
	
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments[]entity.Comments
	for rows.Next() {
		comment := entity.Comments{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		
		comments = append(comments, comment)
	}
	return comments, nil
}