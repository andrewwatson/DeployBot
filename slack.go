package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"
)

const (
	responseTemplate = `{
	"response_type": "{{.Type}}",
	"text": "{{.Text}}",
	"attachments": [
		{
				"text":"{{.AttachmentText}}"
		}
	]
}`
)

func respondToSlack(responseURL, title, message string, ephemeral bool) ([]byte, error) {

	rType := "in_channel"
	if ephemeral {
		rType = "ephemeral"
	}

	data := struct {
		Text           string
		AttachmentText string
		Type           string
	}{title, message, rType}
	tmpl, err := template.New("json").Parse(responseTemplate)
	if err != nil {
		return nil, err
	}

	var jsonOutput bytes.Buffer
	err = tmpl.Execute(&jsonOutput, data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", responseURL, bytes.NewBuffer(jsonOutput.Bytes()))

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
