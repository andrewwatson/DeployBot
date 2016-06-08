package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func performDeployment(version string, environment string, responseURL string) ([]byte, error) {

	log.Printf("Asked to respond to %s later", responseURL)
	wakeUp := time.After(5 * time.Second)
	<-wakeUp

	log.Print("i'm awake!")
	data := struct {
		Text           string
		AttachmentText string
	}{"test", "attachment text"}
	tmpl, err := template.New("json").Parse(responseTemplate)
	if err != nil {
		return nil, err
	}

	var jsonOutput bytes.Buffer
	err = tmpl.Execute(&jsonOutput, data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", responseURL, bytes.NewBuffer(jsonOutput))

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
