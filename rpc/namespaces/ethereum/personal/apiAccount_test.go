package personal

import (
	"os"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/crypto/hd"
	"github.com/tendermint/tendermint/libs/log"
	// "github.com/evmos/evmos/v12/rpc/namespaces/ethereum/personal"
)

func TestCreateAcount(t *testing.T) {
	name := "key_" + time.Now().UTC().Format(time.RFC3339)
	lg := log.NewNopLogger()
	api := NewAPI(lg, nil)

	password := "dev01"
	// create the mnemonic and save the account
	hdPath := api.hdPathIter()

	info, err := api.backend.NewMnemonic(name, keyring.English, hdPath.String(), password, hd.EthSecp256k1)
	if err != nil {
		// return common.Address{}, err
		api.logger.Info("fail create address : ", common.Address{}, "err : ", err)
	}

	pubKey, err := info.GetPubKey()
	if err != nil {
		// return common.Address{}, err
		api.logger.Info("fail create pub address : ", common.Address{}, "err : ", err)
	}

	addr := common.BytesToAddress(pubKey.Address().Bytes())
	api.logger.Info("Your new key was generated", "address", addr.String())
	api.logger.Info("Please backup your key file!", "path", os.Getenv("HOME")+"/.evmos/"+name) // TODO: pass the correct binary
	api.logger.Info("Please remember your password!")

	// return addr, nil
}
