package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

// Adapter : middleware style
type Adapter func(http.Handler) http.Handler

// Comment structure
type comment struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Author string        `json:"author" bson:"author"`
	Text   string        `json:"text" bson:"text"`
	When   time.Time     `json:"when" bson:"when"`
}

// Adapt pass through adapters
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func withDB(db *mgo.Session) Adapter {
	// return the adapter
	return func(h http.Handler) http.Handler {

		// Adapter, when called, return the new handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// copy the database session
			dbsession := db.Copy()
			defer dbsession.Close() // clean up

			// https://github.com/gorilla/context#context
			context.Set(r, "database", dbsession)

			h.ServeHTTP(w, r)
		})
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleRead(w, r)
	case "POST":
		handleInsert(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}

func handleInsert(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "database").(*mgo.Session)

	// decode request body
	var c comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// give the comment an unique ID and time
	c.ID = bson.NewObjectId()
	c.When = time.Now()

	// insert to DB!
	if err := db.DB("go_commentsapp").C("comments").Insert(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// redirect to it
	http.Redirect(w, r, "/comments/"+c.ID.Hex(), http.StatusTemporaryRedirect)
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "database").(*mgo.Session)

	// load comments
	var comments []*comment
	if err := db.DB("go_commentsapp").C("comments").Find(nil).Sort("-when").Limit(100).All(&comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send it
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	// connect to DB
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Cannot dial mongo", err)
	}
	defer db.Close() // clean up session

	// Adapt our handle function using DB
	h := Adapt(http.HandlerFunc(handle), withDB(db))

	// add the handler
	http.Handle("/comments", context.ClearHandler(h))

	// start server
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
