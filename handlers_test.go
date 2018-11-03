package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_rootHandler(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer server.Close()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET method, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}
	cl, err := mongoConnect()
	if cl == nil  {
		t.Error("No connection")
	}
	if err!=nil{
		t.Errorf("Error, %s",err)
	}
	/*if img,err := getImage(cl); img == nil || err!=nil {
		t.Error("No image")
	}*/
	img,err :=getImage(cl)
	if img==nil{
		t.Error("No image")
	}
	if err!=nil{
		t.Errorf("Error, %s",err)
	}


}


