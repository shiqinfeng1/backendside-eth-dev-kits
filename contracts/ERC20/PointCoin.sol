pragma solidity ^0.4.23;

import "./MintableToken.sol";
import "./BurnableToken.sol";

/**
 * @title Standard Burnable Token
 * @dev Adds burnFrom method to ERC20 implementations
 */
contract PointCoin is MintableToken,BurnableToken{
  
}