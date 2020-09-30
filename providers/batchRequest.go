package providers

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/ethereum/go-web3/providers/util"
)

type BatchRequest struct {
	RpcMethods []util.JSONRPCObject
}

func (br *BatchRequest) Length() int {
	return len(br.RpcMethods)
}
func (br *BatchRequest) AddRequest(method string, params interface{}) {
	br.RpcMethods = append(br.RpcMethods, util.JSONRPCObject{Version: "2.0", Method: method, Params: params, ID: rand.Intn(5000)})
}
func (br *BatchRequest) AsJsonString() string {
	resultBytes, err := json.Marshal(br.RpcMethods)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(resultBytes)
}
