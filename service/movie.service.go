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

func CreateMovie(movie models.Movie) (primitive.ObjectID, error) {
	response, err := db.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
		return primitive.ObjectID{}, err
	}
	insertedID := response.InsertedID.(primitive.ObjectID)
	fmt.Println("The movie was created", response.InsertedID)
	return insertedID, nil
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

func GetAllMovies() ([]primitive.M, int) {
	response, err := db.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for response.Next(context.Background()) {
		var movie bson.M

		decodeErr := response.Decode(&movie)
		if decodeErr != nil {
			log.Fatal(decodeErr)
		}

		movies = append(movies, movie)

	}

	defer response.Close(context.Background())
	return movies, len(movies)
}

func GetMovieById(id string) models.Movie {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	movie := models.Movie{}
	err := db.Collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	return movie
}

func UpdateMovie(id string, body map[string]interface{}) int64 {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"id": _id}
	update := bson.M{}
	for fieldName, fieldValue := range body {
		update[fieldName] = fieldValue
	}
	response, err := db.Collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M(update)})
	if err != nil {
		log.Fatal(err)
	}
	return response.ModifiedCount
}
