Berikut adalah template **README.md** yang dapat Anda gunakan untuk proyek GitHub Anda:

```markdown
# Bookstore API

A simple Bookstore REST API built using Go and Gorilla Mux with SQLite for database management. This API allows users to manage books by performing CRUD (Create, Read, Update, Delete) operations.

## Features

- Add new books
- Retrieve all books or a specific book by ID
- Update book information
- Delete books

## Requirements

- [Go](https://golang.org/doc/install) 1.16+
- [Gorilla Mux](https://github.com/gorilla/mux)
- [GORM (Go ORM)](https://gorm.io/)
- SQLite3

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/bookstore-api.git
   cd bookstore-api
   ```

2. Install required dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

   The API will be running on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint          | Description              |
|--------|-------------------|--------------------------|
| GET    | `/books`          | Get all books            |
| GET    | `/books/{id}`     | Get a specific book by ID |
| POST   | `/books`          | Create a new book         |
| PUT    | `/books/{id}`     | Update a book by ID       |
| DELETE | `/books/{id}`     | Delete a book by ID       |

### Example Requests

#### Get all books

```bash
GET /books
```

Response:

```json
[
    {
        "id": 1,
        "title": "The Go Programming Language",
        "author": "Alan A. A. Donovan",
        "price": 35.99
    },
    {
        "id": 2,
        "title": "Learning Go",
        "author": "Jon Bodner",
        "price": 25.50
    }
]
```

#### Create a new book

```bash
POST /books
```

Request Body:

```json
{
    "title": "New Book Title",
    "author": "Author Name",
    "price": 12.99
}
```

Response:

```json
{
    "id": 3,
    "title": "New Book Title",
    "author": "Author Name",
    "price": 12.99
}
```

#### Update a book

```bash
PUT /books/1
```

Request Body:

```json
{
    "title": "Updated Book Title",
    "author": "Updated Author Name",
    "price": 15.99
}
```

Response:

```json
{
    "id": 1,
    "title": "Updated Book Title",
    "author": "Updated Author Name",
    "price": 15.99
}
```

#### Delete a book

```bash
DELETE /books/1
```

Response:

```
204 No Content
```

## Project Structure

```bash
.
├── main.go                # Entry point of the application
├── go.mod                 # Go module file
├── go.sum                 # Dependencies version management
├── pkg
│   ├── config
│   │   └── app.go         # Database connection and configuration
│   ├── controllers
│   │   └── bookcontroller.go # Handlers for book operations
│   ├── models
│   │   └── models.go      # Book model definition
│   ├── routes
│   │   └── routes.go      # API route definitions
│   └── utils
│       └── utils.go       # Utility functions for handling errors, JSON, etc.
├── bookstore.db            # SQLite database (auto-generated)
└── README.md               # Project documentation
```

## How to Contribute

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Anda dapat menyesuaikan detail seperti URL repository dan nama proyek sesuai dengan kebutuhan. Template ini sudah mencakup informasi penting tentang proyek, cara instalasi, dan contoh penggunaan API.
