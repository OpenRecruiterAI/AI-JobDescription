package aijobs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func JobSetup(config Config) *Job {

	return &Job{
		DB: config.DB,
	}
}

func GenerateJobDescriptionFromTitle(jobTitle, apiEndpoint string) (responseString string, err error) {
	if jobTitle == "" {
		return "", errors.New("error generating job descripton: jobtitle is empty")
	}

	if apiEndpoint == "" {
		return "", errors.New("error generating job descripton: api endpoint is empty")
	}

	resp, err := http.Post(apiEndpoint, "application/x-www-form-urlencoded", &bytes.Buffer{})
	if err != nil {
		return "", fmt.Errorf("error generating job descripton: %v", err)
	}

	if resp == nil {
		return "", errors.New("error generating job descripton: failed to get response from the api")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return " ", fmt.Errorf("err while reading the response body: %v", err)
	}

	return string(respBody), nil

}
