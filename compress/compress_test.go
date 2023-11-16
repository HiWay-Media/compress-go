package compress

import (
	"os"
	"testing"
	"fmt"
)

func TestNewCompress(t *testing.T) {
	//
	c, err := GetCompress()
	if err != nil {
		t.Fatalf(err.Error())
	}
	c.IsDebug()
}


func GetCompress() ( ICompress, error ) {
	//
	apiKey 			:= os.Getenv("API_KEY")
	customerName	:= os.Getenv("CUSTOMER_NAME")
	//
	c, err := &NewCompress(customerName, apiKey, false)
	if err != nil {
		return nil, err
	}
	//
	return c, nil
}