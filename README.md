# mb
Repository dedicated to study GO

## Steps to run the application:

### REPL Application
1 - If you gonna use the aplication without database (REPL) just run `go run main.go` 

2 - Choose option 1 at first menu

### Web Application using database

1 - First you need be on main directory

2 - Start the postgres container using the command:
`docker-compose up -d`

3 - Using the user: root and password: root connect to database that will start by default on port 5432

4 - Run the sql script to build the tables:
```
CREATE TABLE to_do_list (
	id serial primary key,
	ListId int NOT NULL,
	Item varchar NOT NULL,
	Done boolean);
```

5 - Start the application using `go run main.go`

ps: If you need to use the web application the endpoints are:

- add an item using method `POST` : `localhost:5012/insert/`
- remove an item using method `DELETE` : `localhost:5012/remove/`
- add an item using method `PATCH` : `localhost:5012/update/`
- list itens using method `GET` : `localhost:5012/list/`