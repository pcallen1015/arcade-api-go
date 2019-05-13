package wins

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"pcallen1015/arcade-api-go/database"
)

const (
	dbCollection = "wins"
)

// WinDocument describes the structure of a Win
type WinDocument struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Player   string             `json:"player,omitempty" bson:"player,omitempty"`
	Game     string             `json:"game,omitempty" bson:"game,omitempty"`
	PlayedOn time.Time          `json:"playedOn,omitempty" bson:"playedOn,omitempty"`
}

func list() ([]WinDocument, error) {
	var results []WinDocument
	cursor, err := database.GetDb().Collection(dbCollection).Find(context.Background(), bson.D{})
	if err != nil {
		return results, err
	}

	for cursor.Next(context.Background()) {
		var result WinDocument
		err = cursor.Decode(&result)
		if err == nil {
			results = append(results, result)
		}
	}

	return results, nil
}

func create(data []byte) (WinDocument, error) {
	var doc WinDocument
	err := json.Unmarshal(data, &doc)
	if err != nil {
		return doc, err
	}
	result, err := database.GetDb().Collection(dbCollection).InsertOne(context.Background(), doc)
	if err != nil {
		return doc, err
	}

	doc.ID = result.InsertedID.(primitive.ObjectID)

	return doc, nil
}
