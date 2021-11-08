package collections

import (
	conn "apiGo/connection"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Productos:
type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	Name        string             `bson:"name,omitempty" json:"name" form:"name"`
	Description string             `bson:"description,omitempty" json:"description" form:"description"`
	Price       float64            `bson:"price,omitempty" json:"price" form:"price"`
	Tags        []string           `bson:"tags,omitempty" json:"tags" form:"tags"`
	Stock       int32              `bson:"stock,omitempty" json:"stock" form:"stock"`
	Img         string             `bson:"img,omitempty" json:"img" form:"img"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

var CollectionProduct = conn.GetCollection(conn.Product)
