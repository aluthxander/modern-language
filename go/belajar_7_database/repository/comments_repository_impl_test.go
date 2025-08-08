package repository

import (
	belajar_7_database "belajar-7-database"
	"belajar-7-database/entity"
	"context"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T){
	commentRepository := NewCommentsRepository(belajar_7_database.GetConnection())
	
	ctx := context.Background()
	comment := entity.Comments{
		Email: "9E9tV@example.com",
		Comment: "Hello World 1",
	}

	result, err :=commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T){
	commentRepository := NewCommentsRepository(belajar_7_database.GetConnection())
	
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}