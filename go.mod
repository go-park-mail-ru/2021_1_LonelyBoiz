module server

// +heroku goVersion 1.13
go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/aws/aws-sdk-go v1.38.31
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/csrf v1.7.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.1
	github.com/microcosm-cc/bluemonday v1.0.7
	github.com/rs/cors v1.7.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)
