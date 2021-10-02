package ijson

import (
	// "encoding/json"
	"fmt"

	"reflect"

	simplejson "github.com/bitly/go-simplejson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Format(res *simplejson.Json) bson.M {
	data := bson.M{}
	r, _ := res.Map()
	for k, v := range r {
		_type := reflect.TypeOf(v)
		// fmt.Println("_type ->", _type.Name())
		if _type.Name() == "" {
			value := Format(res.Get(k))
			data[k] = value
		} else {
			if k == "_id" && _type.Name() == "string" {
				// fmt.Println(k)
				// var documentID bson.RawValue
				// documentID.Type = 7
				// documentID.Value = []byte(v.(string))
				objID, _ := primitive.ObjectIDFromHex(v.(string))

				// fmt.Println(k, objID)
				data[k] = objID
			} else {
				fmt.Println(k, v)
				data[k] = v
			}
		}
	}
	return data
}
