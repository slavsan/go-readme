package cmd

import (
	"io/ioutil"

	"github.com/slavsan/go-readme/internal/app"
	"github.com/spf13/cobra"
)

var yamlToMdCmd = &cobra.Command{
	Use:   "yaml-to-md",
	Short: "Generate README.md from readme.yaml",
	Long:  "Generate README.md from readme.yaml\n",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := ioutil.ReadFile("./readme.yaml")
		if err != nil {
			return err
		}

		r := &app.Readme{}

		markdown, err := r.Parse(string(data))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("./README.md", []byte(markdown), 0644)
		if err != nil {
			return err
		}

		return nil
	},
}
