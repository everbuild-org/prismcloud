package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "prismcloud.dev/protobufs"
)

type Client struct {
	conn          *grpc.ClientConn
	contextCancel context.CancelFunc
	Api           pb.PrismcloudApiserverClient
	Ctx           context.Context
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		panic(err)
	}
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewPrismcloudApiserverClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return &Client{
		conn:          conn,
		Api:           client,
		contextCancel: cancel,
		Ctx:           ctx,
	}, nil
}
