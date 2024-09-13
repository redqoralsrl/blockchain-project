// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IIvoryFactory {
    event PairCreated(address indexed token0, address indexed token1, address pair, uint);

    function feeTo() external view returns (address);
    function feeToSetter() external view returns (address);

    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);

    function createPair(address tokenA, address tokenB) external returns (address pair);

    function setFeeTo(address) external;
    function setFeeToSetter(address) external;
}

abstract contract Factory is IIvoryFactory {
    bytes32 public INIT_CODE_PAIR_HASH;
    address public feeTo;
    address public feeToSetter;

    mapping(address => mapping(address => address)) public getPair;
    address[] public allPairs;

    function allPairsLength() external view virtual override returns (uint);

    function createPair(address tokenA, address tokenB) external virtual override returns (address pair);

    function setFeeTo(address _feeTo) external virtual override;

    function setFeeToSetter(address _feeToSetter) external virtual override;
}