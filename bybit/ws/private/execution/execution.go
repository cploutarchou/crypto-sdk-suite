package execution

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Execution struct {
	*client.Client
}

func New(cli *client.Client) Execution {
	return Execution{cli}
}
