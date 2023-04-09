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
			if arr, ok := res.Get(k).Array(); ok {
				value := FormatArray(arr)
				data[k] = value
			} else {
				value := Format(res.Get(k))
				data[k] = value
			}
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

func FormatArray(arr []*simplejson.Json) []interface{} {
	result := []interface{}{}
	for _, v := range arr {
		_type := reflect.TypeOf(v)
		if _type.Name() == "" {
			if innerArr, ok := v.Array(); ok {
				value := FormatArray(innerArr)
				result = append(result, value)
			} else {
				value := Format(v)
				result = append(result, value)
			}
		} else {
			result = append(result, v.Interface())
		}
	}
	return result
}
