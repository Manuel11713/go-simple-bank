## Prerequisites

**Docker**
https://www.docker.com/

- postgres

```bash
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.0-alpine

docker exec -it postgres15 psql -U root  # just to test the connection
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
