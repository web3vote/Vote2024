import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { Contract } from "ethers";
import { ethers } from "hardhat";

describe("Vote test", async () => {
    let owner: SignerWithAddress
    let user: SignerWithAddress
    let passport: SignerWithAddress
    let vote: Contract

    beforeEach(async () => {
        [owner, user, passport] = await ethers.getSigners()
        let voteFactory = await ethers.getContractFactory("Vote")
        vote = await voteFactory.deploy(passport.address)
        await vote.deployed()
    })

    it("📄 test", async () => {
        // console.log(vote);
    })
})