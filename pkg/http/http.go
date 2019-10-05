package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func GetBearerTokenFromHeader(w http.ResponseWriter, r *http.Request, returnHTTPError bool) (string, error) {
	token := r.Header.Get("Authorization")
	if token == ""  {
		if  returnHTTPError {
			http.Error(w, "Unauthorized.", 401)
		}
		return "", errors.New("No token in the Authorization Header")
	}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) < 2 {
		if  returnHTTPError {
			http.Error(w, "Unauthorized.", 401)
		}
		return "", errors.New("No token in the Authorization Header")
	}

	token = strings.TrimSpace(splitToken[1])
	return token, nil
}

func MarshalAndWriteJSON(w http.ResponseWriter, content interface{}) error {
	js, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJSON(w, js)
	return nil
}

func WriteJSON(w http.ResponseWriter, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
