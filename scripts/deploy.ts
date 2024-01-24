import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { ethers } from 'hardhat';

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log('Deploying contracts with the account:', deployer.address);

  // Deploy Passport contract
  const Passport = await ethers.getContractFactory('Passport');
  const passport = await Passport.deploy();

  await passport.deployed();

  console.log('Passport contract deployed to:', passport.address);

  // Deploy Vote contract with Passport address in its constructor
  const Vote = await ethers.getContractFactory('Vote');
  const vote = await Vote.deploy(passport.address); // Assuming Vote constructor accepts Passport address

  await vote.deployed();

  console.log('Vote contract deployed to:', vote.address);
}

// Run the deployment
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

