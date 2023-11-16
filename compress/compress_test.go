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
		t.Falf(err.Error())
	}
	c.IsDebug()
}