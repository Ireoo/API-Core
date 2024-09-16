package iJson

import (
	"encoding/json"
	"github.com/gookit/color"
	"reflect"

	"github.com/Ireoo/API-Core/libs/debug"
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
			if k == "_id" { // 添加对 uuid 字段的处理
				switch val := v.(type) {
				case string:
					objID, err := primitive.ObjectIDFromHex(val)
					if err != nil {
						data[k] = val
					} else {
						data[k] = objID
					}
				case json.Number:
					strVal := val.String() // 或者使用 .Int64(), .Float64() 进行数字转换
					objID, err := primitive.ObjectIDFromHex(strVal)
					if err != nil {
						data[k] = strVal
					} else {
						data[k] = objID
					}
				default:
					data[k] = v
				}
			} else if k == "uuid" {
				switch val := v.(type) {
				case json.Number:
					strVal := val.String() // 或者使用 .Int64(), .Float64() 进行数字转换
					data[k] = strVal
				default:
					data[k] = v
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
			} else {
				value, err := simplejson.NewJson(b)
				if err != nil {
					debug.Error("[ERROR]" + " " + color.FgDefault.Render(err))
				} else {
					if innerArr, err := value.Array(); err == nil {
						value := FormatArray(innerArr)
						result = append(result, value)
					} else {
						value := Format(value)
						result = append(result, value)
					}
				}
			}
		} else {
			result = append(result, v)
		}
	}
	return result
}
