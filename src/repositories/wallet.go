package repositories

import (
	"context"
	"github/kdswto/webserver_example/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type WalletRepository struct {
	Database *mongo.Database
}

func (r *WalletRepository) GetByUserId(userId int) *models.Wallet {
	if r.Database == nil {
		log.Fatal("WalletRepository.GetById: Database not initialize")
	}

	filter := bson.D{{"user_id", userId}}
	wallet := &models.Wallet{}
	err := r.Database.Collection("wallet").FindOne(context.Background(), filter).Decode(&wallet)
	if err != nil {
		return nil
	}

	return wallet
}
