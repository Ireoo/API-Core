module import

go 1.13

require (
	qiyi.io/basic v1.0.0
  qiyi.io/conf v1.0.0
  qiyi.io/core v1.0.0
  qiyi.io/mongo v1.0.0
)

replace qiyi.io/basic v1.0.0 => ./basic/basic.go
replace qiyi.io/conf v1.0.0 => ./conf/conf.go
replace qiyi.io/core v1.0.0 => ./core/core.go
replace qiyi.io/mongo v1.0.0 => ./mongo/mongo.go