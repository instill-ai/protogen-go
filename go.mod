module github.com/instill-ai/protogen-go

go 1.21
toolchain go1.24.1

retract v0.3.0-alpha

require (
	cloud.google.com/go/longrunning v0.4.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.56.3
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
