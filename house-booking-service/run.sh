goctl api go -api usercenter.api -dir ../
goctl rpc  protoc ./usercenter.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../
goctl model mysql datasource --url="root:admin@tcp(127.0.0.1:3306)/house" --table=user --dir ./user -c
