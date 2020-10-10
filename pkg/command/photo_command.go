package command

import (
	"agileengine/imagecache/pkg/dto"
	"agileengine/imagecache/pkg/model"
	"agileengine/imagecache/pkg/repository"
	"agileengine/imagecache/pkg/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	page := searchPage(1)
	getImageFromPage(page)
	for i := 2; i <= page.PageCount; i++ {
		page := searchPage(1)
		getImageFromPage(page)
	}
}

func getImageFromPage(page *dto.GetPicturesResponse) []interface{} {
	var pictures []interface{}
	for _, picture := range page.Pictures {
		pictures = append(pictures, getParticularImage(picture))
	}
	repository.StoreImages(pictures)
	return pictures
}

func getParticularImage(picture dto.Picture) *model.Image {
	req := createRequestForImage(picture.Id)
	res, err := client.Do(req)
	utils.CheckError(err, "Error tyring to get the image with id: "+picture.Id)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	utils.CheckError(err, "Error trying to read the response from the server when tries to get the images")
	response := new(model.Image)
	utils.CheckError(json.Unmarshal(body, &response), "Error trying to deserialize the json from the response")
	return response
}

func searchPage(page int) *dto.GetPicturesResponse {
	req := createRequestForPage(page)
	res, err := client.Do(req)
	utils.CheckError(err, "Error trying to get the page image from the server")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	utils.CheckError(err, "Error trying to read the response from the server when tries to get the images")
	response := new(dto.GetPicturesResponse)
	utils.CheckError(json.Unmarshal(body, &response), "Error trying to deserialize the json from the response")
	return response
}

func getAuthorizationToken() string {
	token, found := utils.Cache.Get("token")
	if !found {
		req := createAuthorizationRequest()
		res, err := client.Do(req)
		utils.CheckError(err, "We have an error trying to load the photos from the server")
		utils.CheckError(utils.CheckStatusCodeIs200(res), "Error in the response, it must to be a 200")
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		utils.CheckError(err, "Error trying to read the response from the server")
		response := new(dto.AuthResponse)
		utils.CheckError(json.Unmarshal(body, &response), "Error trying to deserialize the json from the response")
		token = response.Token
		utils.Cache.Set("token", token, 5*time.Minute)
	}
	return token.(string)

}

func createAuthorizationRequest() *http.Request {
	reqTokenBody := dto.AuthRequest{
		ApiKey: utils.GetConfigValueFromKey(utils.ApiKey),
	}
	bytesRequest, err := json.Marshal(reqTokenBody)
	utils.CheckError(err, "Error trying to marshal the body for get the authentication body")
	req, err := http.NewRequest("POST", "http://interview.agileengine.com/auth", bytes.NewBuffer(bytesRequest))
	utils.CheckError(err, "Error trying to create the request")
	req.Header.Set("Content-type", "application/json")
	return req
}

func createRequestForPage(page int) *http.Request {
	req, err := http.NewRequest("GET", "http://interview.agileengine.com:80/images?page="+strconv.Itoa(page), nil)
	utils.CheckError(err, "Error trying to create the request")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getAuthorizationToken())
	return req
}

func createRequestForImage(id string) *http.Request {
	req, err := http.NewRequest("GET", "http://interview.agileengine.com:80/images/"+id, nil)
	utils.CheckError(err, "Error trying to create the request")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getAuthorizationToken())
	return req
}
