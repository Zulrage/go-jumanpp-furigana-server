package models

import "gopkg.in/mgo.v2/bson"

type (
    // User represents the structure of our resource
    Furigana struct {
      MongoId bson.ObjectId `json:"mongoId" bson:"_id"`
      Title   string        `json:"title" bson:"title"`
      Content string        `json:"content" bson:"content"`
      Date    string        `json:"date" bson:"date"`
      Type    string        `json:"type" bson:"type"`
      Chapter    string        `json:"chapter" bson:"chapter"`
  }
)
