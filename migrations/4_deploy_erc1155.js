const CustomERC1155 = artifacts.require("CustomERC1155");

module.exports = async function (deployer, network, accounts) {
  const owner = accounts[0];
  await deployer.deploy(CustomERC1155, owner);
};
