
create a new db migration
```bash
$ migrate create -ext sql -dir db/migration -seq <name-file>
```

map all env variables on AWS secrets manager into app.env file
```bash
$ aws secretsmanager get-secret-value --secret-id simple_bank \
 --query SecretString --output text \
 | jq -r 'to_entries|map("(\.key)=\(.value)")|.[]' \
 > app.env
```

## GRPC

- [grpc documentation](https://grpc.io/docs/languages/go/quickstart/)
- [Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)

## 4 Types of gRPC

1. Unary gRPC
2. Client Streaming gRPC
3. Server Streaming gRPC
4. Bidirectional streaming gRPC

## Evans testing grpc cli