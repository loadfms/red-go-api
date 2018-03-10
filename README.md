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
$ ./red-go-api
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

## Tests
### Token
* Retrieve token to be used to authenticate other requests
```sh
curl -X GET   http://localhost:3000/token   -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json'
```

### User
* Retrieve all users in database
```sh
curl -X GET   http://localhost:3000/users   -H 'Authorization: Bearer TOKEN'   -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json'
```

* Retrieve user by id
```sh
curl -X GET   http://localhost:3000/users/{id}   -H 'Authorization: Bearer TOKEN'   -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json'
```

* Create user
```sh
curl -X POST  http://localhost:3000/users -H 'Authorization: Bearer TOKEN' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -d '{"name": "Leonardo", "gravatar": "gravatar.jpg"}'
```

* Update user
```sh
curl -X PUT  http://localhost:3000/users/{id} -H 'Authorization: Bearer TOKEN' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -d '{"name": "Leonardo", "gravatar": "gravatar_2.jpg"}'
```

* Delete user
```sh
curl -X DELETE  http://localhost:3000/users/{id} -H 'Authorization: Bearer TOKEN' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json'
```

### Widget
* Retrieve all widgets in database
```sh
curl -X GET   http://localhost:3000/widgets   -H 'Authorization: Bearer TOKEN'   -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json'
```

* Retrieve widget by id
```sh
curl -X GET   http://localhost:3000/widgets/{id}   -H 'Authorization: Bearer TOKEN'   -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json'
```

* Create widget
```sh
curl -X POST  http://localhost:3000/widgets -H 'Authorization: Bearer TOKEN' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -d '{"name": "First widget", "color": "red", "price": "20", "inventory": 1000,"melts": false}'
```

* Update widget
```sh
curl -X PUT  http://localhost:3000/widgets/{id} -H 'Authorization: Bearer TOKEN' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -d '{"name": "First widget", "color": "blue", "price": "20", "inventory": 1000,"melts": false}'
```

* Delete widget
```sh
curl -X DELETE  http://localhost:3000/widgets/{id} -H 'Authorization: Bearer TOKEN' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json'
```

## Sorry about the mess!
Hello guys! This is my very first app built in Go Lang, so, I'm pretty sure there are many other best ways to do.  
But I will love learn all the amazing things and the right way to deal with this.  
I really loved Go! :)  
Thanks for your attention and I'm looking forward for some positive feedback.  
Best regards  
Leo
