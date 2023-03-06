package controller

import (
	"context"
	"fmt"
	"go-project/model"
	"testing"
)

func TestCreateBook(t *testing.T) {
	testCases := map[string]struct {
		book model.Books
	}{
		"create book success": {
			book: model.Books{
				ID:          "123",
				Title:       "purple hibiscus",
				Author:      "azeez",
				PublishDate: "2023-03-06",
			},
		},
	}
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			booksInterface := NewBooks()
			err := booksInterface.CreateBook(context.Background(), test.book)
			fmt.Println("err", err)
		})
	}
}
