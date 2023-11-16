package test

import (
	"fmt"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
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
