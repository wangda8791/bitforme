package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ReadJSON(id string) []byte {

	filename, _ := filepath.Abs(fmt.Sprintf("../jsons/%s.json", id))
	byteValue, _ := ioutil.ReadFile(filename)
	return byteValue
}

func JSONToByteArray(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

func ByteArrayToJSON(data []byte, v interface{}) {
	json.Unmarshal(data, v)
}

func InterfaceToJSON(v interface{}, result interface{}) {
	v_byte, _ := json.Marshal(v)
	err := json.Unmarshal(v_byte, result)
	if err != nil {
		panic(err)
	}
}

func ReadChartJson(id string) []byte {
	
	filename, _ := filepath.Abs(fmt.Sprintf("../jsons/charts/%s.json", id))
	byteValue, _ := ioutil.ReadFile(filename)
	return byteValue
}
