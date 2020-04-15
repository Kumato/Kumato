package controller

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/kumato/kumato/internal/auth"
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Controller struct {
	types.UnimplementedControllerServer
}

func registerConn(id, addr string) (types.WorkerClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithPerRPCCredentials(auth.JWTClaims{}))

	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	saveConfig(id, addr)

	return types.NewWorkerClient(conn), nil
}

func (c *Controller) Register(ctx context.Context, in *types.Node) (*types.Empty, error) {
	client, err := registerConn(in.GetId(), in.GetAddress())
	if err != nil {
		logger.Fatal("error appeared when dialing gRPC to", in.GetAddress())
		return &types.Empty{}, err
	}

	nodes.Store(in.GetId(), client)
	go writeConfig()

	logger.Warn("node", in.GetId(), "from", in.GetAddress(), "registered")

	go AssignTask()

	return &types.Empty{}, nil
}

func (c *Controller) TaskDone(ctx context.Context, in *types.Task) (*types.Task, error) {
	var t types.Task
	t.ContainerId = in.GetContainerId()
	t.Node = in.GetNode()
	t.Id = in.GetId()
	db.GetTask(&t)

	if t.Id == 0 {
		e := "task " + t.GetContainerId() + " with container id " + t.GetContainerId() + " has not been found on " + t.GetNode()
		logger.Fatal(e)
		return &t, errors.New(e)
	}

	t.FinishTime = in.FinishTime
	t.ExitCode = in.ExitCode

	db.UpdateTask(&t)

	logger.Warn("task", t.GetId(), "has been marked as finished")

	go AssignTask()

	return &t, nil
}
