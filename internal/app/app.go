package app

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// Readme ..
type Readme struct {
}

type doc struct {
	Readme readme `yaml:"readme"`
}

type setup struct {
	Clone   string `yaml:"clone"`
	Build   string `yaml:"build"`
	Start   string `yaml:"start"`
	Testing string `yaml:"testing"`
	Linting string `yaml:"linting"`
}

type readme struct {
	Title         string `yaml:"title"`
	Description   string `yaml:"description"`
	Prerequisites string `yaml:"prerequisites"`
	Install       string `yaml:"install"`
	Usage         string `yaml:"usage"`
	Setup         *setup `yaml:"setup"`
	License       string `yaml:"license"`
	Author        string `yaml:"author"`
}

func (d *doc) ToMarkdown() string {
	var buf strings.Builder

	if d.Readme.Title != "" {
		buf.WriteString(fmt.Sprintf("# %s", d.Readme.Title))
	}

	if d.Readme.Description != "" {
		buf.WriteString(fmt.Sprintf("%s", d.Readme.Description))
	}

	if d.Readme.Install != "" {
		buf.WriteString(fmt.Sprintf("### Installation\n%s", d.Readme.Install))
	}

	if d.Readme.Usage != "" {
		buf.WriteString(fmt.Sprintf("### Usage\n%s", d.Readme.Usage))
	}

	if d.Readme.Setup != nil {
		buf.WriteString(fmt.Sprintf("### Setup\n\n"))

		if d.Readme.Setup.Clone != "" {
			buf.WriteString(fmt.Sprintf("#### Clone\n"))
			buf.WriteString(fmt.Sprintf("%s", d.Readme.Setup.Clone))
		}

		if d.Readme.Setup.Build != "" {
			buf.WriteString(fmt.Sprintf("#### Build\n"))
			buf.WriteString(fmt.Sprintf("%s", d.Readme.Setup.Build))
		}

		if d.Readme.Setup.Start != "" {
			buf.WriteString(fmt.Sprintf("#### Start\n"))
			buf.WriteString(fmt.Sprintf("%s", d.Readme.Setup.Start))
		}

		if d.Readme.Setup.Testing != "" {
			buf.WriteString(fmt.Sprintf("#### Testing\n"))
			buf.WriteString(fmt.Sprintf("%s", d.Readme.Setup.Testing))
		}

		if d.Readme.Setup.Linting != "" {
			buf.WriteString(fmt.Sprintf("#### Linting\n"))
			buf.WriteString(fmt.Sprintf("%s", d.Readme.Setup.Linting))
		}
	}

	if d.Readme.Prerequisites != "" {
		buf.WriteString(fmt.Sprintf("### Prerequisites\n%s", d.Readme.Prerequisites))
	}

	if d.Readme.License != "" {
		buf.WriteString(fmt.Sprintf("### License\n%s", d.Readme.License))
	}

	if d.Readme.Author != "" {
		buf.WriteString(fmt.Sprintf("### Author\n%s", d.Readme.Author))
	}

	return buf.String()
}

// Parse ..
func (r *Readme) Parse(content string) (string, error) {
	var d doc

	err := yaml.Unmarshal([]byte(content), &d)
	if err != nil {
		return "", err
	}

	return d.ToMarkdown(), nil
}
