s1: docker-compose up -d 
=> zookeeper & kakfa running

s2: 
docker exec -it kafka kafka-topics \
  --create --topic topic-1 \
  --bootstrap-server localhost:9092 \
  --partitions 3 --replication-factor 1

docker exec -it kafka kafka-topics \
  --create --topic topic-2 \
  --bootstrap-server localhost:9092 \
  --partitions 3 --replication-factor 1

=> 2 topics created

s3: go run consumer.go
=>📡 Connecting to Kafka broker at:  localhost:9092
 and subscribing to topic:  topic-1
📥 Listening for messages on topic: topic-1

s4: go run producer.go
=> 📡 Connecting to Kafka broker at:  localhost:9092
 and publishing to topic:  topic-1
