pragma solidity ^0.4.23;

import "./MintableToken.sol";
import "./BurnableToken.sol";

/**
 * @title Standard Burnable Token
 * @dev Adds burnFrom method to ERC20 implementations
 */
contract PointCoin is MintableToken,BurnableToken{
    struct Data {
        address Addr; 
    }
    Data[] public addrList;
    mapping(address => uint256) public userIndex;
    function addlist(address _address) internal 
    {
        Data memory value = Data(_address);
        addrList.push(value);
    }

    function buy(address _to,uint256 _amount) public
    {
        if (userIndex[_to] == 0){
            userIndex[_to] = addrList.length;
            addlist(_to);
        }
        mint(_to,_amount);
    }

    function consume(uint256 _value) public
    {
        burn(_value);
    }

    function refund() public
    {
        burn(balances[msg.sender]);   
    }
}