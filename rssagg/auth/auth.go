package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header")
	}

	pair := strings.Split(val, " ")
	if len(pair) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if pair[0] != "ApiKey" {
		return "", errors.New("malformed first part of Authorization header")
	}

	return pair[1], nil
}
