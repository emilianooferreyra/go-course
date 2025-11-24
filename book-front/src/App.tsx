import { useState, useEffect } from "react";
import "./App.css";

interface Book {
  id: number;
  title: string;
  author: string;
  imageUrl: string;
}

function App() {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchBooks();
  }, []);

  const fetchBooks = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await fetch("/api/books");
      if (!response.ok) {
        throw new Error("Error al obtener los libros");
      }
      const data = await response.json();
      setBooks(data || []);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Error desconocido");
      setBooks([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <div className="header">
        <h1>Popular this month</h1>
      </div>

      {loading && <p className="status">Cargando libros...</p>}
      {error && <p className="error">⚠️ {error}</p>}

      {!loading && books.length === 0 && !error && (
        <p className="empty">No hay libros registrados</p>
      )}

      {!loading && books.length > 0 && (
        <div className="books-grid">
          {books.map((book) => (
            <div key={book.id} className="book-card">
              <div className="book-wrapper">
                {book.imageUrl && (
                  <div className="book-image-container">
                    <img
                      src={book.imageUrl}
                      alt={book.title}
                      className="book-image"
                    />
                  </div>
                )}
              </div>
              <button className="add-to-cart">Add to cart</button>
              <div className="book-info">
                <h2>{book.title}</h2>
                <p className="author">By: {book.author}</p>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default App;
