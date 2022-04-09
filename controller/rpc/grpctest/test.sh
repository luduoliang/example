#! /bin/sh
rm -rf *.pb.go
protoc --go_out=plugins=grpc:. *.proto