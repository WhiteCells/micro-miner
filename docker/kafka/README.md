

```sh
docker exec -it kafka1 kafka-console-producer.sh --broker-list localhost:9092,localhost:9093 --topic test-topic
```

```sh
docker exec -it kafka1 kafka-console-consumer.sh --bootstrap-server localhost:9092,localhost:9093 --topic test-topic --from-beginning
```