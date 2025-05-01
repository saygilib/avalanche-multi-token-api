// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract CustomERC1155 is ERC1155URIStorage, Ownable {
    constructor(address initialOwner)
        ERC1155("")
        Ownable(initialOwner)
    {
        // Initial mint of two token types with preset URIs
        _setURI(1, "https://example.com/metadata/gold.json");
        _setURI(2, "https://example.com/metadata/silver.json");

        _mint(initialOwner, 1, 100, "");
        _mint(initialOwner, 2, 250, "");
    }

    function mint(
        address to,
        uint256 id,
        uint256 amount,
        string memory uri
    ) external onlyOwner {
        _setURI(id, uri);
        _mint(to, id, amount, "");
    }
}
