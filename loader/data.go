package loader

type Panel struct {
	Name        string `yaml:"name"`
	Heading     string `yaml:"heading"`
	Description string `yaml:"description"`
}

type Tab struct {
	Name   string  `yaml:"name"`
	Label  string  `yaml:"label"`
	Panels []Panel `yaml:"panels"`
}

type Button struct {
	Name   string  `yaml:"name"`
	Label  string  `yaml:"label"`
	Tabs   []Tab   `yaml:"tabs,omitempty"`
	Panels []Panel `yaml:"panels,omitempty"`
}
