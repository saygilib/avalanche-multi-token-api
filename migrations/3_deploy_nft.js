const CustomERC721 = artifacts.require("CustomERC721");

module.exports = async function (deployer, network, accounts) {
  const owner = accounts[0];
  await deployer.deploy(CustomERC721, owner);
};
