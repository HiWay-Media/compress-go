package test

import (
	"log"
	"testing"
)

func TestGetRestreamers(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	restreamers, err := c.GetRestreamers(1, 20)
	if err != nil {
		t.Fatalf(err.Error())
	}
	log.Println("restreamers ", restreamers)
	//c.IsDebug()
}
