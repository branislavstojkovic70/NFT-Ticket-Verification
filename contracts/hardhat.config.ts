import type { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox-viem";
import * as dotenv from "dotenv";

dotenv.config({path: __dirname + "/.env"});

const config: HardhatUserConfig = {
  solidity: "0.8.30",
  networks: {
    sepolia: {
      url: process.env.NETWORK_URL,
      chainId: parseInt(process.env.CHAIN_ID!),
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
    },
  },
};

export default config;
