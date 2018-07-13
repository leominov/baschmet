package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Options struct {
	HelmVersion               string                 `yaml:"helmVersion"`
	NodeVersion               string                 `yaml:"nodeVersion"`
	MetaChartVersion          string                 `yaml:"metaChartVersion"`
	PushDockerImageWithLatest bool                   `yaml:"pushDockerImageWithLatest"`
	GoogleCloudSDKVersion     string                 `yaml:"googleCloudSDKVersion"`
	BaseURLDomain             string                 `yaml:"baseURLDomain"`
	BaseURLProtocol           string                 `yaml:"baseURLProtocol"`
	PublicHelmRepo            bool                   `yaml:"publicHelmRepo"`
	StaticEnvironments        []string               `yaml:"staticEnvironments"`
	DynamicEnvironments       bool                   `yaml:"dynamicEnvironments"`
	DisabledServiceTestsRaw   []string               `yaml:"disabledServiceTests"`
	Custom                    map[string]interface{} `yaml:"custom"`
	GeneratedBy               string                 `yaml:"generatedBy"`
	DisabledServiceTests      map[string]bool        `yaml:"-"`
}

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
	options.DisabledServiceTests = make(map[string]bool)
	for _, service := range options.DisabledServiceTestsRaw {
		options.DisabledServiceTests[service] = true
	}
	return options, nil
}
