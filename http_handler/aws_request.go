package http_handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"src/local/oscli/configure"
	"src/local/oscli/structs"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

const shortDuration = 100 * time.Millisecond

func GetRequest(url string) *http.Response {
	var config structs.ConfigFile
	config = configure.GetCredentials() //gets ~/.oscli/config.json info {es host}
	client := &http.Client{}
	reqUrl := config.Host + "/" + url
	method := "GET"
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	// signs the AWS request with V4 Signing Algorithm, returns *http.header
	_ = signRequest(req, nil)
	resp, clientErr := client.Do(req)
	if clientErr != nil {
		fmt.Print(clientErr.Error())
	}
	return resp
}

func DeleteRequest(url string) *http.Response {
	var config structs.ConfigFile
	config = configure.GetCredentials() //gets ~/.oscli/config.json info {es host}
	client := &http.Client{}
	reqUrl := config.Host + "/" + url
	method := "DELETE"
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	// signs the AWS request with V4 Signing Algorithm, returns *http.header
	_ = signRequest(req, nil)
	resp, clientErr := client.Do(req)
	if clientErr != nil {
		fmt.Print(clientErr.Error())
	}
	return resp
}

// payload map[string]interface{}
func PutRequest(url string, payload interface{}) *http.Response {
	requestBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestBody))
	body := strings.NewReader(string(requestBody))

	var config structs.ConfigFile
	config = configure.GetCredentials() //gets ~/.oscli/config.json info {es host}
	client := &http.Client{}
	reqUrl := config.Host + "/" + url
	method := "PUT"

	req, err := http.NewRequest(method, reqUrl, body)
	req.Header.Set("Content-type", "application/json")
	if err != nil {

		fmt.Println(err.Error())
	}
	// signs the AWS request with V4 Signing Algorithm, returns *http.header
	_ = signRequest(req, bytes.NewReader(requestBody))
	resp, clientErr := client.Do(req)
	if clientErr != nil {

		fmt.Print(clientErr.Error())
	}

	return resp
}

/*
 signs the request with the v4 aws signing method using the client key and secret and service
 this also allows for IAM roles of the user running the oscli to be leveraged for opensearch cluster permissions
*/

func signRequest(req *http.Request, body io.ReadSeeker) *http.Header {
	now := time.Now()
	var config structs.ConfigFile
	profile := external.WithSharedConfigProfile(config.Profile)
	cfg, extErr := external.LoadDefaultAWSConfig(profile)
	if extErr != nil {
		fmt.Println(extErr.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("..api response")
	}

	signer := v4.NewSigner(cfg.Credentials, func(signer *v4.Signer) {
		signer.Debug = aws.LogDebugWithSigning
	})

	signedHeader, e := signer.Sign(ctx, req, body, "es", cfg.Region, now)
	if e != nil {
		fmt.Printf("failed to sign request: (%v)\n", e.Error())
	}
	return &signedHeader
}
