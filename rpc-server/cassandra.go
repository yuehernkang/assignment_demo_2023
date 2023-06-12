package main

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

type CassandraClient struct {
	session *gocql.Session
}

func (c *CassandraClient) Init() error {
	cluster := gocql.NewCluster("host.docker.internal")
	cluster.Keyspace = "messages"
	cluster.Consistency = gocql.Quorum
	sess, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	//TODO: remove this print statement
	c.session = sess
	return nil
}

func (c *CassandraClient) GetMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error) {
	query := "SELECT chat_id, message FROM messages WHERE chat_id =" + chatID
	iter := c.session.Query(query).Iter()
	scanner := iter.Scanner()
	var (
		messages []*Message
	)
	for scanner.Next() {
		var (
			chat_id string
			message string
		)
		temp := &Message{
			ChatID:  chat_id,
			Message: message,
		}

		messages = append(messages, temp)
	}
	fmt.Println(messages)

	return messages, nil
}

func (c *CassandraClient) TestGetMessage() error {
	query := "SELECT chat_id, message FROM messages WHERE chat_id = 'jack:zack'"
	iter := c.session.Query(query).Iter()
	scanner := iter.Scanner()
	for scanner.Next() {
		var (
			chat_id string
			message string
		)
		err := scanner.Scan(&chat_id, &message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tweet:", chat_id, message)
	}
	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	return nil
}
