# sme-mini

es-scaffold, use lift as stream

## Compile Proto files

Run the command below from the types directory:

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb

## sonar

    sonar-scanner -Dsonar.projectKey=sme-mini -Dsonar.sources=. -Dsonar.host.url=http://127.0.0.1:9001 -Dsonar.login=065fbe7f76bd31892a40c2e5a7ef278365eb1522
