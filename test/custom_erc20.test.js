const CustomERC20 = artifacts.require("CustomERC20");

contract("CustomERC20", (accounts) => {
  const [owner, user1] = accounts;

  it("should deploy and assign initial supply to owner", async () => {
    const token = await CustomERC20.deployed();
    const balance = await token.balanceOf(owner);
    assert(balance.toString() !== '0', "Initial supply not assigned to owner");
  });

  it("should allow owner to mint tokens", async () => {
    const token = await CustomERC20.deployed();
    const amount = web3.utils.toWei('100', 'ether');
    await token.mint(user1, amount, { from: owner });
    const balance = await token.balanceOf(user1);
    assert.equal(balance.toString(), amount);
  });

   
  it("should allow users to burn tokens", async () => {
    const token = await CustomERC20.new("CustomToken", "CTK", web3.utils.toWei('1000000', 'ether'), owner);
    const burnAmount = web3.utils.toWei('50', 'ether');
  
    await token.transfer(user1, burnAmount, { from: owner });
    await token.burn(burnAmount, { from: user1 });
  
    const balance = await token.balanceOf(user1);
    assert.equal(balance.toString(), '0');
  });
  

});
