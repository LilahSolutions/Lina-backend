package requestHelper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"middlend-template/pkg/domain/innerDomain/response"
	"net/http"
	"os"
)

func GetRequest(url string) (*http.Response, response.Status) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("lila", os.Getenv("BACK_PASS")))

	result, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return result, response.SuccessfulSearch
}

func GetRequestBearerToken(url, bearerToken string) (*http.Response, response.Status) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, response.InternalServerError
	}
	// Set authorization header
	bearer := "Bearer " + bearerToken
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	result, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return result, response.SuccessfulSearch
}

// Use authorization without prefix
func GetRequestAccessToken(url, accessToken string) (*http.Response, response.Status) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, response.InternalServerError
	}
	// Set authorization header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", accessToken)

	result, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return result, response.SuccessfulSearch
}

func PostRequest(url string, body interface{}) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.InternalServerError
	}
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(parsedObject))
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("lila", os.Getenv("BACK_PASS")))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.StatusCreated
}

func PostRequestBearerToken(url string, body interface{}, bearerToken string) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.InternalServerError
	}
	fmt.Println(string(parsedObject))

	// Set authorization header
	bearer := "Bearer " + bearerToken

	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(parsedObject))
	req.Header.Add("Authorization", bearer)
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.StatusCreated
}

func PostRequestAccessToken(url string, body interface{}, accessToken string) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.InternalServerError
	}
	fmt.Println(string(parsedObject))

	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(parsedObject))
	req.Header.Add("Authorization", accessToken)
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.StatusCreated
}

// Body must be a marshalled object
func PutRequest(url string, body interface{}) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.InternalServerError
	}
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(parsedObject))
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("lila", os.Getenv("BACK_PASS")))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.SuccessfulUpdate
}

func PutRequestAccessToken(url string, body interface{}, accessToken string) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.InternalServerError
	}
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(parsedObject))
	req.Header.Add("Authorization", accessToken)
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.SuccessfulUpdate
}

func DeleteRequest(url string) (*http.Response, response.Status) {
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("lila", os.Getenv("BACK_PASS")))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.SuccessfulDeletion
}

func DeleteRequestWithAccessToken(url, accessToken string) (*http.Response, response.Status) {
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, response.InternalServerError
	}
	req.Header.Add("Authorization", accessToken)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.InternalServerError
	}
	return res, response.SuccessfulDeletion
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
