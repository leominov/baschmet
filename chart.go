package main

import (
	"fmt"
	"io/ioutil"

	"github.com/coreos/go-semver/semver"
	yaml "gopkg.in/yaml.v2"
)

type Chart struct {
	Name          string            `yaml:"name,omitempty"`
	Version       string            `yaml:"version,omitempty"`
	Description   string            `yaml:"description,omitempty"`
	Keywords      []string          `yaml:"keywords,omitempty"`
	Home          string            `yaml:"home,omitempty"`
	Sources       []string          `yaml:"sources,omitempty"`
	Maintainers   []*Maintainer     `yaml:"maintainers,omitempty"`
	TillerVersion string            `yaml:"tillerVersion,omitempty"`
	Annotations   map[string]string `yaml:"annotations,omitempty"`
	Tags          string            `yaml:"tags,omitempty"`
	Engine        string            `yaml:"engine,omitempty"`
	Icon          string            `yaml:"icon,omitempty"`
	Condition     string            `yaml:"condition,omitempty"`
	KubeVersion   string            `yaml:"kubeVersion,omitempty"`
	APIVersion    string            `yaml:"apiVersion,omitempty"`
	Deprecated    bool              `yaml:"deprecated,omitempty"`
	filePath      string            `yaml:"-"`
}

type Maintainer struct {
	Name  string `yaml:"name,omitempty"`
	Email string `yaml:"email,omitempty"`
	URL   string `yaml:"url,omitempty"`
}

func ChartFromFile(pat string) (*Chart, error) {
	b, err := ioutil.ReadFile(pat)
	if err != nil {
		return nil, err
	}
	chart := &Chart{}
	yaml.Unmarshal(b, &chart)
	if err != nil {
		return nil, err
	}
	chart.filePath = pat
	return chart, nil
}

func (c *Chart) IncPatch() {
	v := semver.New(c.Version)
	v.Patch++
	c.Version = v.String()
}

func (c *Chart) Save() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.filePath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chart) String() string {
	return fmt.Sprintf("[%s] %s", c.Version, c.Name)
}
