// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract Pair {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);

    function name() external view virtual returns (string memory);
    function symbol() external view virtual returns (string memory);
    function decimals() external view virtual returns (uint8);
    function totalSupply() external view virtual returns (uint);
    function balanceOf(address owner) external view virtual returns (uint);
    function allowance(address owner, address spender) external view virtual returns (uint);

    function approve(address spender, uint value) external virtual returns (bool);
    function transfer(address to, uint value) external virtual returns (bool);
    function transferFrom(address from, address to, uint value) external virtual returns (bool);

    function DOMAIN_SEPARATOR() external view virtual returns (bytes32);
    function PERMIT_TYPEHASH() external view virtual returns (bytes32);
    function nonces(address owner) external view virtual returns (uint);

    function permit(address owner, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s) external virtual;

    event Mint(address indexed sender, uint amount0, uint amount1);
    event Burn(address indexed sender, uint amount0, uint amount1, address indexed to);
    event Swap(
        address indexed sender,
        uint amount0In,
        uint amount1In,
        uint amount0Out,
        uint amount1Out,
        address indexed to
    );
    event Sync(uint112 reserve0, uint112 reserve1);

    function MINIMUM_LIQUIDITY() external view virtual returns (uint);
    function factory() external view virtual returns (address);
    function token0() external view virtual returns (address);
    function token1() external view virtual returns (address);
    function getReserves() external view virtual returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    function price0CumulativeLast() external view virtual returns (uint);
    function price1CumulativeLast() external view virtual returns (uint);
    function kLast() external view virtual returns (uint);

    function mint(address to) external virtual returns (uint liquidity);
    function burn(address to) external virtual returns (uint amount0, uint amount1);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external virtual;
    function skim(address to) external virtual;
    function sync() external virtual;

    function initialize(address, address) external virtual;
}