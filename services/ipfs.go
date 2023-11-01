package services

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sukrati192/fileserver/configmanager"
)

var IpfsShell *shell.Shell

func InitIPFS() {
	ipfsConfig := configmanager.GetConfig().IPFS
	IpfsShell = shell.NewShell(fmt.Sprintf("%s:%s", ipfsConfig.Host, ipfsConfig.Port))
}
