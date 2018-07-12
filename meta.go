package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Meta map[string]interface{}

func MetaFromFile(pat string) (*Meta, error) {
	b, err := ioutil.ReadFile(pat)
	if err != nil {
		return nil, err
	}
	meta := &Meta{}
	yaml.Unmarshal(b, &meta)
	if err != nil {
		return nil, err
	}
	return meta, nil
}
