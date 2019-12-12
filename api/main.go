package api

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/spin"
	"github.com/mem-dev/cli/auth"
	"gopkg.in/resty.v1"
)

func getClient() *resty.Client {
	apiClient := resty.
		New().
		SetHostURL("http://localhost:3000").
		SetHeader("Content-Type", "application/json")

	if auth.IsAuthenticated() == true {
		authData := auth.Auth{}
		apiClient.
			SetHeader("Authorization", fmt.Sprintf(`Bearer %s`, authData.JwtToken))

	}
	return apiClient
}

func CreateSnippet(title string, syntax string, source string, snippet string) bool {
	s := spin.New("%s Creating snippet...")
	s.Start()
	defer s.Stop()
	client := getClient()
	payload := fmt.Sprintf(`
	{
		"snippet": {
			"title": "%s",
			"syntax": "%s",
			"source": "%s",
			"content": "dsdghjksdlhss",
			"topic": "a topic"
		}
	}
	`, title, syntax, source)

	resp, err := client.
		R().
		SetBody(payload).
		Post("/api/v2/snippets")

	if err != nil {
		return false
	}

	if resp.StatusCode() != 200 {
		fmt.Printf(resp.Status())
		return false
	}

	fmt.Printf("Your snippet was created, you can convert it to a card from here: https://codecode.ninja/snippets/process")
	return true
}

func Authenticate(authToken string) bool {
	s := spin.New("%s Authenticating...")
	s.Start()
	defer s.Stop()
	client := getClient()
	url := fmt.Sprintf(`/api/v2/authorize/%s`, authToken)

	resp, err := client.R().Get(url)

	if err != nil {
		s.Stop()
		return false
	}

	if resp.StatusCode() != 200 {
		s.Stop()
		fmt.Println(resp.Status())
		return false
	}

	a := auth.Auth{}

	json.Unmarshal(resp.Body(), &a)
	fmt.Println(a)

	a.Persist()

	return true
}
