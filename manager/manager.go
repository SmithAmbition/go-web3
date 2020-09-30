package manager

import (
	"github.com/ethereum/go-web3/web3Manager"
	"github.com/ethereum/go-web3/Accounts"
)

//var (
//	KeystorePath = "keystore"
//	Tom_Manager  = &Manager{
//		AIMan.NewAIMan(providers.NewHTTPProvider("api85.matrix.io", 100, false)),
//		Accounts.NewKeystoreManager(KeystorePath, 1),
//	}
//)
//
type Manager struct {
	*web3Manager.Web3Manager
	*Accounts.KeystoreManager
}
