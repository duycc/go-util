package util

import (
	"encoding/json"
)

func StructToMap(object interface{}) (map[string]interface{}, error) {
	var (
		resMap map[string]interface{}
		err    error
	)
	if object == nil {
		return resMap, nil
	}
	data, err := json.Marshal(object)
	if err != nil {
		return resMap, err
	}
	err = json.Unmarshal(data, &resMap)
	return resMap, err
}
