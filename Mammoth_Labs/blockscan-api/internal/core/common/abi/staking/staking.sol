// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.0;


abstract contract Staking {

    uint256 internal constant BALANCE_COMPACT_PRECISION = 1e10;
    uint16 internal constant COMMISSION_RATE_MIN_VALUE = 0; // 0%
    uint16 internal constant COMMISSION_RATE_MAX_VALUE = 3000; // 30%
    uint64 internal constant TRANSFER_GAS_LIMIT = 30000;
    uint256 public ECOSYSTEM_REWARD = 82191 ether;

    // validator events
    event ValidatorAdded(address indexed validator, address owner, uint8 status, uint16 commissionRate);
    event ValidatorModified(address indexed validator, address owner, uint8 status, uint16 commissionRate);
    event ValidatorRemoved(address indexed validator);
    event ValidatorOwnerClaimed(address indexed validator, uint256 amount, uint64 epoch);
    event ValidatorSlashed(address indexed validator, uint32 slashes, uint64 epoch);
    event ValidatorJailed(address indexed validator, uint64 epoch);
    event ValidatorDeposited(address indexed validator, uint256 amount, uint64 epoch);
    event ValidatorReleased(address indexed validator, uint64 epoch);

    // staker events
    event Delegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch);
    event Undelegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch);
    event Claimed(address indexed validator, address indexed staker, uint256 amount, uint64 epoch);
    event Redelegated(address indexed validator, address indexed staker, uint256 amount, uint256 dust, uint64 epoch);

    enum ValidatorStatus {
        NotFound,
        Active,
        Pending,
        Jail
    }

    struct ValidatorSnapshot {
        uint96 totalRewards;
        uint112 totalDelegated;
        uint32 slashesCount;
        uint16 commissionRate;
    }

    struct Validator {
        address validatorAddress;
        address ownerAddress;
        ValidatorStatus status;
        uint64 changedAt;
        uint64 jailedBefore;
        uint64 claimedAt;
    }

    struct DelegationOpDelegate {
        uint112 amount;
        uint64 epoch;
    }

    struct DelegationOpUndelegate {
        uint112 amount;
        uint64 epoch;
    }

    struct ValidatorDelegation {
        DelegationOpDelegate[] delegateQueue;
        uint64 delegateGap;
        DelegationOpUndelegate[] undelegateQueue;
        uint64 undelegateGap;
    }

    // mapping from validator address to validator
    mapping(address => Validator) internal _validatorsMap;
    // mapping from validator owner to validator address
    mapping(address => address) internal _validatorOwners;
    // list of all validators that are in validators mapping
    address[] internal _activeValidatorsList;
    // mapping with stakers to validators at epoch (validator -> delegator -> delegation)
    mapping(address => mapping(address => ValidatorDelegation)) internal _validatorDelegations;
    // mapping with validator snapshots per each epoch (validator -> epoch -> snapshot)
    mapping(address => mapping(uint64 => ValidatorSnapshot)) internal _validatorSnapshots;

    function ctor(address[] calldata validators, uint256[] calldata initialStakes, uint16 commissionRate) external virtual;

    function getValidatorDelegation(address validatorAddress, address delegator) external view virtual  returns (uint256 delegatedAmount, uint64 atEpoch);

    function getValidatorStatus(address validatorAddress) external view virtual  returns (
        address ownerAddress,
        uint8 status,
        uint256 totalDelegated,
        uint32 slashesCount,
        uint64 changedAt,
        uint64 jailedBefore,
        uint64 claimedAt,
        uint16 commissionRate,
        uint96 totalRewards
    );

    function getValidatorStatusAtEpoch(address validatorAddress, uint64 epoch) external view virtual returns (
        address ownerAddress,
        uint8 status,
        uint256 totalDelegated,
        uint32 slashesCount,
        uint64 changedAt,
        uint64 jailedBefore,
        uint64 claimedAt,
        uint16 commissionRate,
        uint96 totalRewards
    );

    function getValidatorByOwner(address owner) external view virtual  returns (address);

    function releaseValidatorFromJail(address validatorAddress) external virtual;

    function delegate(address validatorAddress) payable external virtual ;

    function undelegate(address validatorAddress, uint256 amount) external virtual ;

    function currentEpoch() external view virtual returns (uint64);

    function nextEpoch() external view virtual returns (uint64);

    function registerValidator(address validatorAddress, uint16 commissionRate) payable external virtual ;

    function addValidator(address account) external virtual ;

    function setEcosystemReward(uint256 _amount) external virtual;

    function removeValidator(address account) external virtual ;

    function activateValidator(address validator) external virtual ;

    function disableValidator(address validator) external virtual ;

    function changeValidatorCommissionRate(address validatorAddress, uint16 commissionRate) external virtual;

    function changeValidatorOwner(address validatorAddress, address newOwner) external virtual ;

    function isValidatorActive(address account) external view virtual  returns (bool);

    function isValidator(address account) external view virtual  returns (bool);

    function getValidators() external view virtual  returns (address[] memory);

    function deposit(address validatorAddress) external payable virtual ;

    function getValidatorFee(address validatorAddress) external view virtual  returns (uint256);

    function getPendingValidatorFee(address validatorAddress) external view virtual  returns (uint256);

    function claimValidatorFee(address validatorAddress) external virtual ;

    function claimValidatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) external virtual ;

    function getDelegatorFee(address validatorAddress, address delegatorAddress) external view virtual  returns (uint256);

    function getPendingDelegatorFee(address validatorAddress, address delegatorAddress) external view virtual  returns (uint256);

    function claimDelegatorFee(address validatorAddress) external virtual ;

    function calcAvailableForRedelegateAmount(address validator, address delegator) external view virtual  returns (uint256 amountToStake, uint256 rewardsDust);

    function redelegateDelegatorFee(address validator) external virtual ;

    function claimDelegatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) external virtual ;

    function slash(address validatorAddress) external virtual ;
}
