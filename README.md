## Description

The project aims at performance comparison of REST and GRPC calls.

## Source Code
`go get github.com/ashu1994/sampleGRPCvsREST`

## Run

    1) go get ./...
    2) go run main_rest.go
    3) cd grpc_server/
    4) go run server.go
    5) cd ../grpc_client
    6) go test --bench Bench
