package action

import (
	"foxy-ci/platform/pipeline"
)

// Run clones and runs a github repo
func Run(gitURL string) {
	task := pipeline.Create(gitURL)
	task.Clone()
	task.Run()
}
