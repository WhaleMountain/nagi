package nagi_net

import (
	"github.com/docker/docker/api/types"

	"nagi-docker/dk"
)

func checkNetwork(netName string) (string, error) {
	cli, ctx := dk.GetDK()

	netResps, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		return "", err
	}

	for _, netResp := range netResps {
		if netResp.Name == netName {
			return netResp.ID, nil
		}
	}

	return "", nil
}

func CreateNewNetwork(netName string, driver string) (string, error) {
	cli, ctx := dk.GetDK()

	if netID, err := checkNetwork(netName+"_"+driver); netID != "" {
		return netID, nil

	} else if err != nil {
		return "", err

	}
	
	netResp, err := cli.NetworkCreate(ctx, netName+"_"+driver, types.NetworkCreate{Driver: driver})
	if err != nil {
		return "", err
	}

	return netResp.ID, nil
}

func DeleteNetWork(netID string) error {
	cli, ctx := dk.GetDK()
	err := cli.NetworkRemove(ctx, netID)
	
	return err
}