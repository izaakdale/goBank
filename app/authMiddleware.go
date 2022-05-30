package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/izaakdale/goBank/domain"
)

type AuthMiddleware struct {
	repo domain.AuthRepo
}

func (authM AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			currentRoute := mux.CurrentRoute(request)
			currentRouteVars := mux.Vars(request)
			authHead := request.Header.Get("Authorization")

			if authHead != "" {
				token := getTokenFromHeader(authHead)

				isAuthorized := authM.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(writer, request)
				} else {
					writer.Header().Add("Content-Type", "application/json")
					writer.WriteHeader(http.StatusForbidden)
					if err := json.NewEncoder(writer).Encode("Insufficient priviledges"); err != nil {
						panic(err)
					}
				}
			} else {
				writer.Header().Add("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)
				if err := json.NewEncoder(writer).Encode("Missing token"); err != nil {
					panic(err)
				}
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer ")
	return splitToken[1]
}
