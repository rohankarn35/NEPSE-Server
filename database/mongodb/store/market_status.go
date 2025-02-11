package store

import (
	"context"
	"log"
	dbmodels "nepseserver/database/models"
	"nepseserver/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MarketStatus(collection *mongo.Collection) {
	marketStatusData, err := server.GetMarketStatus()
	if err != nil {
		log.Fatal("failed to get the market status")
	}

	var marketData = dbmodels.MarketStatus{
		IsOpen: marketStatusData.IsOpen,
	}
	filter := bson.M{}
	update := bson.M{
		"$set": marketData,
	}
	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal("failed to insert or update the market status: ", err)
	}
	log.Print("Market Status Updated")

}
