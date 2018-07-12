package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

func GetTemplateText(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func GenerateTemplate(templ, name string, data interface{}) (string, error) {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}
	buf := bytes.NewBufferString("")
	if messageTempl, err := template.New(name).Funcs(funcMap).Parse(templ); err != nil {
		return "", fmt.Errorf("failed to parse template for %s: %v", name, err)
	} else if err := messageTempl.Execute(buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template for %s: %v", name, err)
	}
	return buf.String(), nil
}
