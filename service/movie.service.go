package service

import (
	"context"
	"fmt"
	"log"

	"github.com/shaikdev/GO-MongoDB/db"
	"github.com/shaikdev/GO-MongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMovie(movie models.Movie) string {
	response, err := db.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	insertedID := response.InsertedID.(string)
	fmt.Println("The movie was created", response.InsertedID)
	return insertedID
}

func DeleteMovieById(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	response, err := db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The movie was deleted", response.DeletedCount)
	return response.DeletedCount
}

func DeleteAllMovies() bool {
	response, err := db.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The all movie was deleted", response.DeletedCount)
	return true
}
