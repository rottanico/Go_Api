package collections

import (
	conn "apiGo/connection"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Productos:
type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	Name        string             `bson:"name" json:"name" form:"name"`
	Description string             `bson:"description" json:"description" form:"description"`
	Price       float64            `bson:"price" json:"price" form:"price"`
	Categorys   []Category         `bson:"categorys" json:"categorys" form:"categorys"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Stock       int16              `bson:"stock" json:"stock" form:"stock"`
}
type Category struct {
	Name string `bson:"name" json:"name" form:"name"`
}

var CollectionProduct = conn.GetCollection(conn.Product)
var CollectionCategory = conn.GetCollection(conn.Category)
