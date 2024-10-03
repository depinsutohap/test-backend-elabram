# Start DEPEDENCY SERVICE WITH DOCKER

```bash
docker run  --name mysql -e MYSQL_ROOT_PASSWORD=root  -p 3307:3306 -v /etc/docker/test-mysql:/etc/mysql/conf.d -v final-mysql-data:/var/lib/mysql -d mysql
```
```bash
docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
```

## Installation
Run migrate.sql to your mysql database.

Please install Golang First.
Setup .env 
```bash
go run ./main.go
```

You can try the API with postman, I attach the postman json
## Explation optimize query
I will make use indexing first for the column that used for filter like id's and date's and then create materialized views to optimize the query.

Test for backend at elabram