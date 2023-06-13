package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	resp := rpc.NewSendResponse()
	resp.Code, resp.Msg = areYouLucky()
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	messages, err := cassandraClient.GetMessagesByChatID(ctx, "jack:zack")
	if err != nil {
		return nil, err
	}

	respMessages := make([]*rpc.Message, 0)

	//TODO: Change all to sender
	for _, msg := range messages {
		temp := &rpc.Message{
			Chat:     req.GetChat(),
			Text:     msg.Text,
			Sender:   msg.Sender,
			SendTime: msg.SendTime,
		}
		respMessages = append(respMessages, temp)
	}
	resp := rpc.NewPullResponse()
	resp.Messages = respMessages
	resp.Code = 0
	resp.Msg = "success"
	fmt.Println(resp)
	return resp, nil
}

func areYouLucky() (int32, string) {
	if rand.Int31n(2) == 1 {
		return 0, "success"
	} else {
		return 500, "oops"
	}
}
