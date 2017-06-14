# grpc-talk
Slides and demo code for my gRPC talk

run `protoc -I$GOPATH/src -I. demo.proto --go_out=plugins=grpc:$GOPATH/src` to generate Go output
