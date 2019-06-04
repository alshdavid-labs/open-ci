package pipeline

type Config struct {
	Actions map[string]Action `yaml:"actions"`
}

type Action struct {
	Image string   `yaml:"image"`
	Steps []string `yaml:"steps"`
}
