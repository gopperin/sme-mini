# sme-mini

general data platform, use lift as stream

## Compile Proto files

Run the command below from the types directory:

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb

## krakend
