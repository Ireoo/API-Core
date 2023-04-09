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
		if _type.Kind() == reflect.Interface || _type.Kind() == reflect.Map {
			if arr, err := res.Get(k).Array(); err == nil {
				value := FormatArray(arr)
				data[k] = value
			} else {
				value := Format(res.Get(k))
				data[k] = value
			}
		} else {
			if k == "_id" && _type.Kind() == reflect.String {
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

func FormatArray(arr []interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range arr {
		_type := reflect.TypeOf(v)
		if _type.Kind() == reflect.Interface || _type.Kind() == reflect.Map {
			if innerArr, err := simplejson.NewJson(v).Array(); err == nil {
				value := FormatArray(innerArr)
				result = append(result, value)
			} else {
				value := Format(simplejson.NewFromAny(v))
				result = append(result, value)
			}
		} else {
			result = append(result, v)
		}
	}
	return result
}
