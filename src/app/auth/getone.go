package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetOneFunc ...
func GetOneFunc() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "application/json")
		params := mux.Vars(request)
		id, _ := primitive.ObjectIDFromHex(params["id"])
		var person Person
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, _ := mongo.Connect(context.TODO(), clientOptions)
		collection := client.Database("golang-mongodb").Collection("people")
		err := collection.FindOne(context.TODO(), Person{ID: id}).Decode(&person)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ message: "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(person)
	}
}
