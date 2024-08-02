// truffle-config.js

const HDWalletProvider = require('@truffle/hdwallet-provider');

const mnemonic = "chat mean immune own rocket basic bargain lab error abandon guide pipe";

module.exports = {
    networks: {
        holesky: {
            provider: () => new HDWalletProvider(
                mnemonic,
                `https://ethereum-holesky.core.chainstack.com/5cb26e90d4925a7f24a1c7b51c2d8263`
            ),
            network_id: "*",  // holesky的链Id是17000，设置*表示适配所有的链
            confirmations: 2,
            timeoutBlocks: 200,
            skipDryRun: true,
            timeoutBlocks: 200,  // 等待区块确认的超时时间
            networkCheckTimeout: 1000000000,  // 网络检查的超时时间
        }
    },
    compilers: {
        solc: {
            version: "0.8.21"
        }
    }
};

