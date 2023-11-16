package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("APP_ENV") == "" {
		err := os.Setenv("APP_ENV", "test")
		if err != nil {
			panic("could not set test env")
		}
	}
	//env.Load()
	m.Run()
}


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
	c, err := NewCompress(customerName, apiKey, false)
	if err != nil {
		return nil, err
	}
	//
	return c, nil
}