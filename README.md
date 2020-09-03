# httq
Http serves a queue over a rest api

## Running

### Kafka

```
docker-compose up kafka zookeeper
go run main.go
```

Obs: It may be improved by pre-creating the consumers and producers and re-utilizem them instead of creating a new one each requests

But it would change the dinamically behavior of `path`