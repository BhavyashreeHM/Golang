package jwt

import (
	"go/token"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super_secret_key")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func generaejwt(username string, password string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(secretKey)

}

func jwtmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){

	next.ServeHTTP(w, r)	
	})

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func UserHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)
	mux.Handle("/user", jwtmiddleware(http.HandlerFunc(UserHandler)))

	log.Println("Server Up and Running on Por :8080")
	http.ListenAndServe(":8080", mux)

}
