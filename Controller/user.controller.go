package controller

import (
	"fmt"
	Dbconnet "gomongo/Db.connet"
	"gomongo/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo"
)

func SetupItemController(e *echo.Echo) {
	e.GET("/users", GetAll)
	e.POST("/user", Insert)

	e.GET("user/:id", GetuserId)
	e.PUT("/user/:id", updateuserId)
	e.DELETE("/user/:id", DeluserId)
	e.GET("/", test)
}

func test(c echo.Context) error {
	a := "dddd"
	return c.JSON(http.StatusOK, a)
}

func GetAll(c echo.Context) error {
	var users []bson.M

	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
	cur, err := coletion.Find(Dbconnet.GetDatabasectx(), bson.M{})
	if err != nil {
		panic(err)
	}
	err = cur.All(Dbconnet.GetDatabasectx(), &users)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, users)
}

func Insert(c echo.Context) error {
	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
	testuser := model.User{
		// Name: "",
		// City: "Samut",
		// Age:  22, //////,ต้องมีอยุ่ตัวสุดท้าย ถ้าเป็น Go
	}

	if err := c.Bind(&testuser); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	res, err := coletion.InsertOne(Dbconnet.GetDatabasectx(), testuser)
	if err != nil {
		panic(err)
	}
	fmt.Print(res)
	return c.String(http.StatusOK, "InsertOk")
}

func GetuserId(c echo.Context) error {
	var user bson.M
	id := c.Param("id")

	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
	// objid, err := primitive.ObjectIDFromHex("6112341143062e8d442b584c")
	objid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}
	err = coletion.FindOne(Dbconnet.GetDatabasectx(), bson.M{"_id": objid}).Decode(&user)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, user)
}

func DeluserId(c echo.Context) error {
	// var user bson.M
	id := c.Param("id")

	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
	// objid, err := primitive.ObjectIDFromHex("6112341143062e8d442b584c")
	objid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}
	// var result = *mongo.DeleteResult
	_, err = coletion.DeleteOne(Dbconnet.GetDatabasectx(), bson.M{"_id": objid})
	if err != nil {
		panic(err)
	}
	// fmt.Println(result)

	return c.String(http.StatusOK, "Delsusecfull")

	// return c.JSON(http.StatusOK, user)
}

func updateuserId(c echo.Context) error {
	// var user bson.M
	id := c.Param("id")
	testuser := model.User{}

	coletion := Dbconnet.GetDatabase().Database("TESTMVC").Collection("user")
	// objid, err := primitive.ObjectIDFromHex("6112341143062e8d442b584c")
	objid, err := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&testuser); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err != nil {
		panic(err)
	}
	// res, err := coletion.UpdateOne(
	// 	Dbconnet.GetDatabasectx(),
	// 	bson.M{"_id": objid},
	// 	bson.D{
	// 		{Key: "$set", Value: bson.D{{Key: "city", Value: "Bank"}}},
	// 	})
	res, err := coletion.UpdateOne(Dbconnet.GetDatabasectx(), bson.M{"_id": objid}, bson.M{"$set": &testuser})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ModifiedCount)

	return nil
}
