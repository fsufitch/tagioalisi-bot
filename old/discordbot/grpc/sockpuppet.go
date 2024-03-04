package grpc

import (
	"context"
	"fmt"

	"github.com/fsufitch/tagioalisi-bot/bot/sockpuppet-module"
	"github.com/fsufitch/tagioalisi-bot/log"
	pb "github.com/fsufitch/tagioalisi-bot/proto"
	"github.com/fsufitch/tagioalisi-bot/security"
)

type SockpuppetServer struct {
	pb.UnimplementedGreeterServer
	Log              *log.Logger
	SockpuppetModule *sockpuppet.Module
	JWT              *security.JWTSupport
}

func (s SockpuppetServer) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.SendMessageReply, error) {
	jwtData, err := s.JWT.ExtractJWTData(request.Jwt)
	if err != nil {
		return nil, fmt.Errorf("failed to extract jwt: %v; %w", request.Jwt, err)
	}

	if err := s.SockpuppetModule.SendMessage(request.ChannelID, request.Content, jwtData.UserID); err != nil {
		return &pb.SendMessageReply{Status: &pb.UnaryStatus{Ok: false, Message: err.Error()}}, nil
	}

	return &pb.SendMessageReply{Status: &pb.UnaryStatus{Ok: true, Message: ""}}, nil
}
