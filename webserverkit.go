package webserverkit

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
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

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SliceContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CopyStringSlice(slice []string) []string {
	return append([]string(nil), slice...)
}

func RemoveStringFromSlice(s []string, r string) []string {
	found := false
	j := 0

	copiedSlice := CopyStringSlice(s)
	for i, v := range copiedSlice {
		if v == r {
			found = true
			j = i
		}
	}

	if found {
		return append(copiedSlice[:j], copiedSlice[j+1:]...)
	}
	return s
}

// min: inclusive, max: exclusive
func RandomNumberRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func Percent(p int) bool {
	val := RandomNumberRange(0, 100)
	return p >= val
}

func PickRandomInt(values []int) int {
	i := RandomNumberRange(0, len(values))
	return values[i]
}

func PickRandomString(values []string) string {
	i := RandomNumberRange(0, len(values))
	return values[i]
}

func fieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))
	for _, s := range fields {
		set[s] = true
	}
	return set
}

func SelectFields(obj interface{}, fields ...string) map[string]interface{} {
	fs := fieldSet(fields...)
	rt, rv := reflect.TypeOf(obj), reflect.ValueOf(obj)
	out := make(map[string]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")

		keys := strings.Split(jsonKey, ",")
		if len(keys) > 0 {
			jsonKey = keys[0]
		}

		if fs[jsonKey] {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}
	return out
}

