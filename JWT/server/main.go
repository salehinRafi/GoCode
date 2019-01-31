package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("zanderazishere")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Super Secret Information")
}

// isAuthorized is a middleware that check to see that the incoming request features the Token header in the request and we then subsequently check to see if the token is valid based off our private mySigningKey.
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				// Check if token is valid
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprint(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})

}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("Hello, server!")
	handleRequests()
}
