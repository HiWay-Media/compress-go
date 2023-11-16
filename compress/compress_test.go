package compress

import (
	"os"
	"testing"
)

func TestNewCompress(t *testing.T) {
	//
	apiKey 			:= os.Getenv("API_KEY")
	customerName	:= os.Getenv("CUSTOMER_NAME")
	//
	c, err := NewCompress(customerName, apiKey, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	c.IsDebug()
}