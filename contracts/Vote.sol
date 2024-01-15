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
 * @notice Passport verification contract for Vote2024
 */


contract Vote is Ownable, AccessControl {


    address private _owner;
    bytes32 public constant moderator = keccak256("moderator");
    bytes32 public constant T3P = keccak256("T3P");
    Passport private PC;
    uint uid_vote_global_counter = 0;


        constructor(address passport_contract) Ownable(msg.sender) {
        _owner = owner();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(moderator, msg.sender);
        PC = Passport(passport_contract);
      //  addNewTTP(msg.sender,msg.sender);
        }




    struct Voting {
        uint uid;
        Org _org ;
        string desc; // description
        uint start_date;
        //uint end_date;
        uint time_to_vote_hours;
        Passport.PassportType id_type_required; // can vote only internal or only external. by default if all citizens have internal passport, everyone is allowed. opposition in exile can be elected only by exiled russians if seted.
       // mapping(string option => uint count) free_results;  // if it's "open promt vote" than we just count how much for each option
        uint votes_total;
      //  mapping(address => bool) users_voted;
    }

    struct Org {
        address ens_domain; // or orginiser
        address operator;
    }

    mapping (uint uid_vote => Voting) public Votings; // All voting event by id
    mapping (address orginiser =>Voting) public VotingsByOrg;   // can be a few

    //mapping (Voting voting => mapping(string option => uint count)) public Results;
    mapping (uint uid_vote => mapping(string option => uint count)) public VoteResultsFreePromt;
    mapping (uint uid_vote => mapping(address user => bool voted)) private UsersVoted;

    function createNewVote(address orginiser_or_ens, address operator, uint start_date_timestamp, uint vote_hours, Passport.PassportType id_type_required) onlyOwner() public{
        Voting storage v = Votings[uid_vote_global_counter];
        Org memory o = Org(orginiser_or_ens, operator);
        v._org = o;
        v.start_date = start_date_timestamp;
        v.time_to_vote_hours = vote_hours;
        v.id_type_required = id_type_required;
        Votings[uid_vote_global_counter] = v;
        VotingsByOrg[orginiser_or_ens] = v;
        uid_vote_global_counter +=1;
    }

}
