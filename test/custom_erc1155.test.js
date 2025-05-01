const CustomERC1155 = artifacts.require("CustomERC1155");

contract("CustomERC1155", (accounts) => {
  const [owner, user1] = accounts;
  let token;

  beforeEach(async () => {
    token = await CustomERC1155.new(owner);
  });

  it("should deploy and mint initial supply for token IDs 1 and 2", async () => {
    const balance1 = await token.balanceOf(owner, 1);
    const balance2 = await token.balanceOf(owner, 2);
    assert(balance1.toNumber() > 0, "Token ID 1 should have been minted");
    assert(balance2.toNumber() > 0, "Token ID 2 should have been minted");
  });

  it("should allow the owner to mint a new token with URI", async () => {
    const tokenId = 3;
    const amount = 10;
    const uri = "https://example.com/metadata/bronze.json";
    await token.mint(user1, tokenId, amount, uri, { from: owner });
    const balance = await token.balanceOf(user1, tokenId);
    const storedUri = await token.uri(tokenId);
    assert.equal(balance.toNumber(), amount);
    assert.equal(storedUri, uri);
  });

  it("should not allow non-owner to mint", async () => {
    const tokenId = 4;
    const amount = 5;
    const uri = "https://fail.com";
    try {
      await token.mint(user1, tokenId, amount, uri, { from: user1 });
      assert.fail("Non-owner was able to mint");
    } catch (err) {
      assert.include(err.message, "revert", "Expected transaction to revert for non-owner mint");
    }
  });

  it("should overwrite URI if token ID is minted again", async () => {
    const tokenId = 5;
    const uri1 = "https://example.com/one.json";
    const uri2 = "https://example.com/two.json";
    await token.mint(user1, tokenId, 1, uri1, { from: owner });
    await token.mint(user1, tokenId, 1, uri2, { from: owner });
    const uri = await token.uri(tokenId);
    assert.equal(uri, uri2);
  });
});
