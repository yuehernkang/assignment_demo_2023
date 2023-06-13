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
	query := "SELECT Chat, Text FROM messages.Messages WHERE Chat =" + chatID
	iter := c.session.Query(query).Iter()
	scanner := iter.Scanner()
	var (
		messages []*Message
	)
	for scanner.Next() {
		var (
			Chat string
			Text string
		)
		err := scanner.Scan(&Chat, &Text)
		if err != nil {
			log.Fatal(err)
			return messages, err
		}

		temp := &Message{
			Chat: Chat,
			Text: Text,
		}

		messages = append(messages, temp)
	}
	fmt.Println(messages)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return messages, err
	}
	return messages, nil
}

func (c *CassandraClient) TestGetMessage() ([]*Message, error) {
	query := "SELECT Chat, Text FROM messages.Messages WHERE Chat = 'jack:zack'"
	iter := c.session.Query(query).Iter()
	scanner := iter.Scanner()
	var (
		messages []*Message
	)
	for scanner.Next() {
		var (
			Chat string
			Text string
		)
		err := scanner.Scan(&Chat, &Text)
		if err != nil {
			log.Fatal(err)
			return messages, err
		}

		temp := &Message{
			Chat: Chat,
			Text: Text,
		}

		messages = append(messages, temp)
	}
	fmt.Println(messages)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return messages, err
	}
	return messages, nil
}
