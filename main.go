package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type LambdaRequest struct {
	RunID string `json:"run_id"`
}

type LambdaResponse struct {
	statusCode int    `json:"statusCode"`
	body       string `json:"body"`
}

func lambdaHandler(event LambdaRequest) (LambdaResponse, error) {
	rid := event.RunID
	if rid == "" {
		return LambdaResponse{
			statusCode: 200,
			body:       "OK",
		}, nil
	}
	log.Printf("handling %s", rid)

	bearer := "Bearer" + os.Getenv("TFE_TOKEN")
	url := fmt.Sprintf("https://app.terraform.io/api/v2/runs/%s/actions/apply", rid)
	ct := "application/vnd.api+json"

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", ct)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string([]byte(body)))

	return LambdaResponse{
		statusCode: 200,
		body:       string([]byte(body)),
	}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
