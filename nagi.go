package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/client"
)

func NewNetwork(cli *client.Client, ctx context.Context, netName string, driver string) (types.NetworkCreateResponse, error) {
	netResp, err := cli.NetworkCreate(ctx, netName, types.NetworkCreate{Driver: driver})
	if err != nil {
		return netResp, err
	}
	return netResp, err
}

func NewContainer(cli *client.Client, ctx context.Context, conName []string, images []string, environment [][]string, guestPorts []string, hostPort []string, netResp types.NetworkCreateResponse) {
	for idx := 0; idx < len(images); idx++ {
		config := &container.Config{
			Image: images[idx],
			Domainname: conName[idx],
			Env: environment[idx],
			ExposedPorts: nat.PortSet{nat.Port(guestPorts[idx]): struct{}{}},
		}
		hostConfig := &container.HostConfig{
			PortBindings: nat.PortMap{
				nat.Port(guestPorts[idx]): []nat.PortBinding{{HostPort: hostPort[idx]}},
			},
		}
		netConfig := &network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				netResp.ID: {
					IPAMConfig: &network.EndpointIPAMConfig{},
				},
			},
		}
		// コンテナ作成
		conResp, err := cli.ContainerCreate(ctx, config, hostConfig, netConfig, conName[idx])
		if err != nil {
			panic(err)
		}
		// コンテナ起動
		if err := cli.ContainerStart(ctx, conResp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}
	}
	//return err
}

func main() {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
	}

	netResp, err := NewNetwork(cli, ctx, "nagi_bridge", "bridge")
	if err != nil {
		panic(err)
	}
	
	environment := [][]string{
		{"-e", "MYSQL_ROOT_PASSWORD=mysql"},
		{"-e", "WORDPRESS_DB_HOST=nagi-db", "WORDPRESS_DB_USER=root", "WORDPRESS_DB_PASSWORD=mysql"},
	}

	NewContainer(cli, ctx, []string{"nagi-db", "nagi-word"},[]string{"mysql:5.7", "wordpress:latest"}, environment, []string{"3306", "80"}, []string{"3306", "8080"}, netResp)
}