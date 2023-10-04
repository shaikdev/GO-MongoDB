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
		return primitive.ObjectID{}, err
	}
	insertedID := response.InsertedID.(primitive.ObjectID)
	fmt.Println("The movie was created", response.InsertedID)
	return insertedID, nil
}

func DeleteMovieById(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	response, _ := db.Collection.DeleteOne(context.Background(), filter)
	return response.DeletedCount
}

func DeleteAllMovies() int64 {
	response, err := db.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		return 0
	}
	fmt.Println("The all movie was deleted", response.DeletedCount)
	return response.DeletedCount
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

func GetMovieById(id string) (models.Movie, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	movie := models.Movie{}
	err := db.Collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func UpdateMovie(id string, body models.Movie) int64 {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{}
	if len(body.Movie) != 0 {
		update["movie"] = body.Movie
	}
	update["watched"] = body.Watched
	setUpdatedBody := bson.M{"$set": update}
	response, _ := db.Collection.UpdateOne(context.Background(), filter, setUpdatedBody)
	return response.ModifiedCount
}
