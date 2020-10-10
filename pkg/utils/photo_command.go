package utils

import (
	"agileengine/imagecache/pkg/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var isImagedInit = false
var client = &http.Client{}

func LoadImages() {
	if isImagedInit {
		return
	}
	searchImages()
}

func searchImages() {
	token := getAuthorizationToken()
	fmt.Println(token)
}

func getAuthorizationToken() string {
	req := createAuthorizationRequest()
	res, err := client.Do(req)
	CheckError(err, "We have an error trying to load the photos from the server")
	CheckError(CheckStatusCodeIs200(res), "Error in the response, it must to be a 200")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	CheckError(err, "Error trying to read the response from the server")
	response := new(dto.AuthResponse)
	CheckError(json.Unmarshal(body, &response), "Error trying to deserialize the json from the response")
	return response.Token
}

func createAuthorizationRequest() *http.Request {
	reqTokenBody := dto.AuthRequest{
		ApiKey: GetConfigValueFromKey(ApiKey),
	}
	bytesRequest, err := json.Marshal(reqTokenBody)
	CheckError(err, "Error trying to marshal the body for get the authentication body")
	req, err := http.NewRequest("POST", "http://interview.agileengine.com/auth", bytes.NewBuffer(bytesRequest))
	CheckError(err, "Error trying to create the request")
	req.Header.Set("Content-type", "application/json")
	return req
}
