//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;






import "@openzeppelin/contracts/access/Ownable.sol";  
import "@openzeppelin/contracts/access/AccessControl.sol";

//import "./Dictionary.sol";

//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
//import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";
//import "hardhat/console.sol";

contract Passport is Ownable, AccessControl {
    uint private _passportFee;
    address private _owner;
    bytes32 public constant moderator = keccak256("moderator");


       constructor(

        string memory plain_id
    ) Ownable(msg.sender) {
        _passportFee = 2000000000000000 wei;
        _owner = owner();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(moderator, msg.sender);
       // Turing = Dictionary(turing_);
        bytes32 h_id = GetKeccakHash(plain_id);


    }



   /**
    * 
    * @dev get keccak256 hash from string
    */
   function GetKeccakHash(string memory text) public pure returns (bytes32) {
        return keccak256(abi.encode(text));
    }
}
