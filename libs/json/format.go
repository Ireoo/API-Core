package iJson

import (
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
		if _type.Name() == "" {
			value := Format(res.Get(k))
			data[k] = value
		} else {
			if k == "_id" && _type.Name() == "string" {
				objID, err := primitive.ObjectIDFromHex(v.(string))
				if err != nil {
					data[k] = v
				} else {
					data[k] = objID
				}
			} else {
				data[k] = v
			}
		}
	}
	return data
}
