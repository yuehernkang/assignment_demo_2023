CREATE KEYSPACE chatapp
    WITH replication = {'class': 'NetworkTopologyStrategy'}
     AND durable_writes = true;

USE chatapp;

CREATE TABLE chatapp.messages (
                                  Chat TEXT,
                                  SendTime TIMESTAMP,
                                  Text TEXT,
                                  Sender TEXT,
                                  PRIMARY KEY(Chat, SendTime)
) WITH CLUSTERING ORDER BY (SendTime DESC);

INSERT INTO ChatApp.messages (Chat, SendTime, Text, Sender) VALUES ('jack:zack', dateof(now()), 'Hello world', 'jack');
INSERT INTO ChatApp.messages (Chat, SendTime, Text, Sender) VALUES ('jack:zack', dateof(now()), 'BYE BYE world', 'zack');
INSERT INTO ChatApp.messages (Chat, SendTime, Text, Sender) VALUES ('tom:jane', dateof(now()), 'hi', 'tom');
INSERT INTO ChatApp.messages (Chat, SendTime, Text, Sender) VALUES ('tom:jane', dateof(now()), 'bye', 'jane');
