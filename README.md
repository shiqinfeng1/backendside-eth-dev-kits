# backendside-eth-dev-kits
it is backend-side ethereum develop kits.

功能如下：
## 服务后端接口调用框架
nothing descripton
## 通过abigen自动生成智能合约绑定的go源文件

文件： service/eth/eth.go

说明： 编译智能合约文件需要安装solc编译器，以及本项目路径下的abigen工具。注意：该工具在以太坊官方版本基础上对参数的处理有一点改进。

例子：
```
      import (
	      "github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
      ）
      
      eth.CompileSolidity("path/to/source/solidity", "path/to/dest/output.go", "ERC20,ERC20Basic,Ownable")
      
```
## 通过hd钱包和keystroe管理用户账户

文件： service/accounbts/hdwallet.go

说明： NewRootHDWallet创建一个hd钱包，并对助记词进行加密保存到文件，下次直接解密助记词文件并导入钱包。创建用户账户时，根据用户名生成唯一path，并新建账户保存到keystore。注意：当前代码需要mysql数据库保存user的账户信息。

例子：
```
      import (
	      "github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
      ）
      accounts.NewRootHDWallet()
      
      NewAccount("userID")
```
## 智能合约接口调用
## 以太坊web3接口调用

文件:  service/eth/web3.go

说明： client通过gorequest连接到以太坊节点。client实现以太坊web3定义的标准接口。

例子：
``` 
      import (
        "github.com/ethereum/go-ethereum/common"
	      "github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
      ）
      con := eth.NewClient("http://localhost:8545")
      addr := common.HexToAddress("0x1dcef12e93b0abf2d36f723e8b59cc762775d513")
      v, err := con.EthGetBalance(addr, nil)
      if err != nil {
	fmt.Println(v, err)
	return nil
      }
      fmt.Println(v)
```
