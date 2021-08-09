package main

import (
	"fmt"
	Dbconnet "gomongo/Db.connet"
	"net/http"

	"github.com/labstack/echo"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	e := echo.New()

	Dbconnet.Connect()
	databases, err := Dbconnet.GetDatabase().ListDatabaseNames(Dbconnet.GetDatabasectx(), bson.M{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer Dbconnet.GetDatabase().Disconnect(Dbconnet.GetDatabasectx())
	fmt.Println(databases)

	e.GET("/items", func(c echo.Context) error {
		return c.JSON(http.StatusOK, databases)
	})

	e.Logger.Fatal(e.Start(":8000"))

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://root:gc742899@cluster0.mqkff.mongodb.net/TEST?authSource=admin&replicaSet=atlas-fe50vv-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)
}
