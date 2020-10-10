package repository

import (
	"agileengine/imagecache/pkg/utils"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitializeMongoClient() {
	if mongoClient == nil {
		clientOptions := options.Client().ApplyURI(utils.GetConfigValueFromKey("MONGO_URL"))
		mongoClientAux, err := mongo.Connect(context.TODO(), clientOptions)
		mongoClient = mongoClientAux
		utils.CheckError(err, "Error connecting mongodb")
	}
}

func StoreImages(images []interface{}) {
	collection := mongoClient.Database("agile_engine").Collection("images")
	collection.InsertMany(context.TODO(), images)
}

func DropImageCollection() {
	collection := mongoClient.Database("agile_engine").Collection("images")
	collection.Drop(context.TODO())
}
