package recipe

type Recipe struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
}

type Metadata struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}
