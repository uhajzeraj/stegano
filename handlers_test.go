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

		ts := httptest.NewServer(http.HandlerFunc(steganoPostHandler))
		defer ts.Close()
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
		client := &http.Client{}



		req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
		if err != nil {
			t.Errorf("Error constructing the POST request, %s", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Error executing the POST request, %s", err)
		}

		//check if the response from the handler is what we except
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusOK, resp.StatusCode)
			return
		}

}
func Test_steganoGetHandler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(steganoGetHandler))
	defer ts.Close()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the Get request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the Get request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}



}

func Test_loginHandler(t *testing.T){

	ts := httptest.NewServer(http.HandlerFunc(loginHandler))
	defer ts.Close()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the Get request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the Get request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}



}

