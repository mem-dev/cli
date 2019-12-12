package cmd

import (
	"fmt"
	"github.com/mem-dev/cli/api"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Snippet() bool {
	questions := []*survey.Question{
		{
			Name:     "title",
			Prompt:   &survey.Input{Message: "I just learned how to ..."},
			Validate: survey.Required,
		},
		{
			Name:     "syntax",
			Prompt:   &survey.Input{Message: "In"},
			Validate: survey.Required,
		},
		{
			Name:   "source",
			Prompt: &survey.Input{Message: "From"},
		},
		{
			Name:   "snippet",
			Prompt: &survey.Editor{Message: "Content"},
		},
	}

	answers := struct {
		Title   string
		Syntax  string
		Source  string
		Snippet string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return api.CreateSnippet(answers.Title, answers.Syntax, answers.Source, answers.Snippet)
}
