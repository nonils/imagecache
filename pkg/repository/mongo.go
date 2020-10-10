package repository

import (
	"agileengine/imagecache/pkg/model"
	"agileengine/imagecache/pkg/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
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
	models := mongo.IndexModel{
		Keys: bsonx.Doc{{Key: "id", Value: bsonx.String("text")},
			{Key: "author", Value: bsonx.String("text")},
			{Key: "camera", Value: bsonx.String("text")},
			{Key: "tags", Value: bsonx.String("text")},
			{Key: "croppedpicture", Value: bsonx.String("text")},
			{Key: "fullpicture", Value: bsonx.String("text")},
		},
	}
	collection.Indexes().CreateOne(context.TODO(), models)
}

func FindImageById(id string) *model.Image {
	collection := mongoClient.Database("agile_engine").Collection("images")
	result := new(model.Image)
	err := collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	utils.CheckError(err, "Error tyring to get the result from mongo")
	return result
}

func SearchImageByTerm(searchTerm string) []model.Image {
	var res []model.Image
	collection := mongoClient.Database("agile_engine").Collection("images")
	cursor, err := collection.Find(context.TODO(), bson.M{"$text": bson.M{"$search": searchTerm}})
	utils.CheckError(err, "Error tyring to get the result from mongo")
	err = cursor.All(context.TODO(), &res)
	utils.CheckError(err, "Error trying to deserialize all from mongo")
	if res == nil {
		res = make([]model.Image, 0)
	}
	return res
}
