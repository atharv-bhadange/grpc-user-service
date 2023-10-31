## GRPC User Service

Sample GRPC service with mock database as map of users.


### Command to generate stubs

```
protoc --go_out=. --go-grpc_out=. proto/user.proto 
```

### Command to run

```
docker compose up -d
```
OR

```
go build -o main && ./main
```

### Command to test

```
go test -v ./tests
```
