package nagi_con

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"

	"nagi-docker/dk"
	"nagi-docker/models"
	"nagi-docker/services/nagi_net"
	"github.com/gin-gonic/gin"
)

type Service_container struct{}

func setConfig(image string, domain string, environment []string, gport string) *container.Config {
	config := &container.Config{
		Image: image,
		Domainname: domain,
		Env: environment,
		ExposedPorts: nat.PortSet{nat.Port(gport): struct{}{}},
	}

	return config
}

func setHostConfig(hport string, gport string) *container.HostConfig {
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(gport): []nat.PortBinding{{HostPort: hport}},
		},
	}

	return hostConfig
}

func setNetConfig(netID string) *network.NetworkingConfig {
	netConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			netID: {
				IPAMConfig: &network.EndpointIPAMConfig{},
			},
		},
	}

	return netConfig
}

func (sc Service_container) CreateNewCompose(c *gin.Context) error {
	cli, ctx := dk.GetDK()

	var con models.Compose
	if err := c.BindJSON(&con); err != nil {
		return err
	}

	for idx := 0; idx < len(con.Images); idx++ {
		config := setConfig(con.Images[idx], con.ConName[idx], con.Environment[idx], con.GuestPorts[idx])
		hostConfig := setHostConfig(con.HostPorts[idx], con.GuestPorts[idx])

		netID, err := nagi_net.CreateNewNetwork(con.ComposeName, con.Driver)
		if err != nil {
			return err
		}
	
		netConfig := setNetConfig(netID)

		// コンテナ作成
		conResp, err := cli.ContainerCreate(ctx, config, hostConfig, netConfig, con.ConName[idx])
		if err != nil {
			return err
		}
		// コンテナ起動
		if err := cli.ContainerStart(ctx, conResp.ID, types.ContainerStartOptions{}); err != nil {
			return err
		}
	}

	return nil
}

func (sc Service_container) CreateNewContainer(c *gin.Context) error {
	cli, ctx := dk.GetDK()

	var con models.Container
	if err := c.BindJSON(&con); err != nil {
		return err
	}

	config := setConfig(con.Image, con.ConName, con.Environment, con.GuestPort)
	hostConfig := setHostConfig(con.HostPort, con.GuestPort)

	netID, err := nagi_net.CreateNewNetwork(con.ConName, con.Driver)
	if err != nil {
		return err
	}


	netConfig := setNetConfig(netID)

	// コンテナ作成
	conResp, err := cli.ContainerCreate(ctx, config, hostConfig, netConfig, con.ConName)
	if err != nil {
		return err
	}
	// コンテナ起動
	if err := cli.ContainerStart(ctx, conResp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}