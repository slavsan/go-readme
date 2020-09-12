package cmd

import (
	"io/ioutil"

	"github.com/slavsan/go-readme/internal/app"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate example readme.yaml",
	Long:  "Generate example readme.yaml\n",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := &app.Readme{}

		exampleYaml := r.ExampleYaml()

		err := ioutil.WriteFile("./README.md", []byte(exampleYaml), 0644)
		if err != nil {
			return err
		}

		return nil
	},
}
