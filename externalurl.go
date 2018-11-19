package main

// func handler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		t, err := template.ParseFiles("web-dev/front-end/stegano.html")
// 		if err != nil {
// 			http.Error(w, "Error processing template", http.StatusInternalServerError)
// 		}
// 		t.Execute(w, nil)

// 	case http.MethodPost:
// 		url := "http://i.imgur.com/m1UIjW1.png"

// 		response, e := http.Get(url)
// 		if e != nil {
// 			log.Fatal(e)
// 		}

// 		defer response.Body.Close()

// 		//open a file for writing
// 		file, err := os.Create("./tmp/asdf.png")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// Use io.Copy to just dump the response body to the file. This supports huge files
// 		_, err = io.Copy(file, response.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		file.Close()
// 		//fmt.Println("Success!")
// 		// Read image from file that already exists
// 		existingImageFile, err := os.Open("./tmp/asdf.png")
// 		if err != nil {
// 			// Handle error
// 			log.Fatal(err)
// 		}
// 		defer existingImageFile.Close()

// 		// Calling the generic image.Decode() will tell give us the data
// 		// and type of image it is as a string. We expect "png"
// 		imageData, imageType, err := image.Decode(existingImageFile)
// 		if err != nil {
// 			// Handle error
// 			log.Fatal(err)
// 		}
// 		fmt.Println(imageData)
// 		fmt.Println(imageType)

// 		// We only need this because we already read from the file
// 		// We have to reset the file pointer back to beginning
// 		existingImageFile.Seek(0, 0)

// 		// Alternatively, since we know it is a png already
// 		// we can call png.Decode() directly
// 		loadedImage, err := png.Decode(existingImageFile)
// 		if err != nil {
// 			// Handle error
// 			log.Fatal(err)
// 		}
// 		fmt.Println(loadedImage)
// 	default:
// 		http.Error(w, "Not implemented", http.StatusNotImplemented)
// 	}

// }
