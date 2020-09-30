package RPCParam
import (
	"github.com/ethereum/go-web3/interfaces"
	import "errors"
)
type InputParam struct {
	encoder map[string] interfaces.encoder
}
func (param *InputParam) RPCInputParam(command string,param ...interface{}) ([]interface{},error) {
	if encoder,exist := param.encoder[command];exist{
		return encoder.EncodeParam(param);
	}else{
		return nil,errors.new(command + " command is not exist");
	}
}
func (param *InputParam)RPCOutputParam(){
	
}