package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func respondToSlack(responseURL string, message string) ([]byte, error) {

	data := struct {
		Text           string
		AttachmentText string
	}{"Deployment Result", message}
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
func performDeployment(version string, environment string, responseURL string) ([]byte, error) {

	log.Printf("Asked to respond to %s later", responseURL)
	wakeUp := time.After(20 * time.Second)
	<-wakeUp

	responseBody, err := respondToSlack(responseURL, fmt.Sprintf("Deployment of %s to %s Successful", version, environment))

	return responseBody, err

}
