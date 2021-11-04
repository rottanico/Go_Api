package controller

import (
	coll "apiGo/collections"
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/form"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct {
}

func (r *Controller) Connection(c echo.Context) error {

	return c.JSON(http.StatusOK, "bienvenido @User")
}
func (r *Controller) GetProducts(c echo.Context) error {
	var results []coll.Product

	cursor, erro := coll.CollectionProduct.Find(context.TODO(), bson.D{})
	if erro != nil {
		c.Logger().Error(erro)
		return c.JSON(http.StatusNotFound, erro)
	}
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) <= 1 {
		return c.JSON(http.StatusOK, results[0])
	}
	return c.JSON(http.StatusOK, results)
}
func (r *Controller) PostProduct(c echo.Context) error {
	p := new(coll.Product)
	if len(c.FormValue("name")) != 0 {
		decoder := form.NewDecoder()
		form, _ := c.FormParams()
		err := decoder.Decode(&p, form)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}
	} else {
		err := c.Bind(&p)

		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	if len(p.Name)&len(p.Description) != 0 {
		coll.CollectionProduct.InsertOne(context.TODO(), &p)
		return c.JSON(http.StatusOK, p)
	} else {
		return c.String(http.StatusBadRequest, "valor no encontrado")
	}

}
func (r *Controller) GetProduct(c echo.Context) error {
	p := new(coll.Product)
	Id := c.QueryParam("id")
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNoContent, err)
	}
	filter := bson.M{"_id": id}
	err = coll.CollectionProduct.FindOne(context.TODO(), filter).Decode(&p)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, p)
}
func (r *Controller) DeleteProduct(c echo.Context) error {
	Id := c.QueryParam("id")
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNoContent, err)
	}
	filter := bson.M{"_id": id}

	value, errr := coll.CollectionProduct.DeleteOne(context.TODO(), filter)
	if errr != nil {
		c.Logger().Error(errr)
		return c.JSON(http.StatusRequestTimeout, errr)
	}
	return c.JSON(http.StatusOK, value)

}

func (r *Controller) GetProductsBy(c echo.Context) error {
	result := new([]coll.Product)
	offset, _ := strconv.ParseInt(c.QueryParam("offset"), 10, 64)

	count, _ := strconv.ParseInt(c.QueryParam("count"), 10, 64)

	opt := options.FindOptions{
		Skip:  &offset,
		Limit: &count}
	cursor, err := coll.CollectionProduct.Find(context.TODO(), bson.M{}, &opt)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusRequestTimeout, err)
	}
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusRequestTimeout, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (r *Controller) InserMany(c echo.Context) error {
	var Products []interface{}
	err := c.Bind(&Products)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusRequestTimeout, err)
	}
	result, errr := coll.CollectionProduct.InsertMany(context.TODO(), Products)
	if errr != nil {
		c.Logger().Error(errr)
		return c.JSON(http.StatusRequestTimeout, errr)
	}
	return c.JSON(http.StatusOK, result)
}
func (r *Controller) UpdateOne(c echo.Context) error {
	return c.String(http.StatusMethodNotAllowed, "in process")
}
