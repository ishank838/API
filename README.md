# API

#Installation
import Database file db.pgsql
psql -U username dbname < dbexport.pgsql

set enviornment variables in .env

export DBHOST=localhost
export DBPORT=5432
export DBUSER=postgres
export DBPASS=password
export DBNAME=database_name

build the API using 
go build

run API using
go run main.go

Run Test using

go test API...

API Endpoints

Create:
POST  url:-  localhost:8080/create
Put user object in body

Update:
PUT  url:- localhost:8080/update/id=your_id
put user object in body

Delete:
DELETE url:- localhost:8080/delete/id=your_id

List:
GET   url:-  localhost:8080/list/limit=limit_value&offset=offset_value
