package main

import (
	"fmt"
	"os"
	"strconv"
)

func realMain(dryRun, incChartVersion bool) int {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Nothing to do.")
		return 0
	}
	options, err := OptionsFromFile("options.yaml")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	baschmet := &Baschmet{
		Options:         options,
		DryRun:          dryRun,
		IncChartVersion: incChartVersion,
		Charts:          args,
	}
	err = baschmet.Start()
	if err != nil {
		fmt.Println(err)
		return 2
	}
	return 0
}

func parseBoolEnv(env string, def bool) (bool, error) {
	var result bool
	resultRaw := os.Getenv(env)
	if len(resultRaw) == 0 {
		result = def
	} else {
		b, err := strconv.ParseBool(resultRaw)
		if err != nil {
			return false, err
		}
		result = b
	}
	return result, nil
}

func main() {
	dryRun, err := parseBoolEnv("DRY_RUN", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	incChartVersion, err := parseBoolEnv("INC_CHART_VERSION", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(realMain(dryRun, incChartVersion))
}
