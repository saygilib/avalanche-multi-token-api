const CustomERC20 = artifacts.require("CustomERC20");

module.exports = async function (deployer, network, accounts) {
  const name = "CustomToken";
  const symbol = "CTK";
  const initialSupply = web3.utils.toWei('1000', 'ether');
  const initialOwner = accounts[0]; 

  await deployer.deploy(CustomERC20, name, symbol, initialSupply, initialOwner);
};
