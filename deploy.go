package main

import (
	"fmt"
	"log"
	"time"
)

func performDeployment(version string, environment string, responseURL string) ([]byte, error) {

	log.Printf("Asked to respond to %s later", responseURL)
	wakeUp := time.After(20 * time.Second)
	<-wakeUp

	detailedResponse := fmt.Sprintf("Deployment of %s to %s Successful", version, environment)
	responseBody, err := respondToSlack(responseURL, "Deployment Update", detailedResponse)

	return responseBody, err

}
