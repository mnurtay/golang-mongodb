package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateFunc ...
func CreateFunc() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "application/json")
		var person Person
		json.NewDecoder(request.Body).Decode(&person)
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, _ := mongo.Connect(context.TODO(), clientOptions)
		collection := client.Database("golang-mongodb").Collection("people")
		result, err := collection.InsertOne(context.TODO(), person)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ message: "` + err.Error() + `" }`))
			return
		}
		client.Disconnect(context.TODO())
		json.NewEncoder(response).Encode(result)
	}
}
