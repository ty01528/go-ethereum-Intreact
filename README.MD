# 创建代币并使用Geth与合约交互

----
### 结构说明：

    - oceCoin.sol             # 创建了一个名为 oceCoin 的代币

    - call.abi                # Remix在合约部署后生成的abi文件

    - interact --|-- callContract.go   # abigen工具生成的原abi文件的go版本
                 |
                 |-- key&val           # 定义了网络信息与合约信息
                 |
                 |-- ...               # 合约交互文件

### 代币部署流程

- 1.将代码复制到remix,按照个人需求修改名称，精度，等
- 2.选中文件并编译。
- 3.连接钱包并部署，在environment中选择wallet Connect进行连接。Deploy进行部署。
- 4.获取部署后的合约。
- 5.部署完成后编辑main.go。
- 6.调用代码并运行。