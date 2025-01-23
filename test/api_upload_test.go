package test

import (
	"fmt"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	/*if os.Getenv("APP_ENV") == "runner" {
		return
	}*/
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	//c.IsDebug()
	filePath := "test.mp4"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("file does not exist: %s", filePath)
	}
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//fmt.Print(fileContent)
	stat, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	response, err := c.Upload(
		fileContent,
		stat.Size(),
		60678,
		"title",
		"tag",
		"",
		filePath,
	)
	//fmt.Print(response)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if response.Response != "OK" {
		t.Fatalf(fmt.Sprintf("error %s", response.Message))
	}
}

/*
func TestUploadMultipart(t *testing.T) {
	if os.Getenv("APP_ENV") == "runner" {
		return
	}
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	c.IsDebug()

	filePath := "test.mp4"
	if err != nil {
		t.Fatalf(err.Error())
	}

	stat, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = c.UploadMultipart(
		file,
		stat.Size(),
		60678,
		"title",
		"tag",
		"",
		"titlemultipart.mp4",
		"/VMFS1/FILES/upload/test",
	)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
*/
