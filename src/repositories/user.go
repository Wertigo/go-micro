package repositories

import (
	"context"
	"github/kdswto/webserver_example/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	Database *mongo.Database
}

func (r *UserRepository) GetAll() []*models.User {
	if r.Database == nil {
		log.Fatal("UserRepository.GetAll: Database not initialize")
	}

	var users []*models.User
	ctx := context.Background()
	usersCursor, err := r.Database.Collection("users").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer usersCursor.Close(ctx)

	for usersCursor.Next(ctx) {
		user := &models.User{}
		err = usersCursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

func (r *UserRepository) GetById(id int) *models.User {
	if r.Database == nil {
		log.Fatal("UserRepository.GetAll: Database not initialize")
	}

	ctx := context.Background()
	filter := bson.D{{"id", id}}
	user := &models.User{}
	err := r.Database.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}

	return user
}
