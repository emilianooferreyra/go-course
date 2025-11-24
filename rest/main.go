package main

import (
	"database/sql"
	"fmt"
	"learning-go/rest/internal/service"
	"learning-go/rest/internal/store"
	"learning-go/rest/internal/transport"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Conectar a SQLLite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear el table si no existe
	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inyectar nuestra dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// Configurar las rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Servidor ejecutandose en http:localhost: 8080")
	fmt.Println("API Endpoints:")
	fmt.Println("GET 		/books - 			Obtener todos los libros")
	fmt.Println("POST 	/books - 			Crear un nuevo libro")
	fmt.Println("GET 		/books/{id} - Obtener un libro en especifico")
	fmt.Println("PUT 		/books/{id} - Obtener todos los libros")
	fmt.Println("DELETE /books/{id} - Eliminar un libro")

	// Empezar y escuchar el servidor
	log.Fatal(http.ListenAndServe(":8080", nil))

}
