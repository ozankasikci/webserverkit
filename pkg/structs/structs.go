package structs

import (
	"reflect"
	"strings"
)

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
