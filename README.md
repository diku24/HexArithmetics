# HexArithmatics
This repository is to learn and understand hex architecture with basic arithmetic's implementations.


# Basic Proto c commands used for rpc
1. arithmetic_svc.proto
```go
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto
```

2. number_msg.proto
```go
protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/number_msg.proto
```