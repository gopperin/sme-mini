# gudp

general user data platform

## Compile Proto files

Run the command below from the types directory:

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb
