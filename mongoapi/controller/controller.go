package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

const dbName = "organization"
const colName = "developer"

var collection *mongo.Collection

func init() {

	// Create a client and connect to the server
	clientOption := options.Client().ApplyURI(connectionString)
	fmt.Println("Client option is", clientOption)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO() , clientOption)
	fmt.Println("Client is", client)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is", collection)

	fmt.Println("Collection instance is ready!")

}

func insertDeveloper(developer models.Developer) string{

	inserted, err := collection.InsertOne(context.Background(),developer)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Inserted One Developer with Id :", inserted.InsertedID)

	return "Developer Inserted"
}

func updateDeveloperAsLead(developerId string) string {

	id, err := primitive.ObjectIDFromHex(developerId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"lead": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count :", result.ModifiedCount)

	return "Developer Details Updated" 
}

func deleteDeveloper(developerId string) string{
	
	id, err := primitive.ObjectIDFromHex(developerId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id":id}

	collection.DeleteOne(context.Background(), filter)

	return "Developer Deleted"
}

func deleteDevelopers() int64 {

	filter := bson.D{{}}
	deletedCount, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}
	
	return deletedCount.DeletedCount
}

func getDevelopers() []primitive.M {
	
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var developers []primitive.M

	for cursor.Next(context.Background()) {
		var developer bson.M
		err := cursor.Decode(&developer)

		if err != nil {
			log.Fatal(err)
		}

		developers = append(developers, developer)
	}

	defer cursor.Close(context.Background())

	return developers
}

func GetDevelopers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	
	developers := getDevelopers()
	fmt.Println("Developers from database are", developers)

	json.NewEncoder(w).Encode(developers)
}

func CreateDeveloper(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var developer models.Developer
	json.NewDecoder(r.Body).Decode(&developer)

	insertDeveloper(developer)

	json.NewEncoder(w).Encode(developer)
}

func UpdateDeveloperAsLead(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateDeveloperAsLead(params["id"])

	json.NewEncoder(w).Encode("Developer Updated")
}

func DeleteDeveloper(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteDeveloper(params["id"])

	json.NewEncoder(w).Encode("Developer Deleted")
}


func DeleteAllDeveloper(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteDevelopers()
	
	json.NewEncoder(w).Encode(count)
}