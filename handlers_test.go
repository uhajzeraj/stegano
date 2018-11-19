package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_rootHandler_notImplemented(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 200 {
		t.Errorf("Expected StatusOK %d, received %d. ", 200, resp.StatusCode)
		return
	}
}

func Test_rootHandler_malformedUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/project/hi/etc",
		ts.URL + "/proj/",
	}
	for _, tstring := range testCases {
		resp, err := http.Get(tstring)
		if err != nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("For route: %s, expected StatusCode %d, received %d", tstring,
				http.StatusOK, resp.StatusCode)
			return
		}
	}

}

// var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func Test_homeHandler_notImplemented(t *testing.T) {

	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(homeHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}
	session, _ := store.Get(req, "user-login")
	session.Values["user"] = "uhajzeraj"

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected SatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_homeHandler_malformedUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(homeHandler))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/home",
		ts.URL + "/hom",
	}
	for _, tstring := range testCases {
		resp, err := http.Get(tstring)
		if err == nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != http.StatusSeeOther {
			t.Errorf("For route: %s, expected StatusCode %d, received %d", tstring,
				http.StatusSeeOther, resp.StatusCode)
			return
		}
	}

}

func Test_userHandler_notImplemented(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(userHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected SatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_userHandler_MalformedURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(userHandler))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/something",
		ts.URL + "/something/123",
	}

	for _, tstring := range testCases {
		resp, err := http.Get(ts.URL)
		if err == nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != http.StatusSeeOther {
			t.Errorf("For route: %s, expected StatusSeeOther %d, received %d. ", tstring, http.StatusSeeOther, resp.StatusCode)
			return
		}
	}
}

func Test_savedHandler_notImplemeted(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(savedHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOthers  %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_logoutHandler_notImplemented(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(logoutHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOthers %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_loginGetHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(loginGetHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 200 {
		t.Errorf("Expected Status OK %d, received %d. ", 200, resp.StatusCode)
		return
	}
}

func Test_loginPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(loginPostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 200 {
		t.Errorf("Expected Status OKS %d, received %d. ", 200, resp.StatusCode)
		return
	}
}

func Test_signupGetHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(signupGetHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 200 {
		t.Errorf("Expected Status OK %d, received %d. ", 200, resp.StatusCode)
		return
	}
}
func Test_signupPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(signupPostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 200 {
		t.Errorf("Expected StatusFound %d, received %d. ", 200, resp.StatusCode)
		return
	}
}
func Test_steganoGetHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(steganoGetHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_steganoPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(steganoPostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_steganoDecodeHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(steganoDecodeHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}
func Test_caesarGetHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(caesarGetHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}
func Test_caesarPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(caesarPostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}
func Test_caesarDecodeHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(caesarDecodeHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_rot13GetHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rot13GetHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}
func Test_rot13PostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rot13PostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}
func Test_deleteImgPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(deleteImgPostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expacected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_changePassPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(changePassPostHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 303 {
		t.Errorf("Expacected StatusSeeOther %d, received %d. ", 303, resp.StatusCode)
		return
	}
}

func Test_deleteAccDeletePostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(deleteAccDeleteHandler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the DELETE request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusSeeOther {
		t.Errorf("Expected StatusSeeOther %d, received %d. ", http.StatusSeeOther, resp.StatusCode)
		return
	}

}
