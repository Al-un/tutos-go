package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/al-un/tutos-go/api_go-todos/server/models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// For localhost
const connectionString = "mongodb://localhost:27017"

// DB name
const dbName = "go_todos"

// Collection (table) name
const collName = "todolist"

// CORS
const allowedHosts = "http://localhost:3000"

// collection object/instance
var collection *mongo.Collection

// Init the connection with MongoDB upon app initialisation
func init() {
	// Client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Try to connect to DB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}

	log.Println("Connected to MongoDB \\o/")

	collection = client.Database(dbName).Collection(collName)
	log.Println("Collection instance created!")
}

// GetAllTask "get all task" route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)

	if r.Method == "GET" {
		payload := getAllTask()
		json.NewEncoder(w).Encode(payload)
	}
}

// CreateTask "create task" route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		var task models.TodoList
		// fmt.Println("To be created r.Body: ", r)
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("To be created task: ", task)

		insertOneTask(task)
		json.NewEncoder(w).Encode(task)
	}
}

// TaskComplete "update status" route
func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "PUT" {
		params := mux.Vars(r)
		taskComplete(params["id"])
		json.NewEncoder(w).Encode(params["id"])
	}
}

// UndoTask "cancel status update" route
func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "PUT" {
		params := mux.Vars(r)
		undoTask(params["id"])
		json.NewEncoder(w).Encode(params["id"])
	}
}

// DeleteTask "delete task" route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		deleteOneTask(params["id"])
		json.NewEncoder(w).Encode(params["id"])
	}
}

// DeleteAllTask is ARMEGEDDON!
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		count := deleteAllTask()
		json.NewEncoder(w).Encode(count)
	}
}

// DeleteDoneTask pouet
func DeleteDoneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", allowedHosts)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		count := deleteDoneTask()
		json.NewEncoder(w).Encode(count)
	}
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)

		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, " result", reflect.TypeOf(result), ref)
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Create task
func insertOneTask(task models.TodoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted a single record %v\n", insertResult.InsertedID)
}

// update status to true to mark it as completed
func taskComplete(task string) {
	fmt.Println("Completing task: ", task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Completing task: modified count: ", result.ModifiedCount)
}

// task undo: mark as un-completed
func undoTask(task string) {
	fmt.Println("Undoing task ", task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Undoing task: modified count: ", result.ModifiedCount)
}

// delete one task
func deleteOneTask(task string) {
	fmt.Println("Deleting task: ", task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted document: ", d.DeletedCount)
}

// Delete all task
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all tasks: ", d.DeletedCount)
	return d.DeletedCount
}

// Delete all task
func deleteDoneTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.M{"status": true}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted done tasks: ", d.DeletedCount)
	return d.DeletedCount
}
