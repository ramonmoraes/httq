# httq
HTTQ serves a "queue" over a rest api

Following those endpoints

```
{infra}/{key}/ [POST] --data
{infra}/{key}/ [GET]
```
Where:

__infra__ relates to the chosen system that will handle the message, may it be a kafka broker as `kafka`, a array in ram as `ram` or a line file as `disk`

Currently suported: `kafka`

__key__ is the argument that correctly routes the message to the queue/topic

Example:

Doing a post http request to `kafka/sometopic/` will write a event to the kafka's topic `sometopic` passing it's body as the event payload

```sh
$ curl --request POST \
  --url http://localhost:8000/kafka/sometopic \
  --header 'content-type: application/json' \
  --data '{"body": "payload"}'

# returns: sometopic[0]@1

$ curl --request GET \
  --url http://localhost:8000/kafka/sometopic

# returns: {"body": "payload"}
```

## Running

### Kafka

```
docker-compose up kafka zookeeper
go run main.go
```

Obs: It may be improved by pre-creating the consumers and producers and re-utilizem them instead of creating a new one each requests

But it would change the dinamically behavior of `path`