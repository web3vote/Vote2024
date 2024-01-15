//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


import "@openzeppelin/contracts/access/Ownable.sol";  
import "@openzeppelin/contracts/access/AccessControl.sol";

//import "./Dictionary.sol";

//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
//import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";
//import "hardhat/console.sol";


/**
 * @title Passport
 * @author zer0_ex_Bekket
 * @notice Passport verification contract for Vote2024
 */




contract Passport is Ownable, AccessControl {
    uint private _passportFee;
    address private _owner;
    bytes32 public constant moderator = keccak256("moderator");
    bytes32 public constant T3P = keccak256("T3P");

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
       // int valid_ttp_count;
        address[] ttp_proofs;
       // int peer_proof_count;
        address[] peer_proof;
        //mapping (address => bool) peer_trust;
    }



    struct TTP {
        address service;
        address ENS_address;    // we can get domain or other info about TTP from there along with description
        //mapping (string => bool) proofs;
        Phase TTP_status;

    }


    // each address have array of other users who proove him. TODO: change to trust to a passport instead of address?
    mapping (address => address[]) private peer_trust_user; 

    mapping (address => string[]) private TTP_proofs;   // each TTP service have array of proven mrz_hashes_strings.

    mapping (address => mapping (PassportType => User)) private Users ;  // from address wallet to PassportType (Internal or International) to a UserCard. Remember that users can and most likely have at least two passports.

    mapping (string hash_mrz_id => User) private PassportBook; // from mrz_hash_id to User

    mapping (address => TTP) TTPS; // mapping of Trusted 3rd parties allowed to proof id. 

    constructor() Ownable(msg.sender) {
        _passportFee = 2000000000000000 wei;
        _owner = owner();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(moderator, msg.sender);
        GetKeccakHash("T3P");
        bytes32 random_data_hash = GetKeccakHash("1234567890>>>>>>>>>>>>>>0987654321");
        string memory random_hash = bytes32ToString(random_data_hash);
        addNewTTP(msg.sender,msg.sender);
        // test call
        ProoveUserByTTP(msg.sender,PassportType.International,random_hash);


    }


    // create user  TODO: consider removing
    function _createUser(address user_address, PassportType id_type, string memory mrz_uid_hash)  internal{
        User storage newUser = Users[user_address][id_type];
        newUser.wallet = user_address;
        newUser._passportType = id_type;
        newUser.mrz_hash = mrz_uid_hash;
     //   newUser.valid_ttp_count = 0; // change it here?
     //   newUser.peer_proof_count = 0;
     // newUser.ttp_proofs = new address[](0);

        Users[user_address][id_type] = newUser;
        PassportBook[mrz_uid_hash] = newUser;
    }

    // create and proove new User with TTP
    function _create_and_proove_ttp(address user_address, PassportType id_type, string memory mrz_uid_hash) internal {
        require (checkUserExist(mrz_uid_hash) == false, "user already exist");
        User storage _u = Users[user_address][id_type];
        _u.wallet = user_address;
        _u._passportType = id_type;
        _u.mrz_hash = mrz_uid_hash;
        _u.ttp_proofs.push(msg.sender); // add ttp service address

        Users[user_address][id_type] = _u;
        PassportBook[mrz_uid_hash] = _u;
    }

    // increase ttp check   TODO: consider making public + check for TTPS auth..?
    function _proove_ttp(string memory mrz_uid_hash) internal {
        require(checkUserExist(mrz_uid_hash) == true, "user is not exist");
       // User storage _u = Users[user_address][id_type];
        User storage _u = PassportBook[mrz_uid_hash];
        address user_address = _u.wallet;
        PassportType id_type = _u._passportType;

        _u.ttp_proofs.push(msg.sender); // add ttp service address
       // _u.valid_ttp_count += 1; 
        
        Users[user_address][id_type] = _u;
        PassportBook[mrz_uid_hash] = _u;
    }


    // proove user submit valid creds
    function ProoveUserByTTP(address user, PassportType id_type, string memory mrz_uid_hash)  onlyRole(T3P) public {
        require(TTPS[msg.sender].service == msg.sender, "TTP service is not registred");

        // check if user exist and if not create new
        bool u_exist = checkUserExist(mrz_uid_hash);
        if (u_exist == false) {
            _create_and_proove_ttp(user, id_type, mrz_uid_hash);
        } else {
            _proove_ttp(mrz_uid_hash);
        }

        string[] storage ttp_proofs = TTP_proofs[msg.sender];
        ttp_proofs.push(mrz_uid_hash); // commit check to global array
        TTP_proofs[msg.sender] = ttp_proofs; // commit check to mapping from TTP to array of checks.

    }


    function addNewTTP(address service_address, address ENS_address) onlyOwner public  {
        TTP storage _t = TTPS[service_address];
        _t.service = service_address;
        _t.ENS_address = ENS_address;
        _t.TTP_status = Phase.Active;
        TTPS[service_address] = _t;
        _grantRole(T3P, service_address);
    }


    function switchTTP(address service_address) onlyOwner public {
       // TTPS[service_address].TTP_status == phase_;
       Phase  _p = TTPS[service_address].TTP_status;
       if (_p == Phase.Active) {
        _p = Phase.Paused;
       } else {
        _p = Phase.Active;
       }
       TTPS[service_address].TTP_status = _p;
    }

    function pauseTTP(address service_address) onlyOwner public {
        TTPS[service_address].TTP_status = Phase.Paused;
    }


    // really public?
    function checkUserExist(string memory mrz_uid_hash) public view returns(bool) {
        User memory u = PassportBook[mrz_uid_hash];
        if (u.wallet != address(0)) {   //TODO: test it!
            return true;
        } else {
            return false;
        }
       
    }

    function checkUserHaveTypeID(string memory mrz_uid_hash, PassportType id_type) public view returns(bool) {
        require (checkUserExist(mrz_uid_hash), "user does not exist");
        User memory u = PassportBook[mrz_uid_hash];
        //require (u._passportType == id_type, "wrong passport type");
        if (u._passportType == id_type) {
            return true;
        } else {
            return false;
        }
    }


    // really necessary?
    function checkUserHavePassedTTP(string memory mrz_uid_hash, PassportType id_type) public view returns(bool){
        require (checkUserExist(mrz_uid_hash), "user does not exist");
        User memory u = PassportBook[mrz_uid_hash];
        require (u._passportType == id_type, "wrong passport type");
        address[] memory proofs = u.ttp_proofs;
        if (proofs.length > 0) {
            return true;
        } else {
            return false;
        }
    }


    // get all MRZ hashes which has been checked by this TTP address
    function _getTTP_checks_by_service(address service) internal view returns (string[] memory) {
        string[] memory checks = TTP_proofs[service];
        return checks;
    }


    // get array of ttp services which has been checked user passport by it's hash. length can be tracted as valid/invalid
    function _getTTP_proofs_of_user(string memory mrz_uid_hash) internal view returns (address[] memory) {
        User memory u = PassportBook[mrz_uid_hash];
        address[] memory user_ttp_proofs = u.ttp_proofs; // array of addresses of ttp services which has been proofed user.
        return user_ttp_proofs;

    }


    //function getUserPassedTTPS()

   /**
    * 
    * @dev get keccak256 hash from string
    */
   function GetKeccakHash(string memory text) public pure returns (bytes32) {
        return keccak256(abi.encode(text));
    }




    function bytes32ToString(bytes32 _bytes32) internal pure returns (string memory) {
        uint8 i = 0;
        while(i < 32 && _bytes32[i] != 0) {
            i++;
        }
        
        bytes memory bytesArray = new bytes(i);
        
        for (uint8 j = 0; j < i; j++) {
            bytesArray[j] = _bytes32[j];
        }
        
        return string(bytesArray);
    }
}
