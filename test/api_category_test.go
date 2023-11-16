package test

import (
	"log"
	"testing"
)

func TestGetCategories(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	categories, err := c.GetCategories()
	if err != nil {
		t.Fatalf(err.Error())
	}
	log.Println("categories ", categories)
	//c.IsDebug()
}
