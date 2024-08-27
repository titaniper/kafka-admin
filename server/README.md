https://kafka.apache.org/32/javadoc/org/apache/kafka/clients/admin/Admin.html

```
go get github.com/ricardo-ch/go-kafka-connect
go get github.com/pkg/errors
go get github.com/ricardo-ch/go-kafka-connect/lib/connectors@v0.0.0-20200403115642-f7b66cb04ed7
go get gopkg.in/resty.v1

```

```
go test ./services/connectors
go test -run Test_GET_CONNECTOR
```


```
go install github.com/swaggo/swag/cmd/swag@latest
swag init
swag init -g cmd/main.go
```
