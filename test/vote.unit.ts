import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect } from "chai";
import { Contract } from "ethers";
import { ethers } from "hardhat";

describe("Vote test", async () => {
    let owner: SignerWithAddress
    let user: SignerWithAddress
    let passport: Contract
    let vote: Contract
    let votingTimestamp = 3378259006

    beforeEach(async () => {
        [owner, user] = await ethers.getSigners()
        let passportFactory = await ethers.getContractFactory("MockPassport")
        passport = await passportFactory.deploy()
        await passport.deployed()

        let voteFactory = await ethers.getContractFactory("Vote")
        vote = await voteFactory.deploy(passport.address)
        await vote.deployed()
    })

    it("⭕ Should NOT call createNewVote under regular user", async () => {
        await expect(vote.connect(user).createNewVote(ethers.constants.AddressZero, user.address, votingTimestamp, 1, 0, 0)).to.be.reverted
    })
    it("📄 Should call createNewVote", async () => {
        await vote.connect(owner).createNewVote(ethers.constants.AddressZero, owner.address, votingTimestamp, 1, 0, 0)
    })
})