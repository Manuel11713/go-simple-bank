## Prerequisites

**Docker**
https://www.docker.com/

- postgres

```bash
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.0-alpine

docker exec -it postgres15 psql -U root  # just to test the connection
```

**SQLC**
https://sqlc.dev/

```bash
brew install sqlc
```

**Table plus**
https://tableplus.com/

**dbdiagram**
https://dbdiagram.io/home

**go-migrate**

```bash
brew install golang-migrate
migrate create -ext sql -dir db/migration -seq init_schema
```

**Go-postgres driver**

```bash
go get github.com/lib/pq
```

**Testify**

```bash
go get github.com/stretchr/testify
```

### Creating a new db inside docker postgres image

```bash
docker exec -it postgres15 /bin/sh

# createdb --username=root --owner=root simple_bank
# psql simple_bank # just to test the connection
# dropdb simple_bank
```

we can concanete commands:

```bash
docker exec -it postgres15 createdb --username=root --owner=root simple_bank

docker exec -it postgres15 psql -U root simple_bank
```

_\*Makefile is ready to run docker postgres image and crate the database_

**Setup SQLC**

```bash
sqlc init
```

Schemas will be the migration folder containing the CREATE DATABASE statements.
Queries will be the folder containing the database queries
Path will be the generated code

```bash
sqlc generate
```

or

```bash
make sqlc
```

**Mocking Data**

gomock:

https://github.com/golang/mock

```bash
mockgen -destination db/mock/store.go -package mockdb github.com/Manuel11713/simple-bank/db/sqlc Store
```

**Create New Migration**

```bash
migrate create -ext sql -dir db/migration -seq add_users
```
