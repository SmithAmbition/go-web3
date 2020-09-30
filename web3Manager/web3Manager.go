/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

package web3Manager

import (
	"strings"

	"github.com/ethereum/go-web3/debug"
	"github.com/ethereum/go-web3/dto"
	"github.com/ethereum/go-web3/eth"

	//"github.com/ethereum/go-web3/net"
	//"github.com/ethereum/go-web3/personal"
	"github.com/ethereum/go-web3/providers"
	//"github.com/ethereum/go-web3/utils"
)

const (
	Coin float64 = 1000000000000000000
)

// web3 - The web3 Module
type Web3Manager struct {
	Provider providers.ProviderInterface
	Eth      *eth.Eth
	Debug    *debug.Debug
	//	Web3      *web3.Web3
	//Net      *net.Net
	//Personal *personal.Personal
	//Utils    *utils.Utils
}

func NewHttpWeb3Manager(rpcUrl string) *Web3Manager {
	secure := false
	rpc := rpcUrl
	if strings.Index(rpcUrl, "http://") == 0 {
		rpc = rpc[len("http://"):len(rpcUrl)]
	} else if strings.Index(rpcUrl, "https://") == 0 {
		rpc = rpc[len("https://"):len(rpcUrl)]
		secure = true
	}
	return NewWeb3Manager(providers.NewHTTPProvider(rpc, 100, secure))
}

// NewAIMan - AIMan Module constructor to set the default provider, Man, Net and Personal
func NewWeb3Manager(provider providers.ProviderInterface) *Web3Manager {
	web3Manager := new(Web3Manager)
	web3Manager.Provider = provider
	web3Manager.Eth = eth.NewEth(provider)
	web3Manager.Debug = debug.NewDebug(provider)
	//aiMan.Net = net.NewNet(provider)
	//aiMan.Personal = personal.NewPersonal(provider)
	//aiMan.Utils = utils.NewUtils(provider)
	return web3Manager
}

// ClientVersion - Returns the current client version.
// Parameters:
//    - none
// Returns:
// 	  - String - The current client version
func (web3 Web3Manager) ClientVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := web3.Provider.SendRequest(pointer, "web3_clientVersion", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
