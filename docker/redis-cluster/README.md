```sh
docker exec -it redis-7001 redis-cli --cluster create redis-7001:7001 redis-7002:7002 redis-7003:7003 redis-7004:7004 redis-7005:7005 redis-7006:7006 --cluster-replicas 1
```
