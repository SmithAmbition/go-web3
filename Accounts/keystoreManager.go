package Accounts

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type KeystoreManager struct {
	Keystore *keystore.KeyStore
	ChainID  *big.Int
	Signer   types.Signer
}

func NewKeystoreManager(keystoreDir string, chainId int64) *KeystoreManager {
	ks := &KeystoreManager{Signer: types.NewEIP155Signer(big.NewInt(chainId))}
	ks.ChainID = big.NewInt(chainId)
	ks.Keystore = keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks
}
func (ks *KeystoreManager) getDecryptedKey(a accounts.Account, auth string) (accounts.Account, *keystore.Key, error) {
	a, err := ks.Keystore.Find(a)
	if err != nil {
		return a, nil, err
	}
	key, err := ks.GetKey(a.Address, a.URL.Path, auth)
	return a, key, err
}
func (ks *KeystoreManager) Unlock(addr common.Address, password string) error {
	return ks.TimedUnlock(addr, password, 0)
}
func (ks *KeystoreManager) TimedUnlock(addr common.Address, password string, timeout time.Duration) error {
	acc := accounts.Account{Address: addr}
	return ks.Keystore.TimedUnlock(acc, password, timeout)
}
func (ks *KeystoreManager) GetKey(addr common.Address, filename, auth string) (*keystore.Key, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	// Make sure we're really operating on the requested key (no swap attacks)
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, addr)
	}
	return key, nil
}

func (ks *KeystoreManager) SignTx(sendTx *types.Transaction, from common.Address) ([]byte, error) {
	acc := accounts.Account{Address: from}
	tx, err := ks.Keystore.SignTx(acc, sendTx, ks.ChainID)
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(tx)
}

func (ks *KeystoreManager) SignTxWithPassphrase(sendTx *types.Transaction, from common.Address, passphrase string) ([]byte, error) {
	acc := accounts.Account{Address: from}
	tx, err := ks.Keystore.SignTxWithPassphrase(acc, passphrase, sendTx, ks.ChainID)
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(tx)
}
