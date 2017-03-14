package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	Text   string
	Status int `json:"status"`
	secret string
}

type ComposedStruct struct {
	MyStruct
	Timestamp time.Time
}

func main() {
	instance1 := ComposedStruct{MyStruct: MyStruct{Text: "text"}, Timestamp: time.Now()}
	instance1.Status = 200
	adHoc := struct{ Field1 int }{Field1: 1}
	fmt.Printf("instance1:%#v;adHoc:%#v", instance1, adHoc)
}
