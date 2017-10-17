## Description

The project aims at comparing the difference in performance of REST and GRPC calls.

## Run

    1) go run main_rest.go
    2) cd grpc_server/
    3) go run server.go
    4) cd ../grpc_client
    5) go test --bench Bench
