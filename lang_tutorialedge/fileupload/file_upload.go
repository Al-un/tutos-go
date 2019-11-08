package fileupload

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Reject GET methods
	if r.Method == "GET" {
		fmt.Printf("Web: refusing GET in %s\n", r.URL.Path)
		http.ServeFile(w, r, "fileupload/error.html")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "Client: Start file uploading...\n")

	// max size
	r.ParseMultipartForm(10 << 20)

	// Get file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Server Error: retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Server: Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File Server: size: %+v\n", handler.Size)
	fmt.Printf("Server: MIME Header: %+v\n", handler.Header)

	// Create a temporary file
	tempFile, err := ioutil.TempFile("fileupload/temp-images", "upload-*.png")
	if err != nil {
		fmt.Println("Server Error: ", err)
	}
	defer tempFile.Close()

	// Read content
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// Write!
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Client: Successfully Uploaded File \\o/\n")
	fmt.Fprintf(w, "<a href=\"/\">Go back<a/>\n")
}

func setupRoutes() {
	// index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Web: Requesting content %s\n", r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// uploader
	http.HandleFunc("/upload", uploadFile)

	// Go!
	http.ListenAndServe(":8080", nil)
}

// RunFileUpload https://tutorialedge.net/golang/go-file-upload-tutorial/
func RunFileUpload() {
	fmt.Println("Web: Initalizing File uploader...")
	setupRoutes()
}
