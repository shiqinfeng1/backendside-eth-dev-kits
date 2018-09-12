package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	ethacc "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
	hdwallet "github.com/shiqinfeng1/go-ethereum-hdwallet"
)

// ImportKeystore 导入keystore下账户到数据库
func ImportKeystore() error {
	//keys := readKeystore(userAddress)
	files, _ := ioutil.ReadDir("./keystore")
	for _, f := range files {
		path, _ := filepath.Abs("./keystore/" + f.Name())
		keyjson, err := ioutil.ReadFile(path)
		if err != nil {
			cmn.Logger.Errorf("Failed to read key: %v", err)
			return err
		}
		m := make(map[string]interface{})
		if err := json.Unmarshal(keyjson, &m); err != nil {
			return err
		}
		if address, ok := m["address"].(string); ok {
			userAddr := "0x" + address
			createAccountInfoToDB(GenerateUserIDForKeystore(userAddr), path, userAddr)
		}
	}
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
func GetadminAddress(id string) (common.Address, error) {
	//TODO: 检查是否是管理员
	accountInfo, err := getAccountInfo(id)
	if err != nil {
		cmn.Logger.Error(err)
		return common.Address{}, err
	}

	return common.HexToAddress(accountInfo.Address), nil
}

func createAccountInfoToDB(userID, path, address string) error {
	address = strings.ToLower(address)
	accountinfo := &db.AccountInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.AccountInfo{}).Where("address = ?", address).Find(accountinfo).RecordNotFound()
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
	//cmn.Logger.Errorf("userID: %v index: %v", userID, index)
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

//GenerateUserIDForKeystore 生成默认的账户名称
func GenerateUserIDForKeystore(userAddr string) string {
	return "KEYSTORE_ACCOUNT_" + userAddr
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

//GetAccountFromKeystore 从keystore中解析得到账户
func GetAccountFromKeystore(userAddress string, _passphrase string) (*keystore.Key, error) {
	keys := readKeystore(userAddress)
	key, err := keystore.DecryptKey(keys, _passphrase)
	if err != nil {
		cmn.Logger.Errorf("Failed to decrypt key: %v", err)
		return nil, err
	}
	return key, nil
}
