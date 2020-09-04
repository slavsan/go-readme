package main

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/slavsan/go-readme/internal/app"
)

type world struct {
	readme            *app.Readme
	yaml              string
	markdown          string
	generatedMarkdown string
	err               error
}

var w *world

func initWorld() *world {
	return &world{
		readme:            &app.Readme{},
		yaml:              "",
		markdown:          "",
		generatedMarkdown: "",
		err:               nil,
	}
}

func aReadmeyamlWithContents(arg1 *messages.PickleStepArgument_PickleDocString) error {
	w.yaml = arg1.Content
	return nil
}

func iParseTheFile() error {
	var err error
	w.generatedMarkdown, err = w.readme.Parse(w.yaml)
	return err
}

func theFollowingREADMEmdWillBeGenerated(arg1 *messages.PickleStepArgument_PickleDocString) error {
	w.markdown = arg1.Content

	if w.markdown != w.generatedMarkdown {
		return fmt.Errorf(
			"failed to parse file correctly\n\twant: %v\n\t got: %v",
			w.markdown, w.generatedMarkdown,
		)
	}

	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// ..
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		w = initWorld()
	})

	ctx.Step(`^a readme\.yaml with contents$`, aReadmeyamlWithContents)
	ctx.Step(`^I parse the file$`, iParseTheFile)
	ctx.Step(`^the following README\.md will be generated$`, theFollowingREADMEmdWillBeGenerated)
}
