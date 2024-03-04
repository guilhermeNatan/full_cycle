comando para gerar as interfaces de comunicação do grpc
https://github.com/ktr0731/evans?tab=readme-ov-file#not-recommended-go-install

protoc --go_out=. --go-grpc_out=.  internal/proto/course_category.proto 


go run cmd/grpcServer/main.go 


evans -r repl


package pb 
service CategoryService 




