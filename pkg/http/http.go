package http

import (
	"encoding/json"
	"net/http"
)

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
