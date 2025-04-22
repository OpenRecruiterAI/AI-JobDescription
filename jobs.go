package aijobs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func GenerateJobDescriptionFromFile() {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	filePart, err := writer.CreateFormFile("file", jdFileHeader.Filename)
	if err != nil {
		fmt.Println("err getting file : ", err)
		c.JSON(200, gin.H{"err": "err creating request to the node api", "success": 0, "response": ""})
	}

	_, err = jdFile.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error resetting file pointer:", err)
		c.JSON(200, gin.H{"err": "Error resetting file pointer", "success": 0, "response": ""})
	}

	Js, err := io.Copy(filePart, jdFile)
	fmt.Println("Js,", Js)
	if err != nil {
		fmt.Println("couldn't write file data:", err)
		c.JSON(200, gin.H{"err": "couldn't write file data", "success": 0, "response": ""})
	}

	writer.Close()

	baseurl := os.Getenv("CHAT_API_URL") + "api/recruit/analyse?type=jobDetail"

	resp, err := http.Post(baseurl, writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println("failed to get data from api", err)
		c.JSON(200, gin.H{"err": "failed to get data from api", "success": 0, "response": ""})
	}

	if resp.StatusCode == 200 {
		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("failed to read the response from the body", err)
			c.JSON(200, gin.H{"err": "failed to read the response from the body", "success": 0, "response": ""})
		}

		c.JSON(200, gin.H{"err": "", "success": 1, "response": string(bodyResp)})

	} else {
		c.JSON(200, gin.H{"err": "failed to get response from the api", "success": 0, "response": ""})
	}
}
