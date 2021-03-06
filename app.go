package main

import (
	"log"
	"net/http"
	"time"

	. "./routes"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var tokenKey = []byte("secret")

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	//Token
	r.HandleFunc("/token", GetTokenHandler).Methods("GET")

	//User
	r.Handle("/users", jwtMiddleware.Handler(AllUsersEndPoint)).Methods("GET")
	r.Handle("/users", jwtMiddleware.Handler(CreateUserEndPoint)).Methods("POST")
	r.Handle("/users/{id}", jwtMiddleware.Handler(UpdateUserEndPoint)).Methods("PUT")
	r.Handle("/users/{id}", jwtMiddleware.Handler(DeleteUserEndPoint)).Methods("DELETE")
	r.Handle("/users/{id}", jwtMiddleware.Handler(FindUserEndpoint)).Methods("GET")

	//Widget
	r.Handle("/widgets", jwtMiddleware.Handler(AllWidgetsEndPoint)).Methods("GET")
	r.Handle("/widgets", jwtMiddleware.Handler(CreateWidgetsEndPoint)).Methods("POST")
	r.Handle("/widgets/{id}", jwtMiddleware.Handler(UpdateWidgetsEndPoint)).Methods("PUT")
	r.Handle("/widgets/{id}", jwtMiddleware.Handler(DeleteWidgetsEndPoint)).Methods("DELETE")
	r.Handle("/widgets/{id}", jwtMiddleware.Handler(FindWidgetsEndpoint)).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "System Admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(tokenKey)

	w.Write([]byte(tokenString))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
