package conf

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoDB struct {
	Host      string        `yaml:"Host"`
	Auth      string        `yaml:"Auth"`
	Username  string        `yaml:"Username"`
	Password  string        `yaml:"Password"`
	Timeout   time.Duration `yaml:"Timeout"`
	PoolLimit int           `yaml:"PoolLimit"`
	URI       string        `yaml:"URI"`
}

type Input struct {
	Where bson.M      `json:"where,omitempty" form:"where,omitempty" query:"where,omitempty"`
	Data  bson.M      `json:"data,omitempty" form:"data,omitempty" query:"data,omitempty"`
	Other interface{} `json:"other,omitempty" form:"other,omitempty" query:"other,omitempty"`
	Table string      `json:"table,omitempty" form:"table,omitempty" query:"table,omitempty"`
	Mode  string      `json:"mode,omitempty" form:"mode,omitempty" query:"mode,omitempty"`
	Auth  string      `json:"auth,omitempty" form:"auth,omitempty" query:"auth,omitempty"`
	App   string      `json:"app,omitempty" form:"app,omitempty" query:"app,omitempty"`
}

type Other struct {
	Page       int64    `json:"page,omitempty" form:"page,omitempty" query:"page,omitempty"`
	Skip       int64    `json:"skip,omitempty" form:"skip,omitempty" query:"skip,omitempty"`
	Limit      int64    `json:"limit,omitempty" form:"limit,omitempty" query:"limit,omitempty"`
	Show       bson.M   `json:"show,omitempty" form:"show,omitempty" query:"show,omitempty"`
	Distinct   bson.M   `json:"distinct,omitempty" form:"distinct,omitempty" query:"distinct,omitempty"`
	Sort       bson.M   `json:"sort,omitempty" form:"sort,omitempty" query:"sort,omitempty"`
	Indexes    []string `json:"indexes,omitempty" form:"indexes,omitempty" query:"indexes,omitempty"`
	Upsert     bool     `json:"upsert" form:"upsert" query:"upsert"`
	Multi      bool     `json:"multi" form:"multi" query:"multi"`
	SortFormat bson.D   `json:"sortFormat,omitempty" form:"sortFormat,omitempty" query:"sortFormat,omitempty"`
}

type AppInfo struct {
	Id   string `bson:"_id"`
	Uuid string `bson:"uuid"`
}

type UserInfo struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type Result struct {
	Success bool
	Data    interface{}
}
