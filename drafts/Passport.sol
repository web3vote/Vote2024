//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;






import "@openzeppelin/contracts/access/Ownable.sol";  
import "@openzeppelin/contracts/access/AccessControl.sol";

//import "./Dictionary.sol";

//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
//import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";
//import "hardhat/console.sol";

contract Leg10n is Ownable, AccessControl {
    uint private _passportFee;
    address private _owner;
    bytes32 public constant moderator = keccak256("moderator");

    struct User {
        address userAddress;
        string tgId;  // unic Id for telegram (value encrypted by bot)
        bool valid;
        address validatorAddress;
        string codeName;
        string public_key;
        bytes32 id_hash;    // hash from plain id... possibly we need to make keccak256 functions work in front-end
    }

    //mappings
    mapping(string => address) public tgIdToAddress;     // encrypted id => address
    mapping(bytes32 => string) public hashToId;          // hash of plain text id to encrypted_id
    mapping(address => User) private users;
    mapping(string => address) public codename_wallets; // codenames protected by dictionary
    mapping(address => mapping(address => bool)) public chain; // from parent to child to flag

    // EVENTS

    event Initialized(address indexed admin);
    event requestDenied(string applyerTg, address wallet);

    event joinRequested(
        string applyerTg,
        address wallet_address,
        address indexed parent_address
    );
    event joinRequestedIndexedTG(
        string applyerTg,
        address wallet_address,
        address indexed parent_address
    );
    event requestAccepted(
        string indexed applyerTg,
        address user_address,
        address parent_address
    );
    event relationChanged(
        address high_node,
        address indexed low_node,
        bool pravda
    ); // we can search by user_address (low_node address) to lookup user relationship

    // Roles Engine
    //Dictionary Turing;

    constructor(
        address turing_,
        address admin_,
        string memory tgid_,
        string memory public_key_,
        string memory plain_id
    ) Ownable() {
        _passportFee = 2000000000000000 wei;
        _owner = owner();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(moderator, msg.sender);
       // Turing = Dictionary(turing_);
        bytes32 h_id = GetKeccakHash(plain_id);
        // Factory init
        tgIdToAddress[tgid_] = msg.sender;
        codename_wallets["0"] = msg.sender;
        users[msg.sender] = User(
            msg.sender,
            tgid_,
            true,
            msg.sender,
            "0",
            "zjXCj9iuse3gHGaAIIgyaiCOsJpQWSCEBBac/zPGrgQ=",
            h_id
        );
        // determine if contract deployed standalone or through factory contract
        if (admin_ == address(0x0)) {
            // standalone deploy, deployer is admin
            _devInitAdmin(
                msg.sender,
                tgid_,
                "zjXCj9iuse3gHGaAIIgyaiCOsJpQWSCEBBac/zPGrgQ=",
                plain_id
            );
        } else {
            // factory deploy, factory inputs admin credentials
            _devInitAdmin(admin_, tgid_, public_key_,plain_id);
        }
    }

    function _updateAddress(
        string memory tgId,
        address userAddress,
        string memory code_name_,
        string memory parent_name,
        string memory plain_id
    ) internal {
       bytes32 hash_id = GetKeccakHash(plain_id);

        require(
            tgIdToAddress[tgId] == address(0x0),
            "There is address connected to that ID already"
        ); // if cell is not empty revert
        require(
               bytes(hashToId[hash_id]).length == 0,
               "This plain_id already registred"
        );
     //   bool flag = Turing.checkDictionaryTree(code_name_, parent_name);
       // require(flag == true, "dictionary error");
        require(
            codename_wallets[code_name_] == address(0x0),
            "codename taken"
        );
        hashToId[hash_id] = tgId;
        tgIdToAddress[tgId] = userAddress;
       // require(codename_wallets[code_name_] == address(0x0), "codename is already taken");
        codename_wallets[code_name_] = userAddress;
    }

    /**
     *  @dev This function is main entrance point, it is create join Request. User shoul know codename of superior inviter.
     * @param applyerTg tgid of user who wants to join
     * @param code_name_  code_name selected by user
     * @param parent_name  code_name of parent_node
     */
    function RequestJoin(
        string memory applyerTg,
        string memory code_name_,
        string memory parent_name,
        string memory public_key,
        string memory plain_id
    ) public payable {
        address applyerAddress = msg.sender; // ЛИЧНАЯ ПОДАЧА ПАСПОРТА В ТРЕТЬЕ ОКОШКО МФЦ

        require(msg.value == _passportFee, "Request fee is not paid");

        address parent_address = codename_wallets[parent_name];
        bytes32 hash_id = GetKeccakHash(plain_id);
        users[msg.sender] = User(
            applyerAddress,
            applyerTg,
            false,
            parent_address,
            code_name_,
            public_key,
            hash_id
        );
        // TODO: add codename_wallets[username] = msg.sender;
        (bool feePaid, ) = _owner.call{value: _passportFee}("");
        require(feePaid, "Unable to transfer fee");
        _updateAddress(applyerTg, applyerAddress, code_name_, parent_name,plain_id);

        emit joinRequested(applyerTg, msg.sender, parent_address);
        emit joinRequestedIndexedTG(applyerTg, msg.sender, parent_address);

        chain[parent_address][msg.sender] = false;
    }

    /**
     *  @dev Accept user intent to join
     * @param applyerTg tgid of user who want to join
     * @param parent_name code_name of parent_node
     */
    function AcceptJoin(string memory applyerTg, string memory parent_name) public {
        address parent_address = codename_wallets[parent_name];
        require(parent_address == msg.sender, "only parent_name can accept it");
        address user_address = GetUserWalletByID(applyerTg);
        bool isRegistred = chain[parent_address][user_address];
        require(isRegistred == false, "already registred");
        users[user_address].valid = true;
        users[user_address].validatorAddress = msg.sender;
        chain[parent_address][user_address] = true;
        emit relationChanged(parent_address, user_address, true);
    }

    /**
     *     @notice This function decline application end erase junk data
     *
     */
    function DeclineRequest(string memory tgid) public {
        address child_address = GetUserWalletByID(tgid);
        //int64 parent_id = GetTgIdByAddress(msg.sender);
        //string memory parent_name = users[msg.sender].codeName;
        string memory user_name_ = users[child_address].codeName;
        require(
            users[child_address].valid == false,
            "already approved OR do not exists yet"
        ); // it also means that record exists
        bool linked = chain[msg.sender][child_address];
        require(linked == false, "already linked OR dont exists");
        delete users[child_address];
        delete tgIdToAddress[tgid];
        delete codename_wallets[user_name_];
        delete chain[msg.sender][child_address];
        emit requestDenied(tgid, child_address);
    }

    // todo -- rework, make banhammer accesible for admin, add deletion of chain of command
    /**
     *  @dev This function is a service function which allow Owner to erase already approved passport
     *  and make clean state contract. NOT FOR USE IN PRODUCTION
     *  it does not clear chain of command
     */
    function devDeleteUser(address user_address) public onlyRole(moderator) {
        string memory _tgId = users[user_address].tgId;
        string memory user_name_ = users[user_address].codeName;
        uint chainID = block.chainid;
        require(chainID == uint(5), "testnet only");
        delete users[user_address];
        delete tgIdToAddress[_tgId];
        delete codename_wallets[user_name_];
        //delete chain[msg.sender][child_address];
        emit requestDenied(_tgId, user_address);
    }

    function _devInitAdmin(
        address admin_,
        string memory tgid_,
        string memory public_key_,
        string memory plain_id
    ) internal {
        tgIdToAddress[tgid_] = admin_;
        codename_wallets["Adam"] = admin_;
        bytes32 hash_id = GetKeccakHash(plain_id);
        users[admin_] = User(admin_, tgid_, true, admin_, "Adam", public_key_,hash_id);
        users[admin_].valid = true;
        users[admin_].validatorAddress = msg.sender;
        address zero = GetWalletByNickName("0");
        chain[zero][admin_] = true;
        hashToId[hash_id] = tgid_;
        emit Initialized(admin_);
    }

    /**
     *  @dev This function is a service function which allow delete profile of user. It does not clear command chain, so it's required to call ClearParent first
     */
    function DeleteUser(address user_address) internal {
        string memory _tgId = users[user_address].tgId;
        string memory user_name_ = users[user_address].codeName;
        delete users[user_address];
        delete tgIdToAddress[_tgId];
        delete codename_wallets[user_name_];
        //delete chain[msg.sender][child_address];
        emit requestDenied(_tgId, user_address);
    }


    // TODO: rework, make ClearParent and DeleteUser by single tx
    /**
     *  @dev delete yourself profile
     */
    function deleteYourSelf() public {
        bool attached = users[msg.sender].valid;
        require(attached == false, "call ClearParent first");
        DeleteUser(msg.sender);
    }

    /**
     *  @dev allow users to clear their parenthesis (chain of command)
     *  @param parent_name name of high node
     *  @param child_name username
     */
    function ClearParent(
        string memory parent_name,
        string memory child_name
    ) public {
        address user_address = GetWalletByNickName(child_name);
        address parent_address = GetWalletByNickName(parent_name);
        require(
            user_address == msg.sender,
            "users allowed only to clear themselfs"
        );
        users[user_address].valid = false;
        chain[parent_address][msg.sender] = false;
    }
    
    /**
     *  @dev setting fee for applying for passport
     */
    function SetPassportFee(uint passportFee_) public onlyOwner {
        _passportFee = passportFee_;
    }

    /**
     *  @dev getter to obtain how much user will pay for apply
     */
    function GetPassportFee() public view returns (uint) {
        return _passportFee;
    }

    function GetUserWalletByID(string memory tgId_) public view returns (address) {
        return tgIdToAddress[tgId_];
    }

    function GetTgIdByAddress(
        address user_wallet
    ) public view returns (string memory tgid) {
        User memory u = GetUserByAddress(user_wallet);
        tgid = u.tgId;
        return tgid;
    }

    function GetUserByAddress(
        address user_wallet
    ) public view returns (User memory) {
        User memory u = users[user_wallet];
        return u;
    }

    //function GetUserByAddress(address user_wallet) public view returns

    function GetWalletByNickName(
        string memory user_name_
    ) public view returns (address) {
        return codename_wallets[user_name_];
    }

    function GetUserByNickName(
        string memory user_name_
    ) public view returns (User memory) {
        address wallet_ = GetWalletByNickName(user_name_);
        User memory u = users[wallet_];
        return u;
    }

    function GetUserByTgId(string memory tgId_) public view returns (User memory) {
        address wallet = GetUserWalletByID(tgId_);
        User memory u = users[wallet];
        return u;
    }

    function GetPublicKeyByAddress(
        address user_address
    ) public view returns (string memory) {
        User memory u = GetUserByAddress(user_address);
        string memory pk = u.public_key;
        return pk;
    }

    function GetOwner() public view returns (address) {
        return _owner;
    }

    function getModeratorIdentifier() public pure returns (bytes32) {
        return moderator;
    }

    /**
     * @dev return encrypted id associated with hash from plain text id
     * @param hash_id hash id of plain text id
     * 
     */
    function GetIdByHash(bytes32 hash_id) public view returns (string memory) {
      return hashToId[hash_id];
    }

   /**
    * 
    * @dev get keccak256 hash from string
    */
   function GetKeccakHash(string memory text) public pure returns (bytes32) {
        return keccak256(abi.encode(text));
    }
}
