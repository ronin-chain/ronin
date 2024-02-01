package shadow

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type validatorSetStorageSlot struct {
	// _candidateIds: []address
	CandidateIdsSlot uint64 `json:"candidateIds_slot" yaml:"candidateIds_slot"`
	// _candidateIndex: mapping[address] => uint256
	CandidateIndexSlot uint64 `json:"candidate_index_slot" yaml:"candidate_index_slot"`
	// _candidateInfo: mapping[address] => ValidatorCandidate
	CandidateInfoSlot uint64 `json:"candidate_info_slot" yaml:"candidate_info_slot"`
	// _validatorMap: mapping[address] => ValidatorFlag
	ValidatorMapSlot uint64 `json:"validator_map_slot" yaml:"validator_map_slot"`
	// _validatorIds: mapping[uint256] => address
	ValidatorIdsSlot uint64 `json:"validatorIds_slot" yaml:"validatorIds_slot"`
	// _validatorCount: uint256
	ValidatorCountSlot       uint64                    `json:"validator_count_slot" yaml:"validator_count_slot"`
	ValidatorCandidateOffset *validatorCandidateOffset `json:"validator_candidate_offset" yaml:"validator_candidate_offset"`
}

type profileStorageSlot struct {
	// _id2Profile: mapping[address] => CandidateProfile
	Id2ProfileSlot uint64 `json:"id2profile_slot" yaml:"id2profile_slot"`
	// _consensus2Id: mapping[TConsensus] => address
	Consensus2IdSlot uint64 `json:"consensus2Id_slot" yaml:"consensus2Id_slot"`
	// _registry: mapping[address] => bool
	RegistrySlot           uint64                  `json:"registry_slot" yaml:"registry_slot"`
	CandidateProfileOffset *candidateProfileOffset `json:"candidate_profile_offset" yaml:"candidate_profile_offset"`
}

type stakingStorageSlot struct {
	// _poolDetail: mapping[address] => PoolDetail
	PoolDetailSlot uint64 `json:"pool_detail_slot" yaml:"pool_detail_slot"`
	// _stakingPool: mapping[address] => PoolFields
	StakingPoolSlot uint64 `json:"staking_pool_slot" yaml:"staking_pool_slot"`
	// _userReward: mapping[address] => UserRewardFields
	UserRewardSlot uint64 `json:"user_reward_slot" yaml:"user_reward_slot"`
	// _adminOfActivePoolMapping: mapping[address] => address
	AdminOfActivePoolMappingSlot uint64                  `json:"admin_of_active_pool_mapping_slot" yaml:"admin_of_active_pool_mapping_slot"`
	PoolDetailOffset             *poolDetailOffset       `json:"pool_detail_offset" yaml:"pool_detail_offset"`
	PoolFieldsOffset             *poolFieldsOffset       `json:"pool_fields_offset" yaml:"pool_fields_offset"`
	UserRewardFieldsOffset       *userRewardFieldsOffset `json:"user_reward_fields_offset" yaml:"user_reward_fields_offset"`
}

type trustedOrgSlot struct {
	// _consensusList: []TConsensus
	ConsensusListSlot uint64 `json:"consensus_list_slot" yaml:"consensus_list_slot"`
	// _governorList: []address
	GovernorListSlot uint64 `json:"governor_list_slot" yaml:"governor_list_slot"`
	// _addedBlock: mapping[address] => uint256
	AddedBlockSlot uint64 `json:"added_block_slot" yaml:"added_block_slot"`
	// _totalWeight: uint256
	TotalWeightSlot uint64 `json:"total_weight_slot" yaml:"total_weight_slot"`
	// _consensusWeight: mapping[TConsensus] => uint256
	ConsensusWeightSlot uint64 `json:"consensus_weight_slot" yaml:"consensus_weight_slot"`
	// _governorWeight: mapping[address] => uint256
	GovernorWeightSlot uint64 `json:"governor_weight_slot" yaml:"governor_weight_slot"`
}

type slashIndicatorSlot struct {
	// _creditScore: mapping[address]uint256
	CreditScoreSlot uint64 `json:"credit_score_slot" yaml:"credit_score_slot"`
}

type validatorCandidateOffset struct {
	AdminOffset          uint64 `json:"admin_offset" yaml:"admin_offset"`
	ConsensusAddrOffset  uint64 `json:"consensus_addr_offset" yaml:"consensus_addr_offset"`
	TreasuryAddrOffset   uint64 `json:"treasury_addr_offset" yaml:"treasury_addr_offset"`
	CommissionRateOffset uint64 `json:"commission_rate_offset" yaml:"commission_rate_offset"`
}

type candidateProfileOffset struct {
	IdOffset        uint64 `json:"id_offset" yaml:"id_offset"`
	ConsensusOffset uint64 `json:"consensus_offset" yaml:"consensus_offset"`
	AdminOffset     uint64 `json:"admin_offset" yaml:"admin_offset"`
	TreasuryOffset  uint64 `json:"treasury_offset" yaml:"treasury_offset"`
	GovernorOffset  uint64 `json:"governor_offset" yaml:"governor_offset"`
	PubkeyOffset    uint64 `json:"pubkey_offset" yaml:"pubkey_offset"`
}

type poolDetailOffset struct {
	PidOffset               uint64 `json:"pid_offset" yaml:"pid_offset"`
	ShadowedPoolAdminOffset uint64 `json:"shadowed_pool_admin_offset" yaml:"shadowed_pool_admin_offset"`
	StakingAmountOffset     uint64 `json:"staking_amount_offset" yaml:"staking_amount_offset"`
	StakingTotalOffset      uint64 `json:"staking_total_offset" yaml:"staking_total_offset"`
	// delegatingAmount: mapping[address] => uint256
	DelegatingAmountOffset uint64 `json:"delegating_amount_offset" yaml:"delegating_amount_offset"`
	// lastDelegatingTimestamp: mapping[address] => uint256
	LastDelegatingTimestampOffset uint64 `json:"last_delegating_timestamp_offset" yaml:"last_delegating_timestamp_offset"`
}

type poolFieldsOffset struct {
	SharesInnerOffset      uint64 `json:"shares_inner_offset" yaml:"shares_inner_offset"`
	SharesLastPeriodOffset uint64 `json:"shares_last_period_offset" yaml:"shares_last_period_offset"`
}

type userRewardFieldsOffset struct {
	LowestAmountOffset uint64 `json:"lowest_amount_offset" yaml:"lowest_amount_offset"`
	LastPeriodOffset   uint64 `json:"last_period_offset" yaml:"last_period_offset"`
}

type validatorCandidate struct {
	Admin          common.Hash
	ConsensusAddr  common.Hash
	TreasuryAddr   common.Hash
	CommissionRate common.Hash
}

type candidateProfile struct {
	Id        common.Hash
	Consensus common.Hash
	Admin     common.Hash
	Treasury  common.Hash
	Governor  common.Hash
	Pubkey    string
}

type poolFields struct {
	SharesInner      common.Hash
	SharesLastPeriod common.Hash
}

type userRewardFields struct {
	LowestAmount common.Hash
	LastPeriod   common.Hash
}

type poolDetail struct {
	Pid                     common.Hash
	ShadowedPoolAdmin       common.Hash
	StakingAmount           common.Hash
	StakingTotal            common.Hash
	DelegatingAmount        common.Hash
	LastDelegatingTimestamp common.Hash
}

type NewValidatorConfig struct {
	CandidateAdmin string   `json:"candidate_admin" yaml:"candidate_admin"`
	ConsensusAddr  string   `json:"consensus_addr" yaml:"consensus_addr"`
	TreasuryAddr   string   `json:"treasury_addr" yaml:"treasury_addr"`
	CommissionRate *big.Int `json:"commission_rate" yaml:"commission_rate"`
	Pubkey         string   `json:"pubkey" yaml:"pubkey"`
	GovernorAddr   string   `json:"governor_addr" yaml:"governor_addr"`
}

// newValidator includes the values, which is inserted into the computed storage slots
type newValidator struct {
	AdminAddr          common.Hash
	ConsensusAddr      common.Hash
	PubKey             string
	GovernorAddr       common.Hash
	Index              common.Hash
	ValidatorCandidate validatorCandidate
	CandidateProfile   candidateProfile
	PoolDetail         poolDetail
	PoolFields         poolFields
	UserRewardFields   userRewardFields
}

var (
	stake = common.HexToHash("0x33b2e3c9fd0803ce8000000") // equivalent to 10^8 eth
)

// NewValidators pre-processes inserted values
func NewValidators(cfg []*NewValidatorConfig, lastPeriod, lastDelegatingTimestamp uint64) []*newValidator {
	newVals := make([]*newValidator, 0)
	maxValue := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	for i, valCfg := range cfg {
		var val newValidator
		val.AdminAddr = common.HexToHash(valCfg.CandidateAdmin)
		val.ConsensusAddr = common.HexToHash(valCfg.ConsensusAddr)
		val.PubKey = valCfg.Pubkey
		val.GovernorAddr = common.HexToHash(valCfg.GovernorAddr)

		val.Index = common.BigToHash(new(big.Int).Sub(maxValue, big.NewInt(int64(i))))

		val.ValidatorCandidate.Admin = val.AdminAddr
		val.ValidatorCandidate.ConsensusAddr = val.ConsensusAddr
		val.ValidatorCandidate.TreasuryAddr = val.AdminAddr
		val.ValidatorCandidate.CommissionRate = common.BigToHash(valCfg.CommissionRate)

		val.CandidateProfile.Id = val.ConsensusAddr
		val.CandidateProfile.Consensus = val.ConsensusAddr
		val.CandidateProfile.Admin = val.AdminAddr
		val.CandidateProfile.Treasury = val.AdminAddr
		val.CandidateProfile.Governor = val.GovernorAddr
		val.CandidateProfile.Pubkey = val.PubKey

		val.PoolDetail.Pid = val.ConsensusAddr
		val.PoolDetail.ShadowedPoolAdmin = val.AdminAddr
		val.PoolDetail.StakingAmount = stake
		val.PoolDetail.StakingTotal = stake
		val.PoolDetail.DelegatingAmount = stake
		val.PoolDetail.LastDelegatingTimestamp = common.BigToHash(big.NewInt(int64(lastDelegatingTimestamp))) //hardfork - 3 day (in second)

		val.PoolFields.SharesInner = common.HexToHash("0x0")
		val.PoolFields.SharesLastPeriod = common.BigToHash(big.NewInt(int64(lastPeriod)))

		val.UserRewardFields.LowestAmount = common.HexToHash("0x0")
		val.UserRewardFields.LastPeriod = common.BigToHash(big.NewInt(int64(lastPeriod)))

		newVals = append(newVals, &val)
	}
	return newVals
}

type Wallet struct {
	WalletAddr common.Address `json:"wallet_addr" yaml:"wallet_addr"`
}
