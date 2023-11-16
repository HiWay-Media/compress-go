package test

import (
	"log"
	"testing"
)

/*
*/
func TestGetRestreamers(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	restreamers, err := c.GetRestreamers()
	if err != nil {
		t.Fatalf(err.Error())
	}
	//log.Println("restreamers ", restreamers)
	//c.IsDebug()
	//
	_, err := c.GetSingleRestreamer(restreamers[0].Name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//
}


