package util

import (
	"encoding/json"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 64)
	return int(num)
}

func StrToInt64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}

func StrToInt64Slice(str, split string) []int64 {
	strs := strings.Split(str, split)
	res := []int64{}
	for _, v := range strs {
		num, _ := strconv.ParseInt(v, 10, 64)
		res = append(res, num)
	}
	return res
}

func StrToStrSlice(str, split string) []string {
	return strings.Split(str, split)
}

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
