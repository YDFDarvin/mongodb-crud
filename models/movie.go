package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty" bson:"title,omitempty"`
	Year  int                `json:"year,omitempty" bson:"year,omitempty"`
}

func InsertOne[E Movie](doc E) {
	collection := mongoClient.Database(dbName).Collection(colName)

	inserted, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	defer (func() {
		log.Println("[InsertDoc] added document: ", inserted.InsertedID)
	})()
}

func InsertMany[E Movie](docs []E) {
	docsCasted := make([]interface{}, len(docs))
	for i, v := range docs {
		docsCasted[i] = v
	}

	collection := mongoClient.Database(dbName).Collection(colName)

	inserted, err := collection.InsertMany(context.TODO(), docsCasted)
	if err != nil {
		panic(err)
	}
	defer (func() {
		log.Println("[InsertDocs] added documents: ", inserted.InsertedIDs)
	})()
}

// Just for now not generic
func UpdateOne[E Movie](docId string, doc E) {
	id, err := primitive.ObjectIDFromHex(docId)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": id}
	// TODO: Just for now not generic, Movie() is type casting
	udpate := bson.M{"$set": bson.M{"title": Movie(doc).Title, "year": Movie(doc).Year}}
	defer (func() {
		log.Println("[UpdateOne] query: ", udpate)
	})()

	collection := mongoClient.Database(dbName).Collection(colName)

	result, err := collection.UpdateOne(context.TODO(), filter, udpate)
	if err != nil {
		panic(err)
	}
	defer (func() {
		log.Println("[UpdateOne] updated documents: ", result.UpsertedID)
	})()
}

func DeleteOne[E Movie](docId string) {
	id, err := primitive.ObjectIDFromHex(docId)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": id}

	collection := mongoClient.Database(dbName).Collection(colName)
	deleted, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer (func() {
		log.Println("[DeleteOne] deleted document: ", deleted.DeletedCount)
	})()
}

func FindOne[E Movie](key string) E {
	var doc E

	// TODO: Just for now not generic
	filter := bson.D{{"title", key}}

	collection := mongoClient.Database(dbName).Collection(colName)
	// Write by reference document
	err := collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		panic(err)
	}
	defer (func() {
		log.Println("[FindOne] found document: ", doc)
	})()

	return doc
}

func FindMany[e Movie](key string) []E {
	var docs []E

	// TODO: Just for now not generic
	filter := bson.D{{"title", key}}

	collection := mongoClient.Database(dbName).Collection(colName)
	// Write by reference document
	cursor, err := collection.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		panic(err)
	}
	defer (func() {
		log.Println("[FindMany] found documents: ", docs)
	})()

	for cursor.Next(context.TODO()) {
		var acc bson.D
		err := cursor.Decode(&acc)
		if err != nil {
			panic(err)
		}
		docs = append(docs, acc)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}

	return docs
}
