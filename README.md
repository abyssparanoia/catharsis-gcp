# catharsis-gcp

## what is this

```
the boilerplate for monorepo application
```

- Base project is https://github.com/golang-standards/project-layout

## Apps

| Package                             | Localhost             | Prodction  |
| :---------------------------------- | :-------------------- | :--------- |
| **[[REST] default](./cmd/default)** | http://localhost:8080 | default.\* |

## development

### Preparation

<!--
- generate rsa pem file

```bash
> openssl genrsa -out ./secret/catharsis-gcp.rsa 1024
> openssl rsa -in ./secret/catharsis-gcp.rsa  -pubout > ./secret/catharsis-gcp.rsa.pub
``` -->

- environment (using dotenv)
  - you should fix a host to default-db if you use docker-compose as server runtime

```bash
> cp .tmpl.env.default .env.default
```

### server starting

- local

```bash
> realize start
```

- docker

```bash
# build image
> docker-compose build

# container start
> docker-compose up -d
```

- example of default server

```bash
> curl --request GET 'http://localhost:8080/ping'
```

<!-- ### database

- generate server code by sql boiler

```bash
> sqlboiler -c ./db/authentication/sqlboiler.toml -o ./pkg/dbmodels/authentication psql
``` -->

## production

### build

```bash
> docker build -f ./docker/production/default/Dockerfile .
```

## structure

```bash
.
├── LICENSE.md
├── Makefile
├── README.md
├── authentication
│   ├── Makefile
│   ├── domain
│   │   ├── model
│   │   │   └── user.go
│   │   └── repository
│   │       ├── mock
│   │       │   └── user.go
│   │       └── user.go
│   ├── handler
│   │   └── authentication.go
│   ├── infrastructure
│   │   ├── internal
│   │   │   └── entity
│   │   │       ├── base.go
│   │   │       └── user.go
│   │   └── repository
│   │       └── user_mock.go
│   ├── intercepter
│   │   └── authentication.go
│   └── service
│       ├── authentication.go
│       ├── authentication_impl.go
│       ├── mock
│       │   └── authentication.go
│       └── test
│           └── authentication_impl_test.go
├── bff
│   ├── domain
│   │   ├── model
│   │   └── repository
│   │       └── authentication.go
│   ├── handler
│   │   ├── authentication.go
│   │   └── ping.go
│   ├── infrastructure
│   │   └── repository
│   │       └── authentication_impl.go
│   ├── middleware
│   │   └── access_control.go
│   └── service
│       ├── authentication.go
│       └── authentication_impl.go
├── cmd
│   ├── README.md
│   ├── authentication
│   │   ├── authentication-server
│   │   ├── dependency.go
│   │   ├── env.go
│   │   └── main.go
│   └── bff
│       ├── bff-server
│       ├── dependency.go
│       ├── env.go
│       ├── main.go
│       └── routing.go
├── docker
│   ├── development
│   │   └── Dockerfile
│   └── production
│       └── authentication
│           └── Dockerfile
├── docker-compose.yml
├── docs
├── go.mod
├── go.sum
├── pkg
│   ├── README.md
│   ├── errcode
│   │   ├── error.go
│   │   └── model.go
│   ├── httpaccesslog
│   │   └── middleware.go
│   ├── jwtauth
│   │   ├── config.go
│   │   ├── context.go
│   │   ├── mock
│   │   │   ├── sign_service.go
│   │   │   └── verify_service.go
│   │   ├── model.go
│   │   ├── sign_client.go
│   │   ├── sign_service.go
│   │   ├── sign_service_impl.go
│   │   ├── verify_client.go
│   │   ├── verify_service.go
│   │   └── verify_service_impl.go
│   ├── log
│   │   ├── config.go
│   │   └── logger.go
│   ├── parameter
│   │   ├── json.go
│   │   └── url.go
│   └── renderer
│       ├── handler.go
│       └── model.go
├── proto
│   ├── authentication.pb.go
│   └── authentication.proto
└── secret
    ├── catharsis-gcp.rsa
    └── catharsis-gcp.rsa.pub

```
