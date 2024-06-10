package test


import (
	"log"
	"testing"
)

func TestGetCustomerS3(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// need to be finish
	zone, err := c.GetCustomerS3Zone()
	if err != nil {
		t.Fatalf(err.Error())
	}
	log.Println("zone ", zone)
	//c.IsDebug()
}
