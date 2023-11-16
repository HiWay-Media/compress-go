package compress

import (
	"testing"
)

func TestGetCategories(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	//c.GetCategories()
	//c.IsDebug()
}