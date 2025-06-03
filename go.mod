module chat_gopher

go 1.20

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/lib/pq v1.10.9
	github.com/sirupsen/logrus v1.9.3
	github.com/swaggo/http-swagger v1.3.1 // для Swagger UI
	github.com/swaggo/swag v1.16.4 // swag CLI
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/swaggo/swag => github.com/swaggo/swag v1.16.4
