# grpc_practice

## server start
```shell
$ go run server/main.go
```

cf. https://grpc.io/docs/languages/go/quickstart/

## client start
```shell
$ go run client/main.go
```

### result
```shell
$ go run client/main.go 1
2021/09/25 12:34:52 name: id:1 name:"Wasabi"

$ go run client/main.go 2
2021/09/25 12:34:55 name: id:2 name:"Karashi"

$ go run client/main.go 3
2021/09/25 12:34:58 getting user failed: rpc error: code = NotFound desc = user not found
exit status 1
```
