# task
task

## to start locally 
```
cp .env.example example
go run main.go
```

## to start by docker
```
cp .env.example example
docker compose up



# post
curl --location 'localhost:3000/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "name": "valid status",
    "status": 1
}'

# get all tasks
curl --location 'localhost:3000/tasks'

# update task
curl --location --request PUT 'localhost:8080/tasks/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "999",
    "status": 0
}'

# delete task
curl --location --request DELETE 'localhost:8080/tasks/1'
```
