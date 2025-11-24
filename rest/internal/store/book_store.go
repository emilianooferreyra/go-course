package store

import (
	"database/sql"
	"learning-go/rest/internal/model"
)

type Store interface {
	GetAll() ([]*model.Book, error)
	GetByID(id int) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id int, book *model.Book) (*model.Book, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*model.Book, error) {
	q := `SELECT id, title, author, image_url FROM books`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*model.Book, 0)
	for rows.Next() {
		b := model.Book{}
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ImageURL); err != nil {
			return nil, err
		}

		books = append(books, &b)
	}

	return books, nil
}

func (s *store) GetByID(id int) (*model.Book, error) {
	q := `SELECT id, title, author, image_url FROM books WHERE id = ?`

	b := model.Book{}

	err := s.db.QueryRow(q, id).Scan(&b.ID, &b.Title, &b.Author, &b.ImageURL)

	if err != nil {
		return nil, err
	}

	return &b, nil

}

func (s *store) Create(book *model.Book) (*model.Book, error) {
	q := `INSERT INTO books (title, author, image_url) VALUES (?, ?, ?)`

	resp, err := s.db.Exec(q, book.Title, book.Author, book.ImageURL)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.ID = int(id)

	return book, nil

}

func (s *store) Update(id int, book *model.Book) (*model.Book, error) {
	q := `UPDATE books SET title = ?, author = ?, image_url = ? WHERE id = ?`

	_, err := s.db.Exec(q, book.Title, book.Author, book.ImageURL, id)
	if err != nil {
		return nil, err
	}

	book.ID = id

	return book, nil
}

func (s *store) Delete(id int) error {
	q := `DELETE from books WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
