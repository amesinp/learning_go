package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/amesinp/learning_go/todo_app_api/utils"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Todo API v1"})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware validates a user's access token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 8 {
			utils.HandleAuthenticationError(w)
			return
		}

		tokenString := authHeader[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		checkRefExpiry := false

		if err != nil || !token.Valid {
			// If the token has expired and the user is trying to refresh then allow it
			if r.RequestURI == RefreshTokenURI && err.Error() == "Token is expired" {
				checkRefExpiry = true
			} else {
				utils.HandleAuthenticationError(w)
				return
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && (!checkRefExpiry || (checkRefExpiry && isRefreshTokenValid(claims))) {
			ctx := context.WithValue(r.Context(), utils.TokenClaimsKey, utils.TokenClaims{
				UserID:         int(claims["sub"].(float64)),
				RefreshTokenID: int(claims["ref"].(float64)),
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			utils.HandleAuthenticationError(w)
		}
	})
}

func isRefreshTokenValid(claims jwt.MapClaims) bool {
	refreshTokenExpiry := time.Unix(int64(claims["refExp"].(float64)), 0)
	return refreshTokenExpiry.Sub(time.Now()) > 0
}

// ConfigureRouter sets up the application routes
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)

	authRoutes(router)
	todoRoutes(router)

	router.HandleFunc("/", index)

	return router
}
