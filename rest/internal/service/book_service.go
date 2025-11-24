package service

import (
	"errors"
	"learning-go/rest/internal/model"
	"learning-go/rest/internal/store"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error) {
	return s.store.GetAll()
}

func (s *Service) GetByID(id int) (*model.Book, error) {
	return s.store.GetByID(id)
}

func (s *Service) CreateBook(book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("necesitamos el titulo")
	}

	return s.store.Create(&book)
}

func (s *Service) UpdateBook(id int, book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("necesitamos el titulo")
	}

	return s.store.Update(id, &book)
}

func (s *Service) DeleteBook(id int) error {
	return s.store.Delete(id)
}
