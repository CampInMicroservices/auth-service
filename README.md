# Auth service

## gRPC API

Auth PROTO:

```
{
    "email": "admin@test.com",
    "password": "test"
}
```

### Modifying Auth Proto

If `proto/auth.proto` is modified, run:

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```

If protoc is not found by bash, run:

```
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```