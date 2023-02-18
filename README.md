# gocom
In this `gocom` monorepo we'll demonstrate Golang based microservices following clean architecture for a ecommerce site.

## This monorepo's structure
* [auth](./auth/README.md) : Auth service which have all the endpoints related to authn/authz.
* [ecom](./ecom/README.md) : ecom service which have all the endpoints to do CRUD operations for a ecommerce site.
* [schemas](./schemas/README.md) :
    - DB migrations
    - DB schemas
    - SQLC queries
    - code generations

## Tech Stack & Tools
- Golang
- Gin Framework
- Ginkgo
- Gomega
- Golang Mockgen
- PostgreSQL
- SQLC
- Golang Migrate
- Docker
- Kubernetes
- GH Actions
- ArgoCD
- OpenAPI Schema

## How to run this project locally?
- clone this repo & go to the root directory (`cd ~/<>/gocom`)
- `docker compose up`:
    - to run all the services according to `docker-compose.yaml` file use this command
    - if you change anything in the proeject and wanna run with that then use `docker compose up --build`
    - you can run any specific service by `docker compose up <service_name_from_docker_compose_file>
- [optional] you can connect to the DB via any(Like: `table plus`) GUI providing the connection info given in the `docker-compose.yaml` file's `db` service:
    ```
        POSTGRES_USER: gocom
        POSTGRES_PASSWORD: gocom123
        POSTGRES_DB: gocomdb
    ```
- open another terminal to apply the migrations files into DB
- `cd schemas`
- `export POSTGRESQL_URL='postgresql://gocom:gocom123@localhost:5432/gocomdb?sslmode=disable'` 
- `migrate -database $POSTGRESQL_URL -path db/migrations up` : apply the migrations files
- [🎉] now the projects setup is successful, for checking the API endpoints of each service pls have a look into each service's README.md file and test those endpoints from `postman` or etc.

## Auth Service's RESTful API Endpoints
- check [here](./auth/openapi.yaml) for OpenAPI schema of auth service's api endpoints
- also look into the [examples](./auth/README.md)

## Ecom Service's RESTful API Endpoints
- check [here](./ecom/openapi.yaml) for OpenAPI schema of ecom service's api endpoints
- also look into the [examples](./ecom/README.md)
