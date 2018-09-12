package accounts

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/howeyc/gopass"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
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
	p, _ := filepath.Abs(cmn.Config().GetString("hdwallet.stored"))
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
		cmn.Logger.Error("NewFromMnemonic fail:", err)
		return err
	}

	ks = keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	return nil
}

//getRootHDWallet 获取hd根钱包
func getRootHDWallet() *hdwallet.Wallet {

	fmt.Printf("Input HDWallet Password: ")
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		return nil
	}
	if wallet != nil && wallet.CheckMnemonic(decryptoMnemonic(string(pass))) == true {
		return wallet
	}
	return nil
}

func useridToIndex(userID string) string {
	var cipherStr = make([]byte, 32)
	sha := sha3.NewKeccak256()
	sha.Write([]byte(userID))
	cipherStr = sha.Sum(nil)

	v := cipherStr[:4]
	if v[0] > 0x80 {
		v[0] = v[0] - 0x80
	}
	i := int(v[0])<<24 + int(v[1])<<16 + int(v[2])<<8 + int(v[3])
	s := strconv.Itoa(i)
	//fmt.Println("userID:", userID, "-> index:", s)
	return s
}

// NewHDAccount 创建新账户
func NewHDAccount(userID string) error {
	if wallet == nil {
		return errors.New("no wallet")
	}
	index := useridToIndex(userID)
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
	newAcc, err := ks.ImportECDSA(key, "m44600"+index)
	if err != nil {
		return err
	}
	createAccountInfoToDB(userID, newAcc.URL.Path, account.Address.Hex())
	cmn.Logger.Debug("CREATE NEW ACCOUNT address:", account.Address.Hex())
	cmn.Logger.Debug("CREATE NEW ACCOUNT path:", newAcc.URL.Path)
	return nil
}
