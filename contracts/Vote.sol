//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


import "@openzeppelin/contracts/access/Ownable.sol";  
import "@openzeppelin/contracts/access/AccessControl.sol";

import "./Passport.sol";

//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
//import "../node_modules/@openzeppelin/contracts/access/AccessControl.sol";
//import "hardhat/console.sol";


/**
 * @title Vote2024
 * @author zer0_ex_Bekket
 * @notice Vote contract for Vote2024
 */


contract Vote is Ownable, AccessControl {


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
      //  addNewTTP(msg.sender,msg.sender);
        }


    enum Phase {
        IsNotStarted,Started,Finished, Finalized
    }

    enum VoteType {
        FreePromt, ENS_Valid, T3P_and_ENS
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



    //mapping (uint uid_vote => Voting) public Votings; // All voting event by id
    Voting[] Votings; // All voting event by uid
    //mapping (address orginiser =>Voting[]) public VotingsByOrg;   // can be a few

    //mapping (Voting voting => mapping(string option => uint count)) public Results;
    mapping (uint uid_vote => mapping(string option => uint count)) public VoteResultsFreePromt;
    mapping (uint uid_vote => mapping(address user => bool voted)) private UsersVoted;

    event FreeVoteCommited (uint indexed uid_event,string promt, uint candidate_total_votes);

    function createNewVote(address orginiser_or_ens, address operator, uint start_date_timestamp, uint vote_hours, Passport.PassportType id_type_required, VoteType vote_type) onlyOwner() public returns(uint){
        Voting storage v = Votings[uid_vote_global_counter];
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
        Votings.push(v);
       // Voting[] storage vbo = VotingsByOrg[orginiser_or_ens];
        //vbo.push(v);
        uid_vote_global_counter +=1; // how much in total
        return uid_vote_global_counter - 1; // real uid is position in array which mean real uid starts with 0
    }


    function checkVotePhase(uint uid_event) public view returns (bool) {
        require(checkVoteExist(uid_event), "vote does not exits");
        Voting memory v = Votings[uid_event];
        uint start_date = v.start_date;
        uint ttv = v.time_to_vote_hours;
        uint end_date = start_date + ttv * 1 hours;
        uint256 current_time = block.timestamp;
        
        if (current_time >= start_date && current_time <= end_date) {
            return true;
        }   else {
            return false;
        }

    }



    function getVoteStatus(uint uid_event) public view returns(Phase) {
        require(checkVoteExist(uid_event), "vote does not exits");
        Voting memory v = Votings[uid_event];
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
        Voting memory v = Votings[uid_event];
        uint sd = v.start_date;
        if (sd == 0) {
            return false;
        } else {
            return true;
        }
    }

    function checkUserDidntVote() internal returns (bool) {
        
    }

    // Free promt
    function CommitChoiceFreePromt(uint uid_event,string memory promt_choice) public {
        require(checkVoteExist(uid_event), "vote does not exits");
        Phase ph = getVoteStatus(uid_event);
        require (ph == Phase.Started, "vote is not in active phase");
        Voting storage v = Votings[uid_event];
        Passport.PassportType id_type_req = v.id_type_required;

       bool id_type_met =PC.CheckUserHaveTypeIDByAddr(msg.sender, id_type_req);
       require(id_type_met == true, "user does not registred requirement id type");
       bool user_voted = UsersVoted[uid_event][msg.sender];
       require(user_voted == false, "user already voted");


       //Voting storage v = Votings[uid];
       uint option_counter_results = VoteResultsFreePromt[uid_event][promt_choice];
       option_counter_results +=1;
       VoteResultsFreePromt[uid_event][promt_choice] = option_counter_results;
       UsersVoted[uid_event][msg.sender] = true;
       v.votes_total +=1;

       Votings[uid_event] = v;
       emit FreeVoteCommited(uid_event,promt_choice,option_counter_results);


    }

    // additional check that promt_choice is registred in ENS and address
    function CommitChoiceENSValid(uint uid_event, string memory promt_choice) public {
        require(checkVoteExist(uid_event), "vote does not exits");
        Phase ph = getVoteStatus(uid_event);
        require (ph == Phase.Started, "vote is not in active phase");
        Voting storage v = Votings[uid_event];
        Passport.PassportType id_type_req = v.id_type_required;

       bool id_type_met =PC.CheckUserHaveTypeIDByAddr(msg.sender, id_type_req);
       require(id_type_met == true, "user does not registred requirement id type");
       bool user_voted = UsersVoted[uid_event][msg.sender];
       require(user_voted == false, "user already voted");


    }


}
