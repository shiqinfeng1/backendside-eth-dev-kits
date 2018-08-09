package accounts

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// KEY, 必须是16,24,32位的[]byte
// 分别对应AES_128,AES_192,AES_256

func cryptoAndSave(text, password string) error {
	// 需要加密的字符串
	plaintext := []byte(text)
	file, err := os.OpenFile(common.Config().GetString("hdwallet.stored"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", common.Config().GetString("hdwallet.stored"))
		os.Exit(0)
	}
	defer file.Close()
	var cipherStr = make([]byte, 32)
	sha := sha3.NewKeccak256()
	sha.Write([]byte(password))
	cipherStr = sha.Sum(nil)
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(cipherStr))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(cipherStr), err)
		return err
	}
	// 加密模式 (ECB、CBC、CFB、OFB)
	// IV是initialization vector的意思
	// 就是加密的初始话矢量，初始化加密函数的变量
	// 也就是加密动作中的 数据操作的偏移量
	cfb := cipher.NewCFBEncrypter(c, commonIV)

	// 存储密码, 必须与块体的长度相同
	ciphertext := make([]byte, len(plaintext))

	// 流化, 必须与块体的长度相同
	cfb.XORKeyStream(ciphertext, plaintext)

	// 写入文件
	err = ioutil.WriteFile(common.Config().GetString("hdwallet.stored"), ciphertext, 0777)
	if err != nil {
		fmt.Println("保存加密后文件失败!")
		return err
	}
	fmt.Println("文件已加密!")

	return nil
}

func decryptoMnemonic(password string) string {

	var cipherStr = make([]byte, 32)
	sha := sha3.NewKeccak256()
	sha.Write([]byte(password))
	cipherStr = sha.Sum(nil)
	// 读文件
	ciphertext, err := ioutil.ReadFile(common.Config().GetString("hdwallet.stored"))
	if err != nil {
		fmt.Println("读取加密文件失败!")
		return ""
	}
	fmt.Println("读取加密文件成功!")

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(cipherStr))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(cipherStr), err)
		return ""
	}

	// 解密模式 (ECB、CBC、CFB、OFB)
	// 也就是说，解密的时候也需要加密时的密钥与偏移量
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)

	//cs, _ := base64.StdEncoding.DecodeString(string(ciphertext))
	// 存储数据, 必须与块体的长度相同
	plaintextCopy := make([]byte, len(ciphertext))

	// 流化, 必须与块体的长度相同
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	return string(plaintextCopy)
}
