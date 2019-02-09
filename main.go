package main

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"localhost/mongo"
)

func main() {
	var result bson.M
	err := mongo.FindOne("api", "user", bson.M{"username": "18551410359"}, bson.M{"_id": 0}, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
