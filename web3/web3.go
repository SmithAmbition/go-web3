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

package web3

import (
	"github.com/ethereum/go-web3/dto"
	"github.com/ethereum/go-web3/providers"
	"github.com/ethereum/go-web3/utils"
	"math/big"
)

// Web3 - The Web3 Module
type Web3 struct {
	provider providers.ProviderInterface
}

// NewWeb3 - Web3 Module constructor to set the default provider
func NewWeb3(provider providers.ProviderInterface) *Web3 {
	web := new(Web3)
	web.provider = provider
	return web
}

// GetBlockByNumber - Returns the information about a block requested by number.
// Parameters:
//    - number, QUANTITY - number of block
//    - transactionDetails, bool - indicate if we should have or not the details of the transactions of the block
// Returns:
//    1. Object - A block object, or null when no transaction was found
//    2. error
func (web *Web3) GetBlockByNumber(number *big.Int, transactionDetails bool) (*dto.Block, error) {

	params := make([]interface{}, 2)
	params[0] = utils.IntToHex(number)
	params[1] = transactionDetails

	pointer := &dto.RequestResult{}

	err := web.provider.SendRequest(pointer, "man_getBlockByNumber", params)

	if err != nil {
		return nil, err
	}
	return pointer.ToBlock();
}

// GetBlockNumber - Returns the number of most recent block.
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current block number the client is on.
func (web *Web3) GetBlockNumber() (*big.Int, error) {

	pointer := &dto.RequestResult{}

	err := web.provider.SendRequest(pointer, "man_blockNumber", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetTransactionCount -  Returns the number of transactions sent from an address.
// Parameters:
//    - DATA, 20 Bytes - address to check for balance.
// Returns:
// 	  - QUANTITY - integer of the number of transactions sent from this address
func (web *Web3) GetTransactionCount(address string, defaultBlockParameter string) (*big.Int, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := web.provider.SendRequest(pointer, "man_getTransactionCount", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

func (web *Web3) SendRawTransaction(rawTx interface{}) (string, error) {

	params := make([]interface{}, 1)
	params[0] = rawTx

	pointer := &dto.RequestResult{}

	err := web.provider.SendRequest(&pointer, "man_sendRawTransaction", params)

	if err != nil {
		return "", err
	}
	return pointer.ToString()
}

func (web *Web3) SendStringRawTransaction(rawTx string) (string, error) {
	params := make([]interface{}, 1)
	params[0] = rawTx

	pointer := &dto.RequestResult{}

	err := web.provider.SendRequest(&pointer, "man_sendRawTransaction", params)

	if err != nil {
		return "", err
	}
	return pointer.ToString()
}

// GetTransactionReceipt - Returns compiled solidity code.
// Parameters:
//    1. DATA, 32 Bytes - hash of a transaction.
// Returns:
//	  1. Object - A transaction receipt object, or null when no receipt was found:
//    - transactionHash: 		DATA, 32 Bytes - hash of the transaction.
//    - transactionIndex: 		QUANTITY - integer of the transactions index position in the block.
//    - blockHash: 				DATA, 32 Bytes - hash of the block where this transaction was in.
//    - blockNumber:			QUANTITY - block number where this transaction was in.
//    - cumulativeGasUsed: 		QUANTITY - The total amount of gas used when this transaction was executed in the block.
//    - gasUsed: 				QUANTITY - The amount of gas used by this specific transaction alone.
//    - contractAddress: 		DATA, 20 Bytes - The contract address created, if the transaction was a contract creation, otherwise null.
//    - logs: 					Array - Array of log objects, which this transaction generated.
func (web *Web3) GetTransactionReceipt(hash string) (*dto.TransactionReceipt, error) {

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult{}

	err := web.provider.SendRequest(pointer, "man_getTransactionReceipt", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToTransactionReceipt()

}
/*
func (web *Web3) SignTxByPrivate(sendTX *common.SendTxArgs1, from string,Privatekey *ecdsa.PrivateKey,ChainId *big.Int)(*common.SendTxArgs1,error) {
	tx1 := sendTX.ToTransaction()
	tx,err:=types.SignTx(tx1, types.NewEIP155Signer(ChainId),Privatekey)
	if err!=nil {
		return sendTX,err
	}

	sendTX.R = (*hexutil.Big)(tx.GetTxR())
	sendTX.S = (*hexutil.Big)(tx.GetTxS())
	sendTX.V = (*hexutil.Big)(tx.GetTxV())
	return sendTX,nil
}
*/
func (web *Web3) GetGasPrice() (*big.Int, error) {
	pointer := &dto.RequestResult{}
	err := web.provider.SendRequest(pointer, "man_getGasPrice", nil)
	if err != nil {
		return nil, err
	}
	price ,err := pointer.ToBigInt()
	return price, err
}
