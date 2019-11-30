package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// sqlite support
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type user struct {
	gorm.Model
	Name  string
	Email string
}

var dbDialect = "sqlite3"
var dbName = "test.db"

func initialMigration() {
	db, err := gorm.Open(dbDialect, dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&user{})
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbDialect, dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	var users []user
	db.Find(&users)
	fmt.Printf("Users: %v", users)

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbDialect, dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&user{Name: name, Email: email})
	fmt.Printf("Creating user: %v, %v\n", name, email)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbDialect, dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user user
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)

	fmt.Printf("Updating user: %v\n", user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbDialect, dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	name := vars["name"]

	var user user
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Printf("Deleting user: %v\n", user)
}

// RunORM https://tutorialedge.net/golang/golang-orm-tutorial/
//
// Import needed:
// _ "github.com/jinzhu/gorm/dialects/sqlite"
// which requires:
// go get github.com/mattn/go-sqlite3
func RunORM() {
	// prepare
	initialMigration()

	// go!
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", getAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", createUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
