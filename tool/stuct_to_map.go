package tool

import "encoding/json"

// StructToMap will encode a structure in a simple map
func StructToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(s)
	_ = json.Unmarshal(j, &m)

	return m
}
