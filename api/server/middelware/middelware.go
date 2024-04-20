package middelware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func DifferentiationRights(next http.Handler, role string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenStr := req.Header.Get("Token")

		if tokenStr == "" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Token empty"))
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		})

		fmt.Println(token)

		if err != nil || !token.Valid {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Token inthalid"))
			return
		}
		claims, ok := token.Claims(jwt.MapClaims)

		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Token inthalid"))
			return
		}

		tokenRole := claims["role"].(string)

		if tokenRole != role {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Token inthalid"))
			return
		}

	})

}
