import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
        { version: "0.8.0", settings: {} }, // Add this line.
        { version: "0.8.20", settings: {} },
        { version: "0.8.21", settings: {} },
        { version: "0.8.23", settings: {} },
    ],
}
};

export default config;
