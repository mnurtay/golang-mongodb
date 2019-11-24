package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllFunc ...
func GetAllFunc() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "application/json")
		var people []Person
		json.NewDecoder(request.Body).Decode(&people)
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.TODO(), clientOptions)
		collection := client.Database("golang-mongodb").Collection("people")
		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ message: "` + err.Error() + `" }`))
			return
		}
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var person Person
			cursor.Decode(&person)
			people = append(people, person)
		}
		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ message: "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(people)
	}
}
