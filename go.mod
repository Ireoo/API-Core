module import

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8
	github.com/mattn/go-colorable v0.1.0
	github.com/mattn/go-isatty v0.0.4
	github.com/valyala/bytebufferpool v1.0.0
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4
	golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
	golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/yaml.v2 v2.2.2
	qiyi.io/libs v1.0.0
)

replace qiyi.io/libs v1.0.0 => ./libs
