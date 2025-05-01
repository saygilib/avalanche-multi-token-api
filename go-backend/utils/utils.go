package utils

import (
	"encoding/json"
)

func StringifyJSON(data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}
