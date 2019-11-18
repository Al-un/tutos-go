package intermediate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

var articles []article
var articleNextID int

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
	fmt.Printf("home page done\n")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET Articles\n")
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	var id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("GetArticle: ", err)
	}

	fmt.Printf("GET Article ID %d\n", id)

	for _, a := range articles {
		if a.ID == id {
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	w.WriteHeader(404)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var art article
	json.Unmarshal(reqBody, &art)
	art.ID = articleNextID
	articles = append(articles, art)

	json.NewEncoder(w).Encode(art)
	articleNextID++

	fmt.Printf("Created: %v\n", art)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

	var id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("DeleteArticle: ", err)
	}

	fmt.Printf("GET Article ID %d\n", id)

	for index, a := range articles {
		if a.ID == id {
			w.WriteHeader(204)
			fmt.Printf("Delete %v\n", a)

			articles = append(articles[:index], articles[index+1:]...)
			return
		}
	}

	fmt.Printf("Not found ID %v\n", id)
	w.WriteHeader(404)
}

func setupRoutes() {
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", getArticles)

	// v2: Mux!
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func populate() {
	articles = []article{
		article{ID: 1, Title: "Article 1", Desc: "Description 1", Content: "Content 1"},
		article{ID: 2, Title: "Article 2", Desc: "Description 2", Content: "Content 2"},
	}

	articleNextID = 3
}

// RunRestAPI https://tutorialedge.net/golang/creating-restful-api-with-golang/
func RunRestAPI() {
	populate()
	setupRoutes()
}
