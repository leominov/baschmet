package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Options map[string]interface{}

func OptionsFromFile(pat string) (*Options, error) {
	b, err := ioutil.ReadFile(pat)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	options := &Options{}
	yaml.Unmarshal(b, &options)
	if err != nil {
		return nil, err
	}
	return options, nil
}
