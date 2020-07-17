# API

## Setup

import Database file db.pgsql
```bash
psql -U username dbname < dbe.pgsql
```
Set enviornment variables in .env
```bash
export DBHOST=localhost
export DBPORT=5432
export DBUSER=postgres
export DBPASS=password
export DBNAME=database_name
```

Build the API using 
```bash
go build
```

Run API using
```bash
go run main.go
```
Run Test using
```bash
go test API...
```
## API Endpoints

Create:

**POST**  url:-  localhost:8080/create
Put user object in body

Update:

**PUT**  url:- localhost:8080/update/id=your_id

put user object in body

Delete:

**DELETE** url:- localhost:8080/delete/id=your_id

List:

**GET**   url:-  localhost:8080/list/limit=limit_value&offset=offset_value
