proto command := protoc -Igreet/proto --go_out=. --go_opt=module=github.com/Amartya-Bhardwaj/grpc  --go-grpc_out=. --go-grpc_opt=module=github.com/Amartya-Bhardwaj/grpc  greet/proto/greet.proto