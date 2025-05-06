package aijobs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

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

func GenerateJobDescriptionFromFile(file multipart.File, fileName, apiEndpoint string) (responseString string, err error) {

	if apiEndpoint == "" {
		return "", errors.New("error generating job descripton: api endpoint is empty")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	filePart, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", fmt.Errorf("error creating file part: %v", err)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("error resetting file pointer: %v", err)
	}

	_, err = io.Copy(filePart, file)
	if err != nil {
		return "", fmt.Errorf("error writing file data: %v", err)
	}

	writer.Close()

	resp, err := http.Post(apiEndpoint, writer.FormDataContentType(), body)
	if err != nil {
		return "", fmt.Errorf("failed to get data from api: %v", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error : %v", err)
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the response from the body: %v", err)
	}

	return string(bodyResp), nil
}
