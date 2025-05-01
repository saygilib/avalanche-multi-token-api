const CustomERC721 = artifacts.require("CustomERC721");

contract("CustomERC721", (accounts) => {
  const [owner, user1] = accounts;
  let nft;

  beforeEach(async () => {
    nft = await CustomERC721.new(owner);
  });

  it("should deploy with correct name and symbol", async () => {
    const name = await nft.name();
    const symbol = await nft.symbol();
    assert.equal(name, "MyCustomNFT");
    assert.equal(symbol, "MCN");
  });

  it("should mint a token with a URI to the given address", async () => {
    const tokenURI = "https://example.com/metadata/1.json";
    const tx = await nft.mint(user1, tokenURI, { from: owner });
    const tokenId = tx.logs[0].args.tokenId.toString();
    const ownerOf = await nft.ownerOf(tokenId);
    const storedUri = await nft.tokenURI(tokenId);
    assert.equal(ownerOf, user1);
    assert.equal(storedUri, tokenURI);
  });

  it("should not allow non-owner to mint", async () => {
    const tokenURI = "https://example.com/unauthorized.json";
    try {
      await nft.mint(user1, tokenURI, { from: user1 });
      assert.fail("Non-owner was able to mint");
    } catch (err) {
      assert.include(err.message, "revert", "Expected transaction to revert for non-owner mint");
    }
  });

  it("should increment token IDs with each mint", async () => {
    await nft.mint(user1, "https://uri1.com", { from: owner });
    await nft.mint(user1, "https://uri2.com", { from: owner });
    const owner1 = await nft.ownerOf(1);
    const owner2 = await nft.ownerOf(2);
    assert.equal(owner1, user1);
    assert.equal(owner2, user1);
  });
});
