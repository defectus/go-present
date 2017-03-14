package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func templateThis(templ string, params map[string]interface{}) string {
	if tmplt, err := template.New("").Parse(templ); err == nil {
		buff := bytes.NewBufferString("")
		if err := tmplt.Execute(buff, params); err == nil {
			appliedTemplate := buff.String()
			return appliedTemplate
		}
	}
	return ""
}

func tryTemplate() {
	template := `loop {{range $index, $value := .loop}}{{$value}},{{end}} until you drop`
	data := map[string]interface{}{"loop": []string{"1", "2", "3"}}
	fmt.Println(templateThis(template, data))
}
