# Go Protocol Buffer
## Install
```
brew install protobuf && \
brew install protoc-gen-go
```

## Summary
- Protocol Buffers data format (improvement over XML and JSON). Invented by Google.
- Define a `protobuf` schema in the `.proto` file
- Compile using the `protoc` binary
```
protoc --go_out=. *.proto
```
- Rename `person.pb.go` package from `__` to `main`
```
go run main.go person.pb.go
-- OUTPUT --
[10 6 69 108 108 105 111 116 16 24]
24
Elliot
```

## Nested fields
part of our `.proto` file:
```proto
message SocialFollowers {
    int32 youtube = 1;
    int32 x = 2;
}

message Person {
    string name = 1;
    int32 age = 2;
    SocialFollowers socialFollowers = 3;
}
```


# Go gRPC Client and Server
## Install
```
brew install protoc-gen-go-grpc
```

## Summary
Making a local chat Go module and compiling with a special option to avoid `unimplemented server` errors.
```
protoc --go_out=./chat --go-grpc_out=require_unimplemented_servers=false:./chat ./chat/chat.proto
```

Client and server `.go` files added. Registers a service
with two rpcs and client sends and receives modified data to and from the server.
