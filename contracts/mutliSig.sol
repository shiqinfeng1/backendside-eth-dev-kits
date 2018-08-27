pragma solidity 0.4.14;

contract SimpleMultiSig {

  uint public nonce;                // 唯一可变的状态
  uint public threshold;            // 不可变状态
  mapping (address => bool) isOwner; // 不可变状态
  address[] public ownersArr;        // 不可变状态

  function SimpleMultiSig(uint threshold_, address[] owners_) {
    if (owners_.length > 10 ||  //签名者数量不允许超过10个
        threshold_ > owners_.length || //
        threshold_ == 0) {
            throw;
    }
    //记录签名者
    for (uint i=0; i< owners_.length; i++) {
      isOwner[owners_[i]] = true;
    }
    ownersArr = owners_;
    threshold = threshold_;
  }

  function execute(uint8[] sigV, bytes32[] sigR, bytes32[] sigS, address destination, uint value, bytes data) {
    if (sigR.length != threshold) {throw;} 
    if (sigR.length != sigS.length || sigR.length != sigV.length) {throw;}

    //计算交易hash. 参考 https://github.com/ethereum/EIPs/issues/191
    bytes32 txHash = sha3(byte(0x19), byte(0), this, destination, value, data, nonce);

    address lastAdd = address(0); // cannot have address(0) as an owner
    for (uint i = 0; i < threshold; i++) {
        address recovered = ecrecover(txHash, sigV[i], sigR[i], sigS[i]);
        if (recovered <= lastAdd || !isOwner[recovered]) throw;
        lastAdd = recovered;
    }

    nonce = nonce + 1;
    if (!destination.call.value(value)(data)) {
        throw;
    }
  }

  function () payable {}
}