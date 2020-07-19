go-metrics-influxdb-grpc-example
================================

Example of [GRPC](https://github.com/grpc/grpc-go) echo service with [go-metrics](https://github.com/rcrowley/go-metrics) 
exported to [InfluxDB](https://influxdb.com/) instance running in [Docker](https://www.docker.com/) by 
[go-metrics-influxdb](https://github.com/zakhio/go-metrics-influxdb) reporter.

Prerequisites:
* Protocol Buffer Compiler (see [protoc installation guide](https://grpc.io/docs/protoc-installation/)).
* Docker Compose (see [docker-compose installation guide](https://docs.docker.com/compose/install/)).

Usage
-----

Generate golang code from `echo.proto` definition:
```shell
./generate-proto.sh
```

Start docker-compose with `InfluxDB` and `Chronograf`:
```shell
docker-compose up -d
```

The configuration for InfluxDB reporter is in `influxdb-config.yaml` file. Don't need to create `localBucketId` database, it will be done automatically.

Server and client available by the next commands: 
```shell
go run server/main.go
go run client/main.go
```

Run them in different terminals.

Echo Calls Rate Graph  
-----------------------------
After `docker-compose up -d` is executed, then `Chronograf` instance is available on `http://localhost:8888`.

If you would like to see the rate for echo calls, then follow next link [http://localhost:8888/sources/0/chronograf/data-explorer](http://localhost:8888/sources/0/chronograf/data-explorer?query=SELECT%20stddev%28%22echo.count%22%29%20AS%20%22stddev_echo.count%22%20FROM%20%22localBucketId%22.%22autogen%22.%22measurement%22%20WHERE%20time%20%3E%20%3AdashboardTime%3A%20AND%20time%20%3C%20%3AupperDashboardTime%3A%20GROUP%20BY%20time%281m%29%20FILL%28null%29). 
Or use the query:

```
SELECT stddev("echo.count") AS "stddev_echo.count" FROM "localBucketId"."autogen"."measurement" WHERE time > :dashboardTime: AND time < :upperDashboardTime: GROUP BY time(1m) FILL(null)
```

License
-------

go-metrics-influxdb-grpc-example is licensed under the MIT license. See the LICENSE file for details.
