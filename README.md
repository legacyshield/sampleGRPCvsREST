#Description
The project aims at comparing the difference in performance of REST and GRPC calls.

#Run
    - go run main_rest.go
    - cd grpc_server/
    - go run server.go
    - cd ../grpc_client
    - go test --bench Bench
