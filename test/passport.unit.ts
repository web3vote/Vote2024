import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect, use } from "chai";
import { Contract } from "ethers";
import { ethers } from "hardhat";

const strToBytes = ethers.utils.formatBytes32String("0x0");
export const getACError = (address: string, role: string) =>
  `AccessControl: account ${address.toLowerCase()} is missing role ${role}`;

describe("Passport test", async () => {
    let owner: SignerWithAddress
    let user: SignerWithAddress
    let passport: Contract

    beforeEach(async () => {
        [owner, user] = await ethers.getSigners()
        let passportFactory = await ethers.getContractFactory("MockPassport")
        passport = await passportFactory.deploy()
        await passport.deployed()
    })

    it("🟡 Should NOT call ProoveUserByTTP under regular user", async () => {
        await expect(passport.connect(user).ProoveUserByTTP(user.address, 0, strToBytes)).to.be.reverted
    })
    it("🟢 Should call ProoveUserByTTP", async () => {
        await passport.createUser(user.address, 0, strToBytes)
        await passport.connect(owner).ProoveUserByTTP(user.address, 0, strToBytes)
        await passport.connect(user).checkUserHavePassedTTP_ByHash(strToBytes)
        await passport.connect(user).CheckUserHavePassedTTP_ByAddr(user.address, 0)
    })
    it("🟡 Should NOT call addNewTTP under regular user", async () => {
        await expect(passport.connect(user).addNewTTP(user.address, ethers.constants.AddressZero)).to.be.reverted
    })
    it("🟡 Should NOT call CheckUserHavePassedTTP_ByAddr because wallet is ZERO address", async () => {
        await expect(passport.connect(user).CheckUserHavePassedTTP_ByAddr(user.address, 0)).to.be.reverted
    })
    it("🟢 Should call CheckUserHavePassedTTP_ByAddr", async () => {
        await passport.createUser(user.address, 0, strToBytes)
        await passport.connect(user).CheckUserHavePassedTTP_ByAddr(user.address, 0)
    })
    it("🟡 Should NOT call create_and_proove_ttp because already exist", async () => {
        await passport.createUser(user.address, 0, strToBytes)
        await expect(passport.connect(user).create_and_proove_ttp(user.address, 0, strToBytes)).to.be.reverted
    })
    it("🟡 Should NOT call proove_ttp because already exist", async () => {
        await expect(passport.connect(user).proove_ttp(strToBytes)).to.be.reverted
    })
    it("🟢 Should call getTTP_checks_by_service", async () => {
        await passport.connect(user).getTTP_checks_by_service(user.address)
    })
    it("🟢 Should call getTTP_proofs_of_user", async () => {
        await passport.connect(user).getTTP_proofs_of_user(strToBytes)
    })
    it("🟡 Should NOT call checkUserHavePassedTTP_ByHash because user dont exist", async () => {
        await expect(passport.connect(user).checkUserHavePassedTTP_ByHash(strToBytes)).to.be.reverted
    })
    it("🟢 Should call checkUserHavePassedTTP_ByHash", async () => {
        await passport.createUser(user.address, 1, strToBytes)
        await passport.connect(user).checkUserHavePassedTTP_ByHash(strToBytes)
    })
    it("🟡 Should NOT call switchTTP under regular user", async () => {
        await expect(passport.connect(user).switchTTP(ethers.constants.AddressZero)).to.be.reverted
    })
    it("🟢 Should call switchTTP", async () => {
        await passport.connect(owner).switchTTP(ethers.constants.AddressZero)
    })
    it("🟢 Should call switchTTP to switch it twice", async () => {
        await passport.connect(owner).switchTTP(ethers.constants.AddressZero)
        await passport.connect(owner).switchTTP(ethers.constants.AddressZero)
    })
    it("🟡 Should NOT call pauseTTP under regular user", async () => {
        await expect(passport.connect(user).pauseTTP(ethers.constants.AddressZero)).to.be.reverted
    })
    it("🟢 Should call pauseTTP", async () => {
        await passport.connect(owner).pauseTTP(ethers.constants.AddressZero)
    })
    it("🟢 Should call CheckUserHaveTypeIDByAddr without user creation", async () => {
        await passport.connect(owner).CheckUserHaveTypeIDByAddr(user.address, 0)
    })
    it("🟢 Should call CheckUserHaveTypeIDByAddr with user creation", async () => {
        await passport.createUser(user.address, 0, strToBytes)
        await passport.connect(owner).CheckUserHaveTypeIDByAddr(user.address, 0)
    })
})
