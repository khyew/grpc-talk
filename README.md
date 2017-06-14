# grpc-talk
Slides and demo code for my gRPC talk

run `protoc -I$GOPATH/src -I. demo.proto --go_out=plugins=grpc:$GOPATH/src` to generate Go output

run `client_csharp/packages/Grpc.Tools.1.3.6/tools/windows_x64/protoc.exe -I. --csharp_out . --grpc_out . demo.proto --plugin=protoc-gen-grpc=client_csharp/packages/Grpc.Tools.1.3.6/tools/windows_x64/grpc_csharp_plugin.exe` to generate C# output
