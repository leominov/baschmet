package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Options map[string]interface{}

func OptionsFromFile(pat string) (*Options, error) {
	b, err := ioutil.ReadFile(pat)
	if err != nil {
		return nil, err
	}
	options := &Options{}
	yaml.Unmarshal(b, &options)
	if err != nil {
		return nil, err
	}
	return options, nil
}
