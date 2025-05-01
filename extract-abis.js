const fs = require('fs');
const path = require('path');

const contracts = ['CustomERC20', 'CustomERC721', 'CustomERC1155'];

const inputDir = path.join(__dirname, 'build', 'contracts');
const outputDir = path.join(__dirname, 'go-backend', 'abi');

// Make sure output directory exists
if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
}

contracts.forEach(contract => {
    const inputPath = path.join(inputDir, `${contract}.json`);
    const outputPath = path.join(outputDir, `${contract}.abi`);

    if (!fs.existsSync(inputPath)) {
        console.error(`ABI source file not found for ${contract}`);
        return;
    }

    const json = JSON.parse(fs.readFileSync(inputPath));
    const abi = JSON.stringify(json.abi, null, 2);

    fs.writeFileSync(outputPath, abi);
    console.log(`ABI extracted for ${contract} -> ${outputPath}`);
});

console.log("🎉 All ABIs extracted successfully.");
