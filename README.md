# Installation guide

Installation uses Confluent Platform
[Quick Start for Confluent Platform](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html#ce-docker-quickstart)
[Docker Image(s)](https://hub.docker.com/u/confluentinc)
[Docker Setup Example](https://github.com/confluentinc/cp-all-in-one/blob/7.0.1-post/cp-all-in-one/docker-compose.yml)

## Start services

```
docker compose up -d
```

check out control-center available at http://localhost:9021/

## List topics

```
docker exec -it broker kafka-topics --bootstrap-server localhost:9092 --list
```

## Create topic

```
docker exec -it broker kafka-topics --create --topic my.topic --replication-factor 1 --partitions 2  --bootstrap-server localhost:9092
```

## Send and receive messages

```
docker exec -it broker kafka-console-producer --bootstrap-server localhost:9092 --topic my.topic
docker exec -it broker kafka-console-consumer --bootstrap-server localhost:9092 --topic my.topic
```
