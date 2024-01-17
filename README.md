# Vote 2024

This is a project of a web3 vote.

Goals of this project:
1. Make an alternative "president election" in Russia 2024.
2. Make russian emmigrants to vote for "commitie-in-exile" who will represent ours interests

Whole project contains from contracts and dapp (this repo) + backend.

# Pre-requirements and problems

### Abstract
Basically, when we are speaking about elections, we firstly think about CAP theorem, which says that *any decentralized system can have 2 out of 3 property: Consistency, Availability, Partition Tolerance, but none can have all 3 in the same time*.

(editor's note) -- there is an exeption for this rule -- system can have all 3 property but not in the *same* time. For example system can be consistent, partition tolarable and *eventually* available. Or it can be available, partition tolerable and *eventually* consistent.

Let's see what properties have putin's elections? and what properties have web3 elections?

#### Comparsion

##### 1. Availability

Putins election have availability -- vote can everyone, you can vote living in woods, you can vote from prison, you can vote from army,
They don't care how people vote, they care only about *who count* votes. If they control them -- they control elections, which mean they can actually have a highly accessible elections therefore creating illusion of democratic society and illusion of the democratic Tsar.

While they control how the votes are calculated they are still in fear of true democratic elections, which means that no one opposition candidate will be allowed to be a candidate for election. They are feared even this almost impossible chance.

Web3 elections basically pre-require access to internet. It is basically possible to deploy the system in the TOR network, and make it practically unblocable. There might be problems in case of total goverment shutdown. From other side web3 elections is allow to **vote for anyone**. 
If you want to choose **Navalny** than you can choose him and noone will procecute you.

I can't say any of these system is fully available.


##### 2. Partition-Tolerance

Both system is partition tolerable. "Izberkoms" are not depending from one another. Blockchain network nodes are also not depending from one another.

##### 3. Consistency

The major flaw in putins "elections" is consistency of data. Which we all have seen how in some regions there are "146%" for putinists.
Which, of course, should make such elections not legitimate, but it does "work" in russia.

In web3 elections we simply have a consistent state of `EVM`, finding it's implementation in smart-contracts.

Which mean that we can make a system where such things as 'carousels' (double-vote) are not possible by designe of the system. It also make it almost impossible to block, which make it censorship-resistant.



# Why


You are probably think, why the hell are we need this, if such elections will not be "legitimate"
The answer are -- nobody have legitimacy anymore.

1. Putin's own election are not legitimate cause 
a. vote manipulation
b. assasination and prosecution of other candidates 
c. nature of authoritharian fashist regime
d. EU will not accept those elections
e. EU frozen contacts with Putins regime by now by fact.

We know, that putin's elections will be compromised and manipulated in anycase, so our ultimate goal is not to "choose president", but to 
mobilize the crowd of anti-putin's powers on the streets. In order to do so we should show our supporters that they are not alone.

2. Navalny is in jail and his organisation do not give a shit about everyone else.
In the circumstances when your own goverment consider every russian emmigrant as a 'traitor' and 'national security threat', many of us have found ourself in a difficult position.

From one side we can no longer await any help from our goverment or it's institution abroads.
From other hand, our so-called "opposition" directly saying "your problems is not our problems".

When Belo-russians got into the same circumstances, their leader -- Tikhanovskaya is doing everything, to help regular belorussians in exile, including solving problems with visa-bans and bank card problems. She does accept responsability for a people who vote for her, and therefore she have legitimacy.

So we need also to choose our representative-in-exile, to start solving our problems.




# How does it work?

The whole project have a few parts:
1. Verification process
2. Vote itself

## Verification process
1. User choose type of id (internal passport / international passport ----> users can have at least two different passports)
2. User POST image to validation service on server (PassportEye + Tesseract). Service use image recognition (Tesseract) to find MRZ (machine-readable zone). It's also check that data from MRZ is the same as in the fields in the picture, also it checks noises and artifacts to detect forgery.
Serice return MRZ to user and delete upload photo.
3. User sign their MRZ with their wallet and send it to server
4. Server got MRZ and make an external call to some Third Party Service (aka TTP aka T3P) to make check that passport data is valid.
5. If all checks are passed, server compute `keccak256(string MRZ + string salt)` -- it gives us hash from an MRZ without necessary to store plain text MRZ into public blockchain. We still need, however, keep hashes from MRZ, therefore preventing one user to vote twice at same ballots (a fatal flow of russian official elections). Original MRZ is also destroyed and is not stored anywhere.
6. Server commit results (hash, address and some metadata that user have passed checks) into blockchain (`contracts/Passport.sol`). 
Now we know that some address may be attached to some identity, without compromising the identity itself.

Also in this way, users are not paying for the documents processing, making easy onboarding for users who are not familiar with crypto.

## Votings itself

1. Admin can create new votings.
2. Voting contract should check that user addresses have passed T3P passport check without getting the actual info about user.
3. User can have more than one passport. There are id_type in `Passport` contract. Voting can require single id_type for making choice commitment
Example: all russians have internal passport and some of them have international passport. Which means we can create one Voting for "president of russia" and make it accessible to every russian, and make a separate Voting for choosing their representative in (EU,USA,World) and make it accessible to only those who have international passport.
4. //TODO: continue



# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```
