package iJson

import (
	"reflect"
	"encoding/json"
	"github.com/gookit/color"

	simplejson "github.com/bitly/go-simplejson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Ireoo/API-Core/libs/debug"
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
			b, err := json.Marshal(v)
			if err != nil {
				debug.Error("[ERROR]" + " " + color.FgDefault.Render(err))
			}
			value, err := simplejson.NewJson(b)
			if err != nil {
				debug.Error("[ERROR]" + " " + color.FgDefault.Render(err))
			}
			if innerArr, err := value.Array(); err == nil {
				value := FormatArray(innerArr)
				result = append(result, value)
			} else {
				value := Format(value)
				result = append(result, value)
			}
		} else {
			result = append(result, v)
		}
	}
	return result
}
