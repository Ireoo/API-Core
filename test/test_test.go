package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/Ireoo/API-Core/libs/conf"
	// "github.com/Ireoo/API-Core/libs/mongodb"
	"github.com/Ireoo/API-Core/libs/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "go.mongodb.org/mongo-driver/bson"
)

func initializeRouter() *gin.Engine { // 修改函数名以避免重定义错误
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/:table/:mode", func(c *gin.Context) {
		router.Table(c, "secret", true)
	})
	r.GET("/:table/:mode", func(c *gin.Context) {
		router.TableGet(c, "secret", true)
	})
	return r
}

func TestRouterEndpoints(t *testing.T) { // 修改函数名以避免重定义错误
	t.Run("Test POST /:table/:mode", func(t *testing.T) {
		r := initializeRouter() // 调用修改后的函数
		req, _ := http.NewRequest("POST", "/testTable/testMode", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		// 根据实际返回内容进行断言
		// assert.Contains(t, w.Body.String(), "expected response")
	})

	t.Run("Test GET /:table/:mode", func(t *testing.T) {
		r := initializeRouter() // 调用修改后的函数
		req, _ := http.NewRequest("GET", "/testTable/testMode", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		// 根据实际返回内容进行断言
		// assert.Contains(t, w.Body.String(), "expected response")
	})

	// 添加更多测试用例以覆盖router中的所有功能
}

// func TestMongoDB(t *testing.T) {
// 	// 初始化MongoDB连接
// 	err := mongo.New("")
// 	assert.NoError(t, err)

// 	t.Run("Test Insert and FindOne", func(t *testing.T) {
// 		doc := bson.M{"name": "test", "value": 1}
// 		_, err := mongo.Insert("testDB", "testCollection", doc)
// 		assert.NoError(t, err)

// 		var result bson.M
// 		err = mongo.FindOne("testDB", "testCollection", bson.M{"name": "test"}, &conf.Other{}, &result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, doc["name"], result["name"])
// 		assert.Equal(t, doc["value"], result["value"])
// 	})

// 	t.Run("Test Count", func(t *testing.T) {
// 		count, err := mongo.Count("testDB", "testCollection", bson.M{"name": "test"})
// 		assert.NoError(t, err)
// 		assert.Equal(t, int64(1), count)
// 	})

// 	t.Run("Test IsEmpty", func(t *testing.T) {
// 		isEmpty := mongo.IsEmpty("testDB", "testCollection")
// 		assert.False(t, isEmpty)
// 	})

// 	t.Run("Test Remove", func(t *testing.T) {
// 		err := mongo.Remove("testDB", "testCollection", bson.M{"name": "test"})
// 		assert.NoError(t, err)

// 		count, err := mongo.Count("testDB", "testCollection", bson.M{"name": "test"})
// 		assert.NoError(t, err)
// 		assert.Equal(t, int64(0), count)
// 	})
// }
