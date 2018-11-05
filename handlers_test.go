package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_rootHandler(t *testing.T){


	testServer := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer testServer.Close()

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("StatusNotFound %d, received %d. ",200, response.StatusCode)
		return
	}




	}
	func Test_steganoPostHandler(t *testing.T) {

		cl, err := mongoConnect()
		if cl == nil  {
			t.Error("No connection")
		}
		if err!=nil{
			t.Errorf("Error, %s",err)
		}

		img,err :=getImage(cl)
		if img==nil{
			t.Error("No image")
		}
		if err!=nil{
			t.Errorf("Error, %s",err)
		}

}