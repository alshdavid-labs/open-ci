package main

import (
	"flag"
	"foxy-ci/cmd/cli/action"
)

func main() {
	run := flag.String("run", "", "Repo URL")
	flag.Parse()

	if *run != "" {
		action.Run(*run)
	}
}
