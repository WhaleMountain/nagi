package dk

import (
	"context"
	"github.com/docker/docker/client"
)

var (
	ctx context.Context
	cli *client.Client
	err error
)

func Init() error {
    ctx = context.Background()
    cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
	}
	return nil
}

func GetDK() (*client.Client, context.Context) {
	return cli, ctx
}