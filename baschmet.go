package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	templatesDir = "templates/"
	chartFile    = "Chart.yaml"
	metaFile     = "meta.yaml"
)

type Baschmet struct {
	DryRun  bool
	Charts  []string
	Options *Options
}

func (b *Baschmet) Start() error {
	for _, chartDir := range b.Charts {
		err := b.ProcessChart(chartDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Baschmet) ProcessChart(chartDir string) error {
	if _, err := os.Stat(chartDir); os.IsNotExist(err) {
		return nil
	}
	vars, err := b.GetChartVariables(chartDir)
	if err != nil {
		return err
	}
	fmt.Println(vars.Chart.String())
	rootDir := path.Join(chartDir, "../..")
	err = b.ProcessFiles(rootDir, templatesDir, vars)
	if err != nil {
		return err
	}
	return nil
}

func (b *Baschmet) GetChartVariables(chartDir string) (*Variables, error) {
	chartFile := path.Join(chartDir, chartFile)
	chart, err := ChartFromFile(chartFile)
	if err != nil {
		return nil, err
	}
	metaFile := path.Join(chartDir, "..", metaFile)
	meta, err := MetaFromFile(metaFile)
	if err != nil {
		return nil, err
	}
	return &Variables{
		Options: b.Options,
		Chart:   chart,
		Meta:    meta,
	}, nil
}

func (b *Baschmet) ProcessFiles(rootDir, templatesDir string, vars *Variables) error {
	templFiles, err := FilePathWalkDir(templatesDir)
	if err != nil {
		return err
	}
	dmp := diffmatchpatch.New()
	for _, templFile := range templFiles {
		relPath := strings.TrimPrefix(templFile, templatesDir)
		resultPath := path.Join(rootDir, relPath)
		fmt.Println(resultPath)
		templ, err := GetTemplateText(templFile)
		if err != nil {
			return err
		}
		text, err := GenerateTemplate(templ, "gitlab", vars)
		if err != nil {
			return err
		}
		originalText, err := GetTemplateText(resultPath)
		if err == nil {
			if originalText == text {
				fmt.Println("nothing changed.")
				continue
			}
			diffs := dmp.DiffMain(originalText, text, true)
			fmt.Println(dmp.DiffPrettyText(diffs))
		}
		if b.DryRun {
			continue
		}
		err = ioutil.WriteFile(resultPath, []byte(text), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
