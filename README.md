[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

RESTful API to manage users and widgets written in Go. It uses MongoDB as storage and JSON Web Tokens as authenticantion.

## Dependencies

* [Install Go tools](https://golang.org/doc/install)
* [Install MongoDB](https://www.mongodb.com/download-center#community)

## Instructions

```sh
git clone https://github.com/loadfms/red-go-api.git
```
* open terminal in cloned directory
```sh
$ go get
```
* build
```sh
$ go build
```

### Execute
#### Windows
* execute red-go-api.exe or
```sh
$ go run app.go
```

#### Linux and osX
```sh
$ ./widgets-spa-go-api
```

* api will be running on http://localhost:3000

## Endpoints

### Token
* GET /get-token http://localhost:3000/token

### User
* GET /users http://localhost:3000/users
* GET /users/:id http://localhost:3000/users/:id
* (Bonus!) POST /users http://localhost:3000/users 
* (Bonus!) PUT /users/:id http://localhost:3000/users/:id
* (Bonus!) DELETE /users/:id http://localhost:3000/users/:id

### Widget
* GET /widgets http://localhost:3000/widgets
* GET /widgets/:id http://localhost:3000/widgets/:id
* POST /widgets http://localhost:3000/widgets
* PUT /widgets/:id http://localhost:3000/widgets/:id
* (Bonus!) DELETE /widgets/:id http://localhost:3000/users/:id
