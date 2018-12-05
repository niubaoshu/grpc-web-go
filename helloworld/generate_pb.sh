#!/usr/bin/env bash

protoc -I=. \
   --js_out=import_style=commonjs:. \
   --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. \
   --go_out=plugins=grpc:. *.proto
