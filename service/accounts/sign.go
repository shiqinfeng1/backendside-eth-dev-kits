package accounts

import (
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

//SignTx 交易签名
func SignTx(userID string, tx *cmn.TransactionRequest) ([]byte, error) {
	index := useridToIndex(userID)
	account, err := getAccountFromHDWallet(index)
	if err != nil {
		return nil, err
	}
	rawTx := types.NewTransaction(
		tx.Nonce.ToInt().Uint64(),
		tx.To,
		tx.Value.ToInt(),
		tx.Gas.ToInt().Uint64(),
		tx.GasPrice.ToInt(),
		tx.Data)

	//pretty.Print("account:", account, "rawTx:", rawTx)
	signedTx, err := wallet.SignTx(*account, rawTx, nil)
	var signedData bytes.Buffer
	signedTx.EncodeRLP(&signedData)
	//pretty.Print("signedData:", signedData.String())
	return signedData.Bytes(), nil
}
