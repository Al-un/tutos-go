package intermediate

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

// var rootPath = "./lang_tutorialedge/intermediate/simple_web_server-static"
var rootPath = "./intermediate/simple_web_server-static"

func handleEchoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
}

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Requesting static file from URL %v\n", r.URL.Path)
	http.ServeFile(w, r, fmt.Sprintf("%s/%s", rootPath, r.URL.Path[1:]))
}

func handleIncrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

// RunSimpleWebServer https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
//
// must be run from the root `tutos-go/`
func RunSimpleWebServer() {
	// http.HandleFunc("/", handleEchoString)

	http.Handle("/static", http.FileServer(http.Dir(rootPath)))
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
	http.HandleFunc("/increment", handleIncrementCounter)
	http.HandleFunc("/", serveStaticFile)

	// log.Fatal(http.ListenAndServe(":8080", nil))

	// log.Fatal(http.ListenAndServeTLS(":8080",
	// 	"./lang_tutorialedge/server.crt",
	// 	"./lang_tutorialedge/server.key",
	// 	nil,
	// ),
	// 	nil,
	// )
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil),
		nil,
	)
}
