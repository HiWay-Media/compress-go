package test

import (
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
	restreamers, err := c.GetRestreamers(0, 10)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//log.Println("restreamers ", restreamers)
	//c.IsDebug()
	//
	_, err = c.GetSingleRestreamer(restreamers[0].Name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//
}

/*
*
 */
func TestGetRestreamersOttAll(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	_, err = c.GetRestreamersOttAll(0, 100)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//log.Println("restreamers ott ", restreamers)
	//
}
