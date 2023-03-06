package controller

import (
	"context"
	"errors"
	"go-project/model"
	"sync"
)

var storage map[string]model.Books

type Books interface {
	CreateBook(ctx context.Context, book model.Books) error
	GetListOfBook(ctx context.Context) ([]*model.Books, error)
	GetBookByID(ctx context.Context, id string) (model.Books, error)
	UpdateBook(ctx context.Context, book model.Books) error
	DeleteBook(ctx context.Context, id string) error
}

type book struct {
}

func InitStorage() {
	once := sync.Once{}
	once.Do(func() {
		storage = make(map[string]model.Books)
	})
}
func NewBooks() Books {
	return Books(book{})
}

func (b book) CreateBook(ctx context.Context, book model.Books) error {
	storage[book.ID] = book
	return nil
}

func (b book) GetListOfBook(ctx context.Context) ([]*model.Books, error) {
	resp := make([]*model.Books, 0)
	for _, bk := range storage {
		bok := bk
		resp = append(resp, &bok)
	}

	return resp, nil
}

func (b book) GetBookByID(ctx context.Context, id string) (model.Books, error) {
	if bk, ok := storage[id]; ok {
		return bk, nil
	}
	return model.Books{}, errors.New("book not found")
}

func (b book) UpdateBook(ctx context.Context, book model.Books) error {
	if _, ok := storage[book.ID]; !ok {
		return errors.New("book not found")
	}
	storage[book.ID] = book
	return nil
}

func (b book) DeleteBook(ctx context.Context, id string) error {
	if _, ok := storage[id]; !ok {
		return errors.New("book not found")
	}
	delete(storage, id)
	return nil
}
