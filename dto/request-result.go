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

/**
 * @file request-result.go
 * @authors:
 *   Reginaldo Costa <ethereum@gmail.com>
 * @date 2017
 */

package dto

import (
	"errors"
	"strconv"

	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-web3/complex/types"
	customerror "github.com/ethereum/go-web3/constants"
)

type RequestResult struct {
	ID      int             `json:"id"`
	Version string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error,omitempty"`
	Data    string          `json:"data,omitempty"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (pointer *RequestResult) UnMarshalResult(result interface{}) error {
	if err := pointer.checkResponse(); err != nil {
		return err
	}
	return json.Unmarshal(pointer.Result, result)
}
func (pointer *RequestResult) ToBigInt() (*big.Int, error) {
	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}
	value := new(hexutil.Big)
	err := json.Unmarshal(pointer.Result, &value)
	return value.ToInt(), err
}
func (pointer *RequestResult) ToStringArray() ([]string, error) {

	new := make([]string, 0)
	err := pointer.UnMarshalResult(&new)
	return new, err

}

func (pointer *RequestResult) ToComplexString() (types.ComplexString, error) {

	result, err := pointer.ToString()
	if err != nil {
		return "", err
	}
	return types.ComplexString(result), nil

}

func (pointer *RequestResult) ToString() (string, error) {
	var new string
	err := pointer.UnMarshalResult(&new)
	return new, err

}

func (pointer *RequestResult) ToInt() (int64, error) {

	result, err := pointer.ToString()
	if err != nil {
		return 0, err
	}
	numericResult, err := strconv.ParseInt(result, 16, 64)

	return numericResult, err

}

/*

func (pointer *RequestResult) ToComplexIntResponse() (types.ComplexIntResponse, error) {

	if err := pointer.checkResponse(); err != nil {
		return types.ComplexIntResponse(0), err
	}

	result := (pointer).Result.(interface{})

	var hex string

	switch v := result.(type) {
	//Testrpc returns a float64
	case float64:
		hex = strconv.FormatFloat(v, 'E', 16, 64)
		break
	default:
		hex = result.(string)
	}

	cleaned := strings.TrimPrefix(hex, "0x")

	return types.ComplexIntResponse(cleaned), nil

}
*/
func (pointer *RequestResult) ToBoolean() (bool, error) {

	var result bool
	err := pointer.UnMarshalResult(&result)

	return result, err

}

func (pointer *RequestResult) ToSignTransactionResponse() (*SignTransactionResponse, error) {

	signTransactionResponse := &SignTransactionResponse{}
	err := pointer.UnMarshalResult(signTransactionResponse)

	return signTransactionResponse, err
}

func (pointer *RequestResult) ToTransactionResponse() (*TransactionResponse, error) {

	transactionResponse := &TransactionResponse{}

	err := pointer.UnMarshalResult(transactionResponse)
	return transactionResponse, err

}

func (pointer *RequestResult) ToTransactionReceipt() (*TransactionReceipt, error) {

	transactionReceipt := &TransactionReceipt{}
	err := pointer.UnMarshalResult(transactionReceipt)
	return transactionReceipt, err

}

func (pointer *RequestResult) ToBlock() (*Block, error) {
	block := &Block{}
	err := pointer.UnMarshalResult(block)
	return block, err

}

func (pointer *RequestResult) ToSyncingResponse() (*SyncingResponse, error) {

	syncingResponse := &SyncingResponse{}

	pointer.UnMarshalResult(syncingResponse)

	return syncingResponse, nil

}

// To avoid a conversion of a nil interface
func (pointer *RequestResult) checkResponse() error {

	if pointer.Error != nil {
		return errors.New(pointer.Error.Message)
	}

	if pointer.Result == nil {
		return customerror.EMPTYRESPONSE
	}

	return nil

}
