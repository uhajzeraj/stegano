package main

// func Test_externalUrlHandler_notImplemented(t *testing.T) {
// 	// instantiate mock HTTP server (just for the purpose of testing
// 	ts := httptest.NewServer(http.HandlerFunc(handler))
// 	defer ts.Close()

// 	//create a request to our mock HTTP server
// 	client := &http.Client{}

// 	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
// 	if err != nil {
// 		t.Errorf("Error constructing the Post request, %s", err)
// 	}

// 	resp, _ := client.Do(req)

// 	//check if the response from the handler is what we except
// 	if resp.StatusCode != 303 {
// 		t.Errorf("Expected SatusSeeOther %d, received %d. ", 303, resp.StatusCode)
// 		return
// 	}
// 	if resp.StatusCode != 200 {
// 		t.Errorf("Expected StatusOK %d, received %d. ", 200, resp.StatusCode)
// 		return
// 	}
// }
