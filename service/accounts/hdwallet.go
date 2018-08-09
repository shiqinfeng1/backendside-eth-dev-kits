package accounts

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/howeyc/gopass"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	hdwallet "github.com/shiqinfeng1/go-ethereum-hdwallet"
)

var wallet *hdwallet.Wallet
var ks *keystore.KeyStore

//NewRootHDWallet 新建hd钱包
func NewRootHDWallet() error {

	fmt.Printf("Input HDWallet Password: ")
	// Silent. For printing *'s use gopass.GetPasswdMasked()
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		return err
	}

	var mnemonic string
	p, _ := filepath.Abs(common.Config().GetString("hdwallet.stored"))
	_, err = os.Stat(p) //os.Stat获取文件信息
	if err == nil || (err != nil && os.IsExist(err)) {
		mnemonic = decryptoMnemonic(string(pass))
	} else {
		fmt.Printf("Input HDWallet Password comfirm: ")
		// Silent. For printing *'s use gopass.GetPasswdMasked()
		pass2, err := gopass.GetPasswdMasked()
		if err != nil {
			return err
		}
		if string(pass) != string(pass2) {
			return errors.New("Password Not The Same")
		}
		mnemonic, err = hdwallet.NewMnemonic(128)
		if err != nil {
			return err
		}
		cryptoAndSave(mnemonic, string(pass))
	}

	wallet, err = hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println("NewFromMnemonic fail:", err)
		return err
	}

	ks = keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	NewAccount("0")
	NewAccount("1")
	return nil
}

//GetRootHDWallet 获取hd根钱包
func getRootHDWallet() *hdwallet.Wallet {

	fmt.Printf("Input HDWallet Password: ")
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		// Handle gopass.ErrInterrupted or getch() read error
		return nil
	}
	if wallet != nil && wallet.CheckMnemonic(decryptoMnemonic(string(pass))) == true {
		return wallet
	}
	return nil
}

// NewAccount 创建新账户
func NewAccount(index string) error {
	if wallet == nil {
		return errors.New("no wallet")
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + index)
	account, err := wallet.Derive(path, true)
	if err != nil {
		return err
	}
	if ks.HasAddress(account.Address) {
		fmt.Println("exsit account:", account.Address.Hex())
		ks.Delete(account, "m44600"+index)
	}

	key, _ := wallet.PrivateKey(account)
	if _, err = ks.ImportECDSA(key, "m44600"+index); err != nil {
		return err
	}

	fmt.Println("new account:", account.Address.Hex())
	return nil
}
