//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
//import '@ensdomains/ens-contracts/contracts/registry/ENS.sol';
//import "@ensdomains/ens-contracts/contracts/resolvers/PublicResolver.sol";

import "./Passport.sol";
//import "hardhat/console.sol";

//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
//import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";

abstract contract ENS {
    function resolver(bytes32 node) public view virtual returns (Resolver);
}

abstract contract Resolver {
    function addr(bytes32 node) public view virtual returns (address);
}

/**
 * @title Vote2024
 * @author zer0_ex_Bekket
 * @notice Vote contract for Vote2024
 */

contract Vote is Ownable, AccessControl {
    //ENS public ens;
    ENS ens = ENS(0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e);

    address private _owner;
    bytes32 public constant moderator = keccak256("moderator");
    bytes32 public constant T3P = keccak256("T3P");
    Passport private PC;
    uint uid_vote_global_counter = 0; // how much in total.

    constructor(address passport_contract) Ownable(msg.sender) {
        _owner = owner();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(moderator, msg.sender);
        PC = Passport(passport_contract);
        // ens = ENS(ens_api);

        //  addNewTTP(msg.sender,msg.sender);
    }

    enum Phase {
        IsNotStarted,
        Started,
        Finished,
        Finalized
    }

    enum VoteType {
        FreePromt,
        ENS_Valid,
        T3P_and_ENS
    }

    struct Voting {
        uint uid;
        address ens_domain; // or orginiser
        address operator;
        string desc; // description
        uint start_date;
        //uint end_date;
        uint time_to_vote_hours;
        Passport.PassportType id_type_required; // can vote only internal or only external. by default if all citizens have internal passport, everyone is allowed. opposition in exile can be elected only by exiled russians if seted.
        // mapping(string option => uint count) free_results;  // if it's "open promt vote" than we just count how much for each option
        uint votes_total;
        VoteType vote_type;
        //  mapping(address => bool) users_voted;
    }

    mapping(uint256 => Voting) public votings;
    //mapping (address orginiser =>Voting[]) public VotingsByOrg;   // can be a few

    //mapping (Voting voting => mapping(string option => uint count)) public Results;
    mapping(uint uid_vote => mapping(bytes32 option => uint count))
        public VoteResultsFreePromt;
    mapping(uint uid_vote => mapping(address user => bool voted))
        private UsersVoted;

    event FreeVoteCommited(
        uint indexed uid_event,
        string promt,
        bytes32 promt_hash,
        uint candidate_total_votes
    );
    event ENSVoteCommited(
        uint indexed uit_event,
        string promt,
        bytes32 promt_hash,
        uint candidate_total_votes
    );
    event ENS_T3P_VoteCommited(
        uint indexed uit_event,
        string promt,
        bytes32 promt_hash,
        uint candidate_total_votes
    );

    // Modifiers
    modifier voteExist(uint uid_event) {
        require(checkVoteExist(uid_event), "vote does not exits");
        _;
    }

    modifier voteActive(uint uid_event) {
        require(checkVotePhase(uid_event), "Vote is not active");
        _;
    }

    // YOU VOTE ONLY ONCE
    modifier YOVO(uint uid_event) {
        bool user_voted = UsersVoted[uid_event][msg.sender];
        require(user_voted == false, "user already voted");
        _;
    }

    modifier canidateMustRegisterENS(string memory promt_choice) {
        address target_address = checkENS_User_by_string(promt_choice);
        require(
            target_address != address(0),
            "promt_choice is not registred in ENS"
        );
        _;
    }

    function createNewVote(
        address orginiser_or_ens,
        address operator,
        uint start_date_timestamp,
        uint vote_hours,
        Passport.PassportType id_type_required,
        VoteType vote_type
    ) public onlyOwner returns (uint) {
        Voting memory v;
        // Org memory o = Org(orginiser_or_ens, operator);
        // v._org = o;
        v.ens_domain = orginiser_or_ens;
        v.operator = operator;
        v.start_date = start_date_timestamp;
        v.time_to_vote_hours = vote_hours;
        v.id_type_required = id_type_required;
        v.uid = uid_vote_global_counter;
        v.vote_type = vote_type;
        //Votings[uid_vote_global_counter] = v;
        votings[uid_vote_global_counter] = v;
        // Voting[] storage vbo = VotingsByOrg[orginiser_or_ens];
        //vbo.push(v);
        uid_vote_global_counter += 1; // how much in total
        return uid_vote_global_counter - 1; // real uid is position in array which mean real uid starts with 0
    }

    function checkVotePhase(
        uint uid_event
    ) public view voteExist(uid_event) returns (bool) {
        // require(checkVoteExist(uid_event), "vote does not exits");
        Voting memory v = votings[uid_event];
        uint start_date = v.start_date;
        uint ttv = v.time_to_vote_hours;
        uint end_date = start_date + ttv * 1 hours;
        uint256 current_time = block.timestamp;

        if (current_time >= start_date && current_time <= end_date) {
            return true;
        } else {
            return false;
        }
    }

    function getVoteStatus(
        uint uid_event
    ) public view voteExist(uid_event) returns (Phase ph) {
        // require(checkVoteExist(uid_event), "vote does not exits");
        Voting memory v = votings[uid_event];
        uint start_date = v.start_date;
        uint ttv = v.time_to_vote_hours;
        uint end_date = start_date + ttv * 1 hours;
        uint256 current_time = block.timestamp;
        // Phase ph;
        if (current_time < start_date) {
            return Phase.IsNotStarted;
        }
        if (checkVotePhase(uid_event) == true) {
            return Phase.Started;
        }
        if (current_time > end_date) {
            return Phase.Finished;
        }
    }

    function checkVoteExist(uint uid_event) public view returns (bool) {
        Voting memory v = votings[uid_event];
        uint sd = v.start_date;
        if (sd == 0) {
            return false;
        } else {
            return true;
        }
    }

    function checkENS_User_by_string(
        string memory name
    ) public view returns (address) {
        bytes32 name_hash = PC.GetKeccakHash(name);
        address user_address = resolve(name_hash);
        return user_address;
    }

    // check that vorer address have required id_type
    function voterHaveIDType(
        address voter,
        Passport.PassportType id_type_req
    ) internal view returns (bool) {
        bool id_type_met = PC.CheckUserHaveTypeIDByAddr(voter, id_type_req);
        require(
            id_type_met == true,
            "user does not registred requirement id type"
        );
        return id_type_met;
    }

    function _commitVote(
        uint uid_event,
        string memory promt_choice,
        VoteType type_of_vote
    ) internal {
        Voting memory v = votings[uid_event];

        require(v.vote_type == type_of_vote, "wrong vote type");

        if (type_of_vote == VoteType.FreePromt) {
            // check free
        }
        if (type_of_vote == VoteType.ENS_Valid) {
            // check that delegate (promt_choice) is registred in ENS by it's hash
            // already checked by modifiers
        }
        if (type_of_vote == VoteType.T3P_and_ENS) {
            // check that promt_choice is registred in ENS and have passed our T3P check.
            // already check in main func
        }

        bytes32 hash = PC.GetKeccakHash(promt_choice); // serialize to hash from string
        uint option_counter_results = VoteResultsFreePromt[uid_event][hash];
        option_counter_results += 1;
        VoteResultsFreePromt[uid_event][hash] = option_counter_results;
        UsersVoted[uid_event][msg.sender] = true;
        v.votes_total += 1;

        votings[uid_event] = v;

        // Emit Events
        if (type_of_vote == VoteType.FreePromt) {
            emit FreeVoteCommited(
                uid_event,
                promt_choice,
                hash,
                option_counter_results
            );
        }
        if (type_of_vote == VoteType.ENS_Valid) {
            emit ENSVoteCommited(
                uid_event,
                promt_choice,
                hash,
                option_counter_results
            );
        }
        if (type_of_vote == VoteType.T3P_and_ENS) {
            emit ENS_T3P_VoteCommited(
                uid_event,
                promt_choice,
                hash,
                option_counter_results
            );
        }
    }

    // Free promt
    function CommitChoiceFreePromt(
        uint uid_event,
        string memory promt_choice
    ) public voteExist(uid_event) voteActive(uid_event) YOVO(uid_event) {
        Voting memory v = votings[uid_event];
        Passport.PassportType id_type_req = v.id_type_required;
        require(voterHaveIDType(msg.sender, id_type_req));
        _commitVote(uid_event, promt_choice, v.vote_type);
    }

    // additional check that promt_choice is registred in ENS and address
    function CommitChoiceENSValid(
        uint uid_event,
        string memory promt_choice
    )
        public
        voteExist(uid_event)
        voteActive(uid_event)
        YOVO(uid_event)
        canidateMustRegisterENS(promt_choice)
    {
        Voting memory v = votings[uid_event];
        Passport.PassportType id_type_req = v.id_type_required;
        require(voterHaveIDType(msg.sender, id_type_req));
        _commitVote(uid_event, promt_choice, v.vote_type);
    }

    function CommitChoice_ENS_and_T3P(
        uint uid_event,
        string memory promt_choice
    ) public voteExist(uid_event) voteActive(uid_event) YOVO(uid_event) {
        Voting memory v = votings[uid_event];
        Passport.PassportType id_type_req = v.id_type_required;
        require(voterHaveIDType(msg.sender, id_type_req));

        address target_address = checkENS_User_by_string(promt_choice);
        require(
            target_address != address(0),
            "promt_choice is not registred in ENS"
        );
        bool checked_Passport = PC.CheckUserHavePassedTTP_ByAddr(
            target_address,
            id_type_req
        );
        require(checked_Passport == true, "candidate didn't process T3P check");

        _commitVote(uid_event, promt_choice, v.vote_type);
    }

    // ENS Resolve
    function resolve(bytes32 node) public view returns (address) {
        Resolver resolver = ens.resolver(node);
        return resolver.addr(node);
    }
}
