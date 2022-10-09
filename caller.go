package dexsdk

import (
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewCaller(link string) (*ethclient.Client, error) {
	var cli *ethclient.Client
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "wss" {
		cli, err = ethclient.Dial(u.String())
	} else {
		rpcCli, err := rpc.Dial(u.String())
		if err != nil {
			return nil, err
		}
		cli = ethclient.NewClient(rpcCli)
	}
	return cli, nil
}
