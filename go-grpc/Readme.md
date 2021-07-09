事前安裝
```
sudo apt install protobuf-compiler
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

server side
```
go run server.go
```

client side
```
go run client/client.go
```
