package main

import (
	"fmt"
	"os"
	"strconv"
)

func realMain(dryRun bool) int {
	options, err := OptionsFromFile("options.yaml")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Nothing to do.")
		return 0
	}
	baschmet := &Baschmet{
		Options: options,
		DryRun:  dryRun,
		Charts:  args,
	}
	err = baschmet.Start()
	if err != nil {
		fmt.Println(err)
		return 2
	}
	return 0
}

func main() {
	var dryRun bool
	dryRunRaw := os.Getenv("DRY_RUN")
	if len(dryRunRaw) == 0 {
		dryRun = true
	} else {
		b, err := strconv.ParseBool(dryRunRaw)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dryRun = b
	}
	os.Exit(realMain(dryRun))
}
