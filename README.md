# swivel-server

## Dependencies
Go lang

Postgres


## Step to install

Install goose for DB migrations
``go get -u github.com/pressly/goose/cmd/goose``

Install glide
``curl https://glide.sh/get | sh``

Install packages
`` glide install ``

Build -> Builds are created in out folder
``make build``

Run Tests
``make test ``

Run DB Migrations up
``make goose up``

Run DB Migrations down
``make goose down``

TODO:

Add Lint and fix lint to makefile