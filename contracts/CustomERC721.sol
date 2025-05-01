// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract CustomERC721 is ERC721URIStorage, Ownable {
    uint256 private _tokenIdCounter;

    constructor(address initialOwner)
        ERC721("MyCustomNFT", "MCN")
        Ownable(initialOwner)
    {}

    function mint(address to, string memory uri) public onlyOwner returns (uint256) {
        _tokenIdCounter++;
        uint256 tokenId = _tokenIdCounter;
        _mint(to, tokenId);
        _setTokenURI(tokenId, uri);
        return tokenId;
    }
}
