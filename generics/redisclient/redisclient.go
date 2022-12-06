package redisclient

import (
	"encoding/json"
)

var persons string = `[
		{"Name": "John", "Age": 20},
		{"Name": "Jane", "Age": 21}
	]
`

var employees string = `[
		{"Name": "John", "Position": "Manager"},
		{"Name": "Jane", "Position": "Developer"}
	]
`
var db map[string]string = map[string]string{
	"persons":   persons,
	"employees": employees,
}

func Read[T any](key string) (T, error) {
	var data T

	result := db[key]

	err := json.Unmarshal([]byte(result), &data)
	return data, err
}
