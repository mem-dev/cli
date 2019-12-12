package cmd

import (
	"errors"
	"fmt"
	"github.com/mem-dev/cli/api"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Login() bool {
	questions := []*survey.Question{
		{
			Name:   "authToken",
			Prompt: &survey.Input{Message: "Please go to https://codecode.ninja/auth_token and enter your secret string:"},
			Validate: func(val interface{}) error {
				if val == nil {
					return errors.New("No secret string entered")
				}
				return nil
			},
		},
	}

	answers := struct {
		AuthToken string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return api.Authenticate(answers.AuthToken)
}
