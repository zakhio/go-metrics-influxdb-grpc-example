go-metrics-influxdb-grpc-example
================================

Example of GRPC service with go-metrics exported to local InfluxDB.

Installation
------------

```shell
./generate-proto.sh

docker-compose up -d
go run main.go
go run cmd/cli.go
```

License
-------

go-metrics-influxdb-grpc-example is licensed under the MIT license. See the LICENSE file for details.
