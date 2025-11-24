# Ejemplos de Consultas API REST

## Obtener todos los libros
```bash
# GET - Obtener todos los libros
curl -X GET http://localhost:8080/books
```

## Obtener un libro por ID
```bash
# GET - Obtener un libro específico por su ID
curl -X GET http://localhost:8080/books/1
```

## Crear un nuevo libro
```bash
# POST - Crear un nuevo libro con titulo y autor
curl -X POST -H "Content-type: application/json" -d '{"title": "Go Programing", "author": "John Doe"}' http://localhost:8080/books
```

## Actualizar un libro existente
```bash
# PUT - Actualizar un libro existente (por ID) con nuevos datos
curl -X PUT -H "Content-type: application/json" -d '{"title": "Advanced Go Programing", "author": "JANE Doe"}' http://localhost:8080/books/1
```

## Eliminar un libro
```bash
# DELETE - Eliminar un libro específico por su ID
curl -X DELETE http://localhost:8080/books/1
```

---

## Explicación de Headers y Parámetros

- `-X`: Especifica el método HTTP (GET, POST, PUT, DELETE)
- `-H`: Define los headers de la solicitud (en este caso, el tipo de contenido JSON)
- `-d`: Los datos a enviar en el body de la solicitud (para POST y PUT)
- `Content-type: application/json`: Indica que estamos enviando datos en formato JSON
