# First API in Go

This is a simple Go-based API that serves as an introductory project for building RESTful services in Go. The project demonstrates basic API functionality using best practices to create a maintainable.

## Features

- Basic REST API implementation
- Uses standard Go libraries with minimal external dependencies

## Installation

To get started with the project, you need to have Go installed on your machine.

1. Clone the repository:

   ```bash
   git clone https://github.com/ikidoncc/first-api-go.git
   cd first-api-go
   ```

2. Install the required dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run cmd/server/main.go
   ```
   or just run with air:
   ```bash
   air
   ```

   The server will start on the default port (`3000`).

## Project Structure

The project follows a simple structure:

```
first-api-go/
│
├── cmd/
│   └── server/          # Main application entry point
│
├── internal/
│   └── app/             # Business logic and app-related functionality
│
├── .air.toml            # Configuration for live reloading
├── .gitignore           # Git ignore file
├── go.mod               # Go module dependencies
├── go.sum               # Go checksum dependencies
└── .tool-versions       # Tool versions
```

## Endpoints

The API exposes a set of simple endpoints. Here are the basics:

| **Method** | **Endpoint**    | **Description** |
|------------|-----------------|-----------------|
| GET        | /api/users      | Returns an array of all users. |
| GET        | /api/users/{id} | Returns the user object with the specified ID. |
| PUT        | /api/users/{id} | Updates the user with the specified ID using the data from the request body and returns the modified user. |
| POST       | /api/users      | Creates a new user using the information provided in the request body. |
| DELETE     | /api/users/{id} | Removes the user with the specified ID and returns the deleted user. |


## License

This project is licensed under the MIT License.