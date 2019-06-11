package app

import (
	"net/http"
	u "github.com/Go_Rest_Api/utils"
	"strings"
	"github.com/Go_Rest_Api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	// "context"
	// "fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		notAuth := []string{"/api/user/new", "api/user/login"}
		requestPath := r.URL.Path

		for _ , value :=range notAuth{

			if value == requestPath{
				next.ServeHTTP(w,r)
				return
			}
		}

		response := make(map[string]interface{})

		tokenHeader := r.Header.Get("Authorization")

		// Missing token header
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// 
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response =u.Message(false, "Invalid auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// Grab the token part, what we are truly interest
		tokenPart := splitted[1]
		tk := &models.Token{}

		// 
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token)(interface{},error){
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Respond(w,response)
			return 
		}

		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Respond(w, response)
			return
		}

		// IF everything is ok, procceed with the request and set teh caller to the user
		// retrieved from the parsed token
		// fmt.Sprint("User", tk.Username)
		// ctx := context.WithValue(r.Context(),"user", tk.UserId)
		// r = r.WithContext(ctx)
		// next.ServeHTTP(w, r)
	})
}