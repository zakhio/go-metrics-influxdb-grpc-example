go-metrics-influxdb-grpc-example
================================

Example of GRPC service with go-metrics exported to local InfluxDB instance running in Docker.

Uses:
* https://github.com/go-yaml/yaml
* https://github.com/grpc/grpc-go
* https://github.com/protocolbuffers/protobuf
* https://github.com/rcrowley/go-metrics
* https://github.com/zakhio/go-metrics-influxdb

Usage
-----

```shell
./generate-proto.sh
docker-compose up -d
go run server/main.go
go run client/main.go
```

License
-------

go-metrics-influxdb-grpc-example is licensed under the MIT license. See the LICENSE file for details.
