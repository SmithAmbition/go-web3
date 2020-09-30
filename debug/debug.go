package debug

import (
	"github.com/ethereum/go-web3/dto"
	"github.com/ethereum/go-web3/providers"
)

// Debug - The Debug Module
type Debug struct {
	provider providers.ProviderInterface
}

// NewDebug - Debug Module constructor to set the default provider
func NewDebug(provider providers.ProviderInterface) *Debug {
	debug := new(Debug)
	debug.provider = provider
	return debug
}

// TraceTransaction - Returns Trace Transaction info.
// Parameters:
//    - hash string

func (debug *Debug) TraceTransaction(txHash string) (*dto.RequestResult, error) {
	params := make([]string, 1)
	params[0] = txHash
	pointer := &dto.RequestResult{}

	err := debug.provider.SendRequest(pointer, "debug_traceTransaction", params)

	if err != nil {
		return nil, err
	}

	return pointer, nil

}
