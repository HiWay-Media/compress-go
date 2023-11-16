package test

import (
	"testing"
)

func TestGetRestreamers(t *testing.T) {
	//
	_, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	//c.GetCategories()
	//c.IsDebug()
}
