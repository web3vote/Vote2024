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

    enum PassportType {
        Internal, International
    }

    enum Phase {
        Paused, Active
    }

    struct User {
        address wallet;
        string mrz_hash; //TODO: check if possible to convert to bytes32. do not forget about salt
        PassportType _passportType;
        int valid_ttp_count;
        address[] ttp_proofs;
        int peer_proof_count;
        //mapping (address => bool) peer_trust;
        UserPeerTrust peer_trust_data;
    }



    struct TTP {
        address service;
        address ENS_address;    // we can get domain or other info about TTP from there along with description
        //mapping (string => bool) proofs;
        TTP_Proofs proofs_data;
        Phase TTP_status;

    }

    // mapping utility structs
    struct UserPeerTrust {
        mapping (address => bool) peer_trust;
    }
    struct TTP_Proofs {
        mapping (string => bool) proofs;
    }

    
    mapping (address => mapping (PassportType => User)) private Users ;  // from address wallet to PassportType (Internal or International) to a UserCard. Remember that users can and most likely have at least two passports.

    mapping (string hash_mrz_id => User) private PassportBook;

    mapping (address => TTP) TTPS; // mapping of Trusted 3rd parties allowed to proof id. 

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


    // create user
    function _createUser(address user_address, PassportType id_type, string memory mrz_uid_hash)  internal{
        User storage newUser = Users[user_address][id_type];
        newUser.wallet = user_address;
        newUser._passportType = id_type;
        newUser.mrz_hash = mrz_uid_hash;
        newUser.valid_ttp_count = 0; // change it here?
        newUser.peer_proof_count = 0;
        newUser.ttp_proofs = new address[](0);
        
        UserPeerTrust storage upt;
        upt.peer_trust[address(0)] = false;
        newUser.peer_trust_data = upt;

        Users[user_address][id_type] = newUser;

    }

    // proove user submit valid creds
    function ProoveUserByTTP(address user, PassportType id_type, string memory mrz_uid_hash)  onlyRole(moderator) public {
        require(TTPS[msg.sender].service == msg.sender, "TTP service is not registred");
    }


    function addNewTTP(address service_address, address ENS_address) onlyOwner public  {
        
    }


    function switchTTP(address service_address, Phase phase_) onlyOwner public {
        TTPS[service_address].TTP_status == phase_;
    }


    // really public?
    function checkUserExist(string memory mrz_uid_hash) public returns(bool) {

    }

    function checkUserHaveTypeID(string memory mrz_uid_hash, PassportType id_type) public returns(bool) {

    }


    // really necessary?
    function checkUserHavePassedTTP(string memory mrz_uid_hash, PassportType id_type) public {

    }

   /**
    * 
    * @dev get keccak256 hash from string
    */
   function GetKeccakHash(string memory text) public pure returns (bytes32) {
        return keccak256(abi.encode(text));
    }
}
