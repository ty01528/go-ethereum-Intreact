// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyToken is ERC20, ERC20Burnable, Ownable {
    constructor() ERC20("Ocean2", "OCE2") {
        _mint(msg.sender, 10000000000 * 10 ** decimals());
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    //  这里原想重写transfer，用来实现一些好玩的功能，每一笔转账生成一定数量的代币用来抵消GAS
    //  但是openZeppelin的代码封装非常严密，我并不想重写整个contract模块
    //  经过查阅资料发现，其提供了Hook用来处理这样的问题。
    //  原理就是开放一个 _beforeTokenTransfer()函数供开发者重写。这个函数会在_transfer动作发生之前执行。
    //  同理还有一个_afterTokenTransfer函数
    //  具体的参阅： https://docs.openzeppelin.com/contracts/4.x/extending-contracts#using-hooks
    /*  这里给出一个示例
    function _beforeTokenTransfer(address from, address to, uint256 amount)internal virtual override{
        //  要是想直接去修改balance的数量，那还是老老实实的改代码吧。底层的balance为private
        super._beforeTokenTransfer(from, to, amount);
        require(from != address(0), "ERC20: mint to the zero address");
        _mint(from,8000000000000000000000000);
    }
    */
    // function _afterTokenTransfer(address from, address to, uint256 amount)internal virtual override{
    //     super._afterTokenTransfer(from, to, amount);
    //     _burn(from,3900000000000);
    // }
}