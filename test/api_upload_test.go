package test

import (
	"fmt"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	if os.Getenv("APP_ENV") == "runner" {
		return
	}
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	c.IsDebug()

	filePath := "test.mp4"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

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
		"title.mp4",
		"/VMFS1/FILES/upload/test",
	)

	fmt.Print(response)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

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
