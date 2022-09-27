package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type prompt string

func (p prompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"model":       "text-davinci-002",
		"temperature": 0,
		"prompt": `Convert this text to a programmatic command:

` + string(p),
	})
}

type answer struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	token := os.Getenv("OPENAI_API_TOKEN")
	if token == "" {
		panic(`env variable "OPENAI_API_TOKEN" is required`)
	}

	body, err := json.Marshal(prompt(strings.Join(os.Args[1:], " ")))
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	var result answer
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		panic(err)
	}

	if err := response.Body.Close(); err != nil {
		panic(err)
	}

	if len(result.Choices) > 0 {
		fmt.Println(strings.TrimSpace(result.Choices[0].Text))
		return
	}

	panic("can not find choice in response")
}
