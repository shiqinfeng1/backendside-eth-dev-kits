pragma solidity ^0.4.23;

import "./StandardBurnableToken.sol";
import "./DetailedERC20.sol";

/**
 * @title Standard Burnable Token
 * @dev Adds burnFrom method to ERC20 implementations
 */
contract OMC is StandardBurnableToken,DetailedERC20 {
    using SafeMath for uint256;
    address techAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513; //  10%
    address communityAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513; // 5% 
    address legalAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513; // 2%
    address marketAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513; // 3%
    address foundationAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513; // 40%
    address charityAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513;   //10%
    address posAddr = 0x1dcef12e93b0abf2d36f723e8b59cc762775d513;   //30%
    
    constructor() 
        DetailedERC20("OMChain Token", "OMC",6)
        public {
            totalSupply_ =             1000000000e6;
            balances[techAddr] =        100000000e6;
            balances[communityAddr] =    50000000e6;
            balances[legalAddr] =        20000000e6;
            balances[marketAddr] =       30000000e6;
            balances[foundationAddr] =  400000000e6;
            balances[charityAddr] =     100000000e6;
            balances[posAddr] =         300000000e6;
        }
  
}