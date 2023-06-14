#!/bin/bash

until cqlsh -e "SHOW HOST" cassandra-db; do
    >&2 echo "Cassandra is unavailable - sleeping"
    sleep 2
done

>&2 echo "Cassandra is up - executing script"

cqlsh host.docker.internal

# Create keyspace
cqlsh host.docker.internal -e "CREATE KEYSPACE IF NOT EXISTS chatapp WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};"

# Use keyspace
cqlsh host.docker.internal -e "USE chatapp;"

# Create table
cqlsh host.docker.internal -e "CREATE TABLE IF NOT EXISTS messages (
  Chat TEXT,
  SendTime TIMESTAMP,
  Text TEXT,
  Sender TEXT,
  PRIMARY KEY(Chat, SendTime)
) WITH CLUSTERING ORDER BY (SendTime DESC);"

cqlsh host.docker.internal -e "
INSERT INTO chatapp.messages (Chat, SendTime, Text, Sender) VALUES ('jack:zack', dateof(now()), 'Hello world', 'jack');
INSERT INTO chatapp.messages (Chat, SendTime, Text, Sender) VALUES ('jack:zack', dateof(now()), 'BYE BYE world', 'zack');
INSERT INTO chatapp.messages (Chat, SendTime, Text, Sender) VALUES ('tom:jane', dateof(now()), 'hi', 'tom');
INSERT INTO chatapp.messages (Chat, SendTime, Text, Sender) VALUES ('tom:jane', dateof(now()), 'bye', 'jane');
"

echo "Keyspace and table creation completed"
