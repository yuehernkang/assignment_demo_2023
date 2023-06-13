package main

import (
	"context"
	"fmt"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	resp := rpc.NewSendResponse()
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	roomID, err := getRoomID(req.GetChat())
	if err != nil {
		return nil, err
	}
	messages, err := cassandraClient.GetMessagesByChatID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	respMessages := make([]*rpc.Message, 0)

	for _, msg := range messages {
		temp := &rpc.Message{
			Chat:     msg.Chat,
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

func getRoomID(chat string) (string, error) {
	roomID := chat
	return roomID, nil
}
