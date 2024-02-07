import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect } from "chai";
import { Contract } from "ethers";
import { ethers } from "hardhat";

const strToBytes = ethers.utils.formatBytes32String("string.eth");
describe("Vote test", async () => {
    let owner: SignerWithAddress
    let user: SignerWithAddress
    let passport: Contract
    let vote: Contract
    let votingTimestamp2077 = 3378259006
    let votingTimestampLate = 1705729995

    beforeEach(async () => {
        [owner, user] = await ethers.getSigners()
        let passportFactory = await ethers.getContractFactory("MockPassport")
        passport = await passportFactory.deploy()
        await passport.deployed()

        let voteFactory = await ethers.getContractFactory("Vote")
        vote = await voteFactory.deploy(passport.address)
        await vote.deployed()
    })
    it("🟡 Should NOT call createNewVote under regular user", async () => {
        await expect(vote.connect(user).createNewVote(ethers.constants.AddressZero, user.address, votingTimestamp2077, 1, 0, 0)).to.be.reverted
    })
    it("🟢 Should call createNewVote", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestamp2077, 1, 0, 0)
        let votingObj = await vote.votings(0)
        expect(votingObj.uid).to.equal(0)
        expect(votingObj.ens_domain).to.equal(ethers.constants.AddressZero)
        expect(votingObj.operator).to.equal(owner.address)
        expect(votingObj.desc).to.equal("")
        expect(votingObj.start_date).to.equal(votingTimestamp2077)
        expect(votingObj.time_to_vote_hours).to.equal(1)
        expect(votingObj.id_type_required).to.equal(0)
        expect(votingObj.vote_type).to.equal(0)
        expect(await vote.checkVoteExist(0)).to.equal(true)
        expect(await vote.getVoteStatus(0)).to.equal(0) // not started
        expect(await vote.checkVotePhase(0)).to.equal(false)
    })
    it("🟢 Should call checkVoteExist of unexisted vote", async () => {
        expect(await vote.checkVoteExist(0)).to.equal(false)
    })
    it("🟡 Should NOT call checkVotePhase because vote does not exits", async () => {
        await expect(vote.checkVotePhase(0)).to.be.reverted
    })
    it("🟢 Should call checkVotePhase of true", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 12000, 0, 0)
        expect(await vote.checkVotePhase(0)).to.equal(true)
        expect(await vote.getVoteStatus(0)).to.equal(1)
    })
    it("🟢 Should call getVoteStatus of finished votitng", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 1, 0, 0)
        expect(await vote.getVoteStatus(0)).to.equal(2)
    })
    it("🟡 Should NOT call CommitChoiceFreePromt because vote is not active", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 1, 0, 0)
        await expect(vote.CommitChoiceFreePromt(0, strToBytes)).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoiceFreePromt because user does not registred requirement id type", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 12000, 0, 0)
        await expect(vote.connect(user).CommitChoiceFreePromt(0, strToBytes)).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoiceFreePromt because user alredy voted", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 12000, 0, 0)
        await passport.createUser(owner.address, 0, strToBytes)
        await vote.connect(owner).CommitChoiceFreePromt(0, strToBytes)
        await expect(vote.connect(owner).CommitChoiceFreePromt(0, strToBytes)).to.be.reverted
    })
    it("🟢 Should call CommitChoiceFreePromt", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 12000, 0, 0)
        await passport.createUser(owner.address, 0, strToBytes)
        await vote.connect(owner).CommitChoiceFreePromt(0, strToBytes)
    })
    it("🟡 Should NOT call CommitChoiceENSValid because vote does not exits", async () => {
        await expect(vote.CommitChoiceENSValid(0, strToBytes)).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoiceENSValid because vote is not active", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 1, 0, 1)
        await expect(vote.CommitChoiceENSValid(0, strToBytes)).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoiceENSValid because", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 12000, 0, 1)
        await passport.connect(owner).createUser(owner.address, 0, strToBytes)
        await expect(vote.connect(owner).CommitChoiceENSValid(0,strToBytes)).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoice_ENS_and_T3P because vote does not exits", async () => {
        await expect(vote.CommitChoice_ENS_and_T3P(0, "string")).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoice_ENS_and_T3P because vote is not active", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 1, 0, 2)
        await expect(vote.CommitChoice_ENS_and_T3P(0, "string")).to.be.reverted
    })
    it("🟡 Should NOT call CommitChoice_ENS_and_T3P because ", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestampLate, 12000, 0, 2)
        await passport.connect(owner).createUser(owner.address, 0, strToBytes)
        // await vote.connect(owner).CommitChoice_ENS_and_T3P(0, "string.eth")
        // await vote.connect(owner).CommitChoiceENSValid(0,"string")
    })
})