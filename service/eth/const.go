package eth

const (
	// HexPrefix 16进制前缀
	HexPrefix = "0x"

	// ERC20MethodTransfer erc20方法transfer的16进制字符串
	ERC20MethodTransfer = "a9059cbb"
	// ERC20TransferLength erc20方法transfer输入长度
	ERC20TransferLength = 138
	// PointsMethodBuyPoints BuyPoints
	PointsMethodBuyPoints = ""
	// PointsBuyPointsLength 方法输入数据长度
	PointsBuyPointsLength = 138
	// PointsMethodConsumePoints ConsumePoints
	PointsMethodConsumePoints = ""
	// PointsConsumePointsLength 方法输入数据长度
	PointsConsumePointsLength = 74
	// PointsMethodRefundPoints RefundPoints
	PointsMethodRefundPoints = ""
	// PointsRefundPointsLength 方法输入数据长度
	PointsRefundPointsLength = 10

	// ERC20Name erc20 name 的16进制字符串
	ERC20Name = "0x06fdde03"
	// ERC20Symbol erc20 symbol 的16进制字符串
	ERC20Symbol = "0x95d89b41"
	// ERC20Decimals erc20 decimals 的16进制字符串
	ERC20Decimals = "0x313ce567"
	// ERC20AbiDefaultLength ERC20 Abi字符串 Default Length
	ERC20AbiDefaultLength = 194

	// ConfirmedNum confirmed num till to 12
	ConfirmedNum = int64(12)

	// EthGasLimit eth gas limit transaction gas limit
	EthGasLimit = int64(21000)
	// DefaultErc20GasPrice default gas price for erc20
	DefaultErc20GasPrice = int64(60000)

	// DefaultErc20Icon default erc20 icon
	DefaultErc20Icon = "https://ops.58wallet.io/home/img/avatar-DEFAULT@2x.png"

	//ReceiptStatusSuccessful Receipt Status Successful
	ReceiptStatusSuccessful = "0x1"

	// EventTransferHash hash of Event Transfer
	EventTransferHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// AddressInHashIndex Address in hash  from index
	AddressInHashIndex = 26
)
