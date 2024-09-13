// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library MarketTypes {
    // Market hash

    // MarketPlace(bool isDealer,address edemSigner,address collection,uint256 tokenId,uint256 nftAmount,uint256 price,address protocolAddress,address tradeTokenAddress,uint256 nonce,uint256 startTime,uint256 endTime)
    bytes32 internal constant MARKET_PLACE_HASH =
    0x821b23108860b2262fa623d45a26073021dcb49a4ad5e89488e070794630832a;

    // Market 판매 목록 혹은 제안 목록
    struct MarketPlace {
        bool isDealer; // true -> NFT seller / false -> NFT proposer
        address edemSigner; // signer of the MarketPlace order
        address collection; // NFT contract Address
        uint256 tokenId; // NFT token Id
        uint256 nftAmount; // Must be 1 for ERC721, 1+ for ERC1155
        uint256 price; // NFT price 1ETH => 1000000000000000000
        address protocolAddress; // trade execution / protocol Fee
        address tradeTokenAddress; // curreny with WETH ...
        uint256 nonce; // New sales information is unique unless it ignores existing sales information
        uint256 startTime; // start timestamp
        uint256 endTime; // end timestamp
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    // Market 상품을 사는 사람 / 상품에 제안을 하는 사람
    struct User {
        bool isDealer; // true -> NFT seller / false -> NFT proposer
        address takerAddress; // msg.sender
        uint256 price; // purchase Amount
        uint256 tokenId; // NFT token Id
    }

    // hash
    function hash(
        MarketPlace memory marketPlace
    ) internal pure returns (bytes32) {
        return
            keccak256(
            abi.encode(
                MARKET_PLACE_HASH,
                marketPlace.isDealer,
                marketPlace.edemSigner,
                marketPlace.collection,
                marketPlace.tokenId,
                marketPlace.nftAmount,
                marketPlace.price,
                marketPlace.protocolAddress,
                marketPlace.tradeTokenAddress,
                marketPlace.nonce,
                marketPlace.startTime,
                marketPlace.endTime
            )
        );
    }
}

abstract contract Edem {
    address public immutable WETH;
    bytes32 public immutable DOMAIN_SEPARATOR;
    bytes4 public constant INTERFACE_ID_ERC721 = 0x80ac58cd;
    bytes4 public constant INTERFACE_ID_ERC1155 = 0xd9b67a26;
    address public protocoleFeeRecipient;

    mapping(address => uint256) userOrderNonce;

    event CancelAllOrders(address indexed user, uint256 newMinNonce);
    event CancelMultipleOrders(address indexed user, uint256[] orderNonces);
    event NewTradeTokenManager(address indexed tradeTokenManager);
    event NewProtocolExecutionManager(address indexed protocolExecutionManager);
    event NewProtocolFeeRecipient(address indexed protocoleFeeRecipient);
    event NewRoyaltyManager(address indexed royaltyManager);
    event NewTransferManager(address indexed transferManager);

    event RoyaltyPayment(
        address indexed collection,
        uint256 indexed tokenId,
        address indexed royaltyRecipient,
        address currency,
        uint256 amount
    );

    event UserSeller(
        bytes32 marketPlaceHash,
        uint256 orderNonce,
        address indexed userSeller,
        address indexed marketProposer,
        address indexed protocolAddress,
        address tradeTokenAddress,
        address collection,
        uint256 tokenId,
        uint256 nftAmount,
        uint256 price
    );

    event UserProposer(
        bytes32 marketPlaceHash,
        uint256 orderNonce,
        address indexed userProposer,
        address indexed marketSeller,
        address indexed protocolAddress,
        address tradeTokenAddress,
        address collection,
        uint256 tokenId,
        uint256 amount,
        uint256 price
    );

    function isUserOrderNonceExecutedOrCancelled(
        address user,
        uint256 orderNonces
    ) external view virtual returns (bool);
    function cancelAll(uint256 minNonce) virtual external;
    function cancelSellAndSuggest(uint256[] calldata orderNonces) external virtual;
    function proposerPayETH(
        MarketTypes.User calldata userProposer,
        MarketTypes.MarketPlace calldata marketSeller
    ) external virtual payable;
    function proposerPayETHAndWETH(
        MarketTypes.User calldata userProposer,
        MarketTypes.MarketPlace calldata marketSeller
    ) external virtual payable;
    function proposerPay(
        MarketTypes.User calldata userProposer,
        MarketTypes.MarketPlace calldata marketSeller
    ) external virtual;
    function suggestApprove(
        MarketTypes.User calldata userSeller,
        MarketTypes.MarketPlace calldata marketProposer
    ) external virtual;
}