require('dotenv').config();
const HDWalletProvider = require('@truffle/hdwallet-provider');

module.exports = {
  networks: {
    fuji: {
      provider: () =>
        new HDWalletProvider(process.env.PRIVATE_KEY, process.env.RPC_URL),
      network_id: 43113,
      gas: 5000000,
      gasPrice: 10000000000,
      timeoutBlocks: 200,
      skipDryRun: true
    }
  },
  compilers: {
    solc: {
      version: "0.8.21"
    }
  }
};
