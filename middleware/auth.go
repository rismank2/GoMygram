package middleware

import (
	"MyGram/models"
	"MyGram/service"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var tempKey = &tempcontext{"user"}

type tempcontext struct {
	data string
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		servic := service.NewUserService()
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		fmt.Println(splitToken)
		if len(splitToken) > 1 {
			reqToken = splitToken[1]
			fmt.Println("Auth token")
			temp_id, err := servic.VerifToken(reqToken)
			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			}
			fmt.Print(" id: ")
			fmt.Println(temp_id)
			user := models.User{Id: int(temp_id)}

			ctx := context.WithValue(r.Context(), tempKey, &user)
			r = r.WithContext(ctx)

		} else {
			w.Write([]byte(fmt.Sprint(errors.New("token cannot be empty"))))
		}
		next.ServeHTTP(w, r)

	})
}

func RunUser(ctx context.Context) *models.User {
	rawData, _ := ctx.Value(tempKey).(*models.User)
	return rawData
}
