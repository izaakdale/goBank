package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type AuthRepo interface {
	IsAuthorized(string, string, map[string]string) bool
}

type RemoteAuthRepo struct {
}

func NewAuthRepo() RemoteAuthRepo {
	return RemoteAuthRepo{}
}

func (remoteAuthRepo RemoteAuthRepo) IsAuthorized(token string, routeName string, routeVars map[string]string) bool {

	url := buildVerifyUrl(token, routeName, routeVars)

	if response, err := http.Get(url); err != nil {
		fmt.Println("Error while sending verify request " + err.Error())
		return false
	} else {
		// here from the url we want a list of Resources, and whether the current user allowed to access them
		verifyMap := map[string]bool{}
		if err := json.NewDecoder(response.Body).Decode(&verifyMap); err != nil {
			fmt.Println("Error decoding verification " + err.Error())
			return false
		}
		return verifyMap["is_authorized"]
	}
}

func buildVerifyUrl(token string, routeName string, routeVars map[string]string) string {

	host := os.Getenv("AUTH_SERVER_URL")
	path := "/auth/verify"
	scheme := "http"

	url := url.URL{Host: host, Path: path, Scheme: scheme}
	query := url.Query()

	query.Add("token", token)
	query.Add("routeName", routeName)
	for k, v := range routeVars {
		query.Add(k, v)
	}
	url.RawQuery = query.Encode()
	return url.String()
}
