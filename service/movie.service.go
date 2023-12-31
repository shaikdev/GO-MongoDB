package service

import (
	"context"
	"fmt"
	"log"
	"reflect"

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

func UpdateMovie(id string, body models.Movie) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{}
	if len(body.Movie) > 0 {
		update["movie"] = body.Movie
	}
	if body.Watched || !body.Watched {
		update["watched"] = body.Watched
	}
	fmt.Println("update", update)
	setUpdatedBody := bson.M{"$set": update}
	_, err := db.Collection.UpdateOne(context.Background(), filter, setUpdatedBody)
	return err

}

func MakeKeyAndValuePair(movie models.Movie) map[string]interface{} {
	fmt.Println("", movie)
	fieldMap := make(map[string]interface{})

	movieField := reflect.TypeOf(movie)
	movieValue := reflect.ValueOf(movie)

	for i := 0; i < movieField.NumField(); i++ {
		fieldKey := movieField.Field(i).Name
		fieldValue := movieValue.Field(i).Interface()
		fieldMap[fieldKey] = fieldValue
	}

	return fieldMap

}
