package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	valueToJSON := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{Key: "my key", Value: "my value"}
	jason, _ := json.Marshal(valueToJSON)
	fmt.Println("json is: ", string(jason))
	restored := &map[string]interface{}{}
	json.Unmarshal(jason, restored)
	fmt.Printf("json is: %#v\n", *restored)
}
