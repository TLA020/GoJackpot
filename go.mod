module goprac

go 1.14

require (
	github.com/TLA020/go_marketcap-client v0.0.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofiber/cors v0.2.2
	github.com/gofiber/fiber v1.14.4
	github.com/gofiber/jwt v0.2.0
	github.com/gofiber/websocket v0.5.1
	github.com/goombaio/namegenerator v0.0.0-20181006234301-989e774b106e
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/klauspost/compress v1.10.11 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	golang.org/x/sys v0.0.0-20200828194041-157a740278f4 // indirect
)

replace (
 github.com/TLA020/go_marketcap-client => ../go_marketcap-client
)
