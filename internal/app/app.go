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

type readme struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func (d *doc) ToMarkdown() string {
	var buf strings.Builder

	if d.Readme.Title != "" {
		buf.WriteString(fmt.Sprintf("# %s", d.Readme.Title))
	}

	if d.Readme.Description != "" {
		buf.WriteString(fmt.Sprintf("%s", d.Readme.Description))
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
