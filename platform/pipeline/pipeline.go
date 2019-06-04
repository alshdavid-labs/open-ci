package pipeline

import (
	"io/ioutil"
	"os/exec"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

type Pipeline struct {
	ID     string
	GitURL string
}

func Create(gitURL string) *Pipeline {
	return &Pipeline{
		ID:     uuid.New().String(),
		GitURL: gitURL,
	}
}

func (p *Pipeline) Clone() {
	Cmd("mkdir -p tmp")
	Cmd("git clone --depth=1 " + p.GitURL + " tmp/" + p.ID)
	Cmd("rm -rf tmp/" + p.ID + "/.git")
}

func (p *Pipeline) Run() {
	configFile, _ := ioutil.ReadFile("tmp/" + p.ID + "/.foxy-ci.yml")
	config := &Config{}
	yaml.Unmarshal(configFile, config)

	for name, action := range config.Actions {
		runAction(name, action)
	}
}

func runAction(name string, action Action) {
	ID := uuid.New().String()
	Cmd("docker run --name " + ID + " -t -d " + action.Image)
	for _, step := range action.Steps {
		Cmd("docker exec " + ID + " " + step)
	}
	Cmd("docker stop " + ID)
	Cmd("docker rm" + ID)
}

func cleanUp(ID string) {
	Cmd("rm -rf tmp/" + ID)
}

func Cmd(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
