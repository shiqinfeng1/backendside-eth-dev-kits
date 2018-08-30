package accounts

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	ethacc "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/howeyc/gopass"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
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

//GetRootHDWallet 获取hd根钱包
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

// NewAccount 创建新账户
func NewAccount(userID string) error {
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

func readKeystore(userAddress string) []byte {
	if userAddress[:2] == "0x" || userAddress[:2] == "0X" {
		userAddress = userAddress[2:]
	}
	files, _ := ioutil.ReadDir("./keystore")
	for _, f := range files {
		if bytes.Contains([]byte(f.Name()), []byte(strings.ToLower(userAddress))) {
			p, _ := filepath.Abs("./keystore/" + f.Name())
			keys, _ := ioutil.ReadFile(p)
			return keys
		}
	}
	return []byte{}
}

// GetTransactOptsFromHDWallet 交易调用参数
func GetTransactOptsFromHDWallet(userID string) (*bind.TransactOpts, error) {
	index := useridToIndex(userID)
	accInfo, err := getAccountInfo(userID)
	if err != nil {
		return nil, err
	}
	//首先导入上面生成的账户密钥（json）和密码
	keys, _ := ioutil.ReadFile(accInfo.Path)
	transactOpts, err := bind.NewTransactor(strings.NewReader(string(keys)), "m44600"+index)
	return transactOpts, err
}

// GetTransactOptsFromKeystore is the collection of authorization data required to create a valid Ethereum transaction.
func GetTransactOptsFromKeystore(userAddress string, _passphrase string) (*bind.TransactOpts, error) {
	keys := readKeystore(userAddress)
	key, err := keystore.DecryptKey(keys, _passphrase)
	if err != nil {
		cmn.Logger.Errorf("Failed to decrypt key: %v", err)
		return nil, err
	}
	// 对keystore采取对称加密解析出私钥
	return bind.NewKeyedTransactor(key.PrivateKey), nil
}

func getAccountFromHDWallet(index string) (*ethacc.Account, error) {
	if wallet == nil {
		return nil, errors.New("no wallet")
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + index)
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func createAccountInfoToDB(userID, path, address string) error {

	accountinfo := &db.AccountInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.AccountInfo{}).Where("address = ? and user_id = ?", address, userID).Find(accountinfo).RecordNotFound()
	if notFound {
		accountinfo.UserID = userID
		accountinfo.Path = path
		accountinfo.Address = address
		err := dbconn.Create(accountinfo).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	} else {
		err := dbconn.Model(accountinfo).Update("path", path).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	}
	dbconn.MysqlCommit()
	return nil
}

func getAccountInfo(userID string) (*db.AccountInfo, error) {
	accountInfo := &db.AccountInfo{}
	dbconn := db.MysqlBegin()
	err := dbconn.Model(&db.AccountInfo{}).Where("user_id = ?", userID).Find(&accountInfo).Error
	dbconn.MysqlRollback()
	if err != nil {
		return nil, err
	}
	return accountInfo, nil
}

//GetUserAddress 获取用户数据
func GetUserAddress(userID string) (common.Address, error) {
	accountInfo, err := getAccountInfo(userID)
	if err != nil {
		cmn.Logger.Error(err)
		return common.Address{}, err
	}

	return common.HexToAddress(accountInfo.Address), nil
}

//GetadminAddress 获取管理员账户
func GetadminAddress() (common.Address, error) {
	accountInfo, err := getAccountInfo("15422339579")
	if err != nil {
		cmn.Logger.Error(err)
		return common.Address{}, err
	}

	return common.HexToAddress(accountInfo.Address), nil
}

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
