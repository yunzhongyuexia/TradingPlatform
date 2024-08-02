//1_deploy.js

const MintToken = artifacts.require("MintToken");
// 上面的这个MintToken是编译后生成的文件名，也是上一步智能合约中contract MintToken is ERC20中的 MintToken，需要把下面的也改成这个
module.exports = async function (deployer) {
    // 设置重试次数和延迟时间
    const maxRetries = 5;
    const delay = 15000;

    // 部署 BToken 合约
    for (let attempt = 1; attempt <= maxRetries; attempt++) {
        try {
            console.log(`Attempt ${attempt} to deploy MintToken`);
            await deployer.deploy(MintToken);  //这个必须要跟上面一致
            console.log(`Successfully deployed MintToken on attempt ${attempt}`);
            break; // 如果部署成功，跳出循环
        } catch (error) {
            console.error(`Attempt ${attempt} failed: ${error.message}`);
            if (attempt < maxRetries) {
                console.log(`Waiting for ${delay / 1000} seconds before retrying...`);
                await new Promise(resolve => setTimeout(resolve, delay));
            } else {
                console.error('Max retries reached. Deployment failed.');
                throw error; // 达到最大重试次数后抛出错误
            }
        }
    }
};

