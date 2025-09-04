package shadow

import (
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/yaml.v2"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ShadowSwitchConfig  Config
	NumberOfTrustedOrgs int64
)

func init() {
	var configPath string
	configPath = os.Getenv("SHADOW_FORK_CONFIG_PATH")
	if configPath == "" {
		configPath = "config.mainnet10.yaml"
		log.Warn("get the default shadow fork config path", "path", configPath)
	}

	log.Info("Shadow fork config", "path", configPath)

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Error("[NewShadowSwitch] Fail to read shadow switch config file", "err", err)
		return
	}

	if err := yaml.Unmarshal(yamlFile, &ShadowSwitchConfig); err != nil {
		log.Error("[NewShadowSwitch] fail to unmarshal shadow switch config", "err", err)
		return
	}

	// Number of trusted orgs must be equal to the number of new validators
	NumberOfTrustedOrgs = int64(len(ShadowSwitchConfig.NewValidatorConfigs))
}

type shadowSwitch struct {
	NewValidators         []*newValidator
	ValidatorSetAddress   common.Address
	ProfileAddress        common.Address
	StakingAddress        common.Address
	TrustedOrgAddress     common.Address
	SlashIndicatorAddress common.Address

	state                   *state.StateDB
	ValidatorSetStorageSlot *validatorSetStorageSlot
	ProfileStoragesSlot     *profileStorageSlot
	StakingStorageSlot      *stakingStorageSlot
	TrustedOrgSlot          *trustedOrgSlot
	SlashIndicatorSlot      *slashIndicatorSlot

	candidateConsensusAddr []common.Address
	candidateGovernorAddr  []common.Address

	NewBlankWallets    []Wallet
	BlankWalletBalance *big.Int
}

type Config struct {
	ValidatorSetStorageSlot *validatorSetStorageSlot `json:"validator_set_storage_slot" yaml:"validator_set_storage_slot"`
	ProfileStorageSlot      *profileStorageSlot      `json:"profile_storage_slot" yaml:"profile_storage_slot"`
	StakingStorageSlot      *stakingStorageSlot      `json:"staking_storage_slot" yaml:"staking_storage_slot"`
	TrustedOrgSlot          *trustedOrgSlot          `json:"trusted_org_slot" yaml:"trusted_org_slot"`
	SlashIndicatorSlot      *slashIndicatorSlot      `json:"slash_indicator_slot" yaml:"slash_indicator_slot"`

	NewValidatorConfigs   []*NewValidatorConfig `json:"new_validator_configs" yaml:"new_validator_configs"`
	ValidatorSetAddress   common.Address        `json:"validator_set_address" yaml:"validator_set_address"`
	ProfileAddress        common.Address        `json:"profile_address" yaml:"profile_address"`
	StakingAddress        common.Address        `json:"staking_address" yaml:"staking_address"`
	TrustedOrgAddress     common.Address        `json:"trusted_org_address" yaml:"trusted_org_address"`
	SlashIndicatorAddress common.Address        `json:"slash_indicator_address" yaml:"slash_indicator_address"`

	LastPeriod              uint64 `json:"last_period" yaml:"last_period"`
	LastDelegatingTimestamp uint64 `json:"last_delegating_timestamp" yaml:"last_delegating_timestamp"`

	NewBlankWallets    []Wallet `json:"new_blank_wallets" yaml:"new_blank_wallets"`
	BlankWalletBalance *big.Int `json:"blank_wallet_balance" yaml:"blank_wallet_balance"`
	ChainID            uint64   `json:"chain_id" yaml:"chain_id"`
}

func NewShadowSwitch(state *state.StateDB, candidateConsensusAddr, candidateGovernorAddr []common.Address) (*shadowSwitch, error) {
	newValidators := NewValidators(ShadowSwitchConfig.NewValidatorConfigs, ShadowSwitchConfig.LastPeriod, ShadowSwitchConfig.LastDelegatingTimestamp)

	return &shadowSwitch{
		NewValidators:         newValidators,
		ValidatorSetAddress:   ShadowSwitchConfig.ValidatorSetAddress,
		ProfileAddress:        ShadowSwitchConfig.ProfileAddress,
		StakingAddress:        ShadowSwitchConfig.StakingAddress,
		TrustedOrgAddress:     ShadowSwitchConfig.TrustedOrgAddress,
		SlashIndicatorAddress: ShadowSwitchConfig.SlashIndicatorAddress,

		state: state,

		ValidatorSetStorageSlot: ShadowSwitchConfig.ValidatorSetStorageSlot,
		ProfileStoragesSlot:     ShadowSwitchConfig.ProfileStorageSlot,
		StakingStorageSlot:      ShadowSwitchConfig.StakingStorageSlot,
		TrustedOrgSlot:          ShadowSwitchConfig.TrustedOrgSlot,
		SlashIndicatorSlot:      ShadowSwitchConfig.SlashIndicatorSlot,

		candidateConsensusAddr: candidateConsensusAddr,
		candidateGovernorAddr:  candidateGovernorAddr,

		// Create new blank wallets and add balance to them.
		NewBlankWallets:    ShadowSwitchConfig.NewBlankWallets,
		BlankWalletBalance: ShadowSwitchConfig.BlankWalletBalance,
	}, nil
}

// Run function presents logic for validator switching during shadow fork.
// It changes the storage slot of 5 contracts (ValidatorSet, Profile, Staking, TrustedOrg, SlashIndicator)
// Ref: https://github.com/axieinfinity/ronin-dpos-contracts
func Run(state *state.StateDB, candidateConsensusAddr, candidateGovernorAddr []common.Address) error {
	shadowSwitch, err := NewShadowSwitch(state, candidateConsensusAddr, candidateGovernorAddr)
	if err != nil {
		log.Debug("fail to process shadow switch", "err", err)
		return err
	}

	shadowSwitch.updateValidatorSet()
	shadowSwitch.updateProfile()
	shadowSwitch.updateStaking()
	shadowSwitch.updateTrustedOrg()
	shadowSwitch.updateSlashIndicator()

	shadowSwitch.createBlankWallets() // update balance for blank wallets

	return nil
}

// UpdateValidatorSet modifies the storage of RoninValidatorSet contract
func (s *shadowSwitch) updateValidatorSet() {
	newSize := big.NewInt(int64(len(s.NewValidators)))

	// update the size of _candidateIds: []address
	candidateIdsLoc := state.GetLocSimpleVariable(s.ValidatorSetStorageSlot.CandidateIdsSlot)
	s.state.SetState(s.ValidatorSetAddress, candidateIdsLoc, common.BigToHash(newSize))

	// update the size of _validatorIds: []address
	validatorIdsLoc := state.GetLocSimpleVariable(s.ValidatorSetStorageSlot.ValidatorIdsSlot)
	s.state.SetState(s.ValidatorSetAddress, validatorIdsLoc, common.BigToHash(newSize))

	// update _validatorCount: uint256
	cLoc := state.GetLocSimpleVariable(s.ValidatorSetStorageSlot.ValidatorCountSlot)
	s.state.SetState(s.ValidatorSetAddress, cLoc, common.BigToHash(newSize))

	// remove all current records in _candidateIds, _candidateIndex, _candidateInfo, _validatorIds
	for i, cval := range s.candidateConsensusAddr {
		// remove _candidateIds
		iloc := state.GetLocDynamicArrAtElement(candidateIdsLoc, uint64(i), 1)
		s.state.SetState(s.ValidatorSetAddress, iloc, common.Hash{})
		// remove _candidateIndex
		indexLoc := state.GetLocMappingAtKey(cval.Hash(), s.ValidatorSetStorageSlot.CandidateIndexSlot)
		s.state.SetState(s.ValidatorSetAddress, indexLoc, common.Hash{})
		// remove _candidateInfo
		infoLoc := state.GetLocMappingAtKey(cval.Hash(), s.ValidatorSetStorageSlot.CandidateInfoSlot)
		s.updateValidatorCandidate(infoLoc, &validatorCandidate{})
		// remove _validatorMap
		mLoc := state.GetLocMappingAtKey(cval.Hash(), s.ValidatorSetStorageSlot.ValidatorMapSlot)
		s.state.SetState(s.ValidatorSetAddress, mLoc, common.Hash{})
		// remove _validatorIds
		vLoc := state.GetLocMappingAtKey(common.BigToHash(big.NewInt(int64(i))), s.ValidatorSetStorageSlot.ValidatorIdsSlot)
		s.state.SetState(s.ValidatorSetAddress, vLoc, common.Hash{})
	}

	amount := new(big.Int).Mul(big.NewInt(1e18), s.BlankWalletBalance)
	// inserting the new validators
	for i, val := range s.NewValidators {
		// add balance for consensus address
		s.state.AddBalance(common.HexToAddress(val.ConsensusAddr.Hex()), amount, tracing.BalanceChangeUnspecified)
		// update _candidateIds: []address
		iloc := state.GetLocDynamicArrAtElement(candidateIdsLoc, uint64(i), 1)
		s.state.SetState(s.ValidatorSetAddress, iloc, val.ConsensusAddr)
		// update _candidateIndex: mapping[address] => uint256
		indexLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ValidatorSetStorageSlot.CandidateIndexSlot)
		s.state.SetState(s.ValidatorSetAddress, indexLoc, val.Index)
		// update _candidateInfo: mapping[address] => ValidatorCandidate
		infoLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ValidatorSetStorageSlot.CandidateInfoSlot)
		s.updateValidatorCandidate(infoLoc, &val.ValidatorCandidate)
		// update _validatorMap: mapping[address] => ValidatorFlag
		mLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ValidatorSetStorageSlot.ValidatorMapSlot)
		s.state.SetState(s.ValidatorSetAddress, mLoc, common.HexToHash("0x3"))
		// update _validatorIds: mapping[uint256] => address
		vLoc := state.GetLocMappingAtKey(common.BigToHash(big.NewInt(int64(i))), s.ValidatorSetStorageSlot.ValidatorIdsSlot)
		s.state.SetState(s.ValidatorSetAddress, vLoc, val.ConsensusAddr)
	}
}

// UpdateProfile modifies the storage of Profile contract
func (s *shadowSwitch) updateProfile() {
	amount := new(big.Int).Mul(big.NewInt(1e18), s.BlankWalletBalance)
	for _, val := range s.NewValidators {
		// add balance for admin address
		s.state.AddBalance(common.HexToAddress(val.AdminAddr.Hex()), amount, tracing.BalanceChangeUnspecified)
		// update _id2Profile: mapping[address] => CandidateProfile
		profileLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ProfileStoragesSlot.Id2ProfileSlot)
		s.updateCandidateProfile(profileLoc, &val.CandidateProfile)

		// update _consensus2Id: mapping[TConsensus] => address
		consensusLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ProfileStoragesSlot.Consensus2IdSlot)
		s.state.SetState(s.ProfileAddress, consensusLoc, val.ConsensusAddr)

		// update _registry: mapping[address] => bool
		registryAdminLoc := state.GetLocMappingAtKey(val.AdminAddr, s.ProfileStoragesSlot.RegistrySlot)
		s.state.SetState(s.ProfileAddress, registryAdminLoc, common.HexToHash("0x1"))
		registryConsensusLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ProfileStoragesSlot.RegistrySlot)
		s.state.SetState(s.ProfileAddress, registryConsensusLoc, common.HexToHash("0x1"))
	}
}

// UpdateStaking modifies the storage of Staking contract
func (s *shadowSwitch) updateStaking() {
	for _, val := range s.NewValidators {
		// update _poolDetail: mapping[address] => PoolDetail
		poolDetailLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.StakingStorageSlot.PoolDetailSlot)
		s.updatePoolDetail(poolDetailLoc, val, &val.PoolDetail)

		// update _stakingPool: mapping[address] => PoolFields
		poolFieldsLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.StakingStorageSlot.StakingPoolSlot)
		s.updatePoolFields(poolFieldsLoc, &val.PoolFields)

		// update _userReward: mapping[address]mapping[address] => UserRewardFields
		userLoc := getLocMappingAtKey(val.AdminAddr, state.GetLocMappingAtKey(val.ConsensusAddr, s.StakingStorageSlot.UserRewardSlot))
		s.updateUserRewardFields(userLoc, &val.UserRewardFields)

		// update _adminOfActivePoolMapping: mapping[address] => address
		adminLoc := state.GetLocMappingAtKey(val.AdminAddr, s.StakingStorageSlot.AdminOfActivePoolMappingSlot)
		s.state.SetState(s.StakingAddress, adminLoc, val.ConsensusAddr)
	}
}

// UpdateTrustedOrg modifies the storage slot of RoninTrustedOrganization contract
func (s *shadowSwitch) updateTrustedOrg() {
	newSize := big.NewInt(NumberOfTrustedOrgs)

	// update the size of _consensusList
	consensusLoc := state.GetLocSimpleVariable(s.TrustedOrgSlot.ConsensusListSlot)
	s.state.SetState(s.TrustedOrgAddress, consensusLoc, common.BigToHash(newSize))

	// update the size of _governorList
	governorLoc := state.GetLocSimpleVariable(s.TrustedOrgSlot.GovernorListSlot)
	s.state.SetState(s.TrustedOrgAddress, governorLoc, common.BigToHash(newSize))

	// update _totalWeight: uint256
	totalLoc := state.GetLocSimpleVariable(s.TrustedOrgSlot.TotalWeightSlot)
	s.state.SetState(s.TrustedOrgAddress, totalLoc, common.BigToHash(big.NewInt(int64(100*NumberOfTrustedOrgs))))

	// remove all records in _consensusList, _governorList, _consensusWeight, _governorWeight
	for i, val := range s.candidateConsensusAddr {
		// remove _consensusList
		cLoc := state.GetLocDynamicArrAtElement(consensusLoc, uint64(i), 1)
		s.state.SetState(s.TrustedOrgAddress, cLoc, common.Hash{})
		// remove _governorList
		gLoc := state.GetLocDynamicArrAtElement(governorLoc, uint64(i), 1)
		s.state.SetState(s.TrustedOrgAddress, gLoc, common.Hash{})
		// remove _addedBlock
		aLoc := state.GetLocMappingAtKey(val.Hash(), s.TrustedOrgSlot.AddedBlockSlot)
		s.state.SetState(s.TrustedOrgAddress, aLoc, common.Hash{})
		// remove _consensusWeight
		cwLoc := state.GetLocMappingAtKey(val.Hash(), s.TrustedOrgSlot.ConsensusWeightSlot)
		s.state.SetState(s.TrustedOrgAddress, cwLoc, common.Hash{})
	}

	// remove _governorWeight
	for _, gov := range s.candidateGovernorAddr {
		gwLoc := state.GetLocMappingAtKey(gov.Hash(), s.TrustedOrgSlot.GovernorWeightSlot)
		s.state.SetState(s.TrustedOrgAddress, gwLoc, common.Hash{})
	}

	amount := new(big.Int).Mul(big.NewInt(1e18), s.BlankWalletBalance)
	// add new org to trusted orgs
	for i, val := range s.NewValidators {
		// only update new trusted orgs with the first 6 validators.
		if i < int(NumberOfTrustedOrgs) {
			// add balance for governor address
			s.state.AddBalance(common.HexToAddress(val.GovernorAddr.Hex()), amount, tracing.BalanceChangeUnspecified)
			// update _consensusList: []TConsensus
			cLoc := state.GetLocDynamicArrAtElement(consensusLoc, uint64(i), 1)
			s.state.SetState(s.TrustedOrgAddress, cLoc, val.ConsensusAddr)
			// update _governorList: []address
			gLoc := state.GetLocDynamicArrAtElement(governorLoc, uint64(i), 1)
			s.state.SetState(s.TrustedOrgAddress, gLoc, val.GovernorAddr)
			// update _addedBlock: mapping[address]uint256
			aLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.TrustedOrgSlot.AddedBlockSlot)
			s.state.SetState(s.TrustedOrgAddress, aLoc, common.Hash{})
			// update _consensusWeight: mapping[TConsensus] => uint256
			cwLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.TrustedOrgSlot.ConsensusWeightSlot)
			s.state.SetState(s.TrustedOrgAddress, cwLoc, common.BigToHash(big.NewInt(100)))
			// update _governorWeight: mapping[address] => uint256
			gwLoc := state.GetLocMappingAtKey(val.GovernorAddr, s.TrustedOrgSlot.GovernorWeightSlot)
			s.state.SetState(s.TrustedOrgAddress, gwLoc, common.BigToHash(big.NewInt(int64(100))))
		}
	}
}

// updateSlashIndicator modifies the storage slot of SlashIndicator contract
func (s *shadowSwitch) updateSlashIndicator() {
	// update _creditScore: mapping[address]uint256
	for _, val := range s.NewValidators {
		cLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.SlashIndicatorSlot.CreditScoreSlot)
		s.state.SetState(s.SlashIndicatorAddress, cLoc, common.BigToHash(big.NewInt(int64(600))))
	}
}

func (s *shadowSwitch) createBlankWallets() {
	// We add a balance of `BlankWalletBalance` RON to each blank wallet.
	amount := new(big.Int).Mul(big.NewInt(1e18), s.BlankWalletBalance)
	for _, w := range s.NewBlankWallets {
		s.state.AddBalance(w.WalletAddr, amount, tracing.BalanceChangeUnspecified)
	}
}

func (s *shadowSwitch) updateValidatorCandidate(loc common.Hash, value *validatorCandidate) {
	validatorCandidateOffset := s.ValidatorSetStorageSlot.ValidatorCandidateOffset
	s.state.SetState(s.ValidatorSetAddress, state.GetLocStructElement(loc, validatorCandidateOffset.AdminOffset), value.Admin)
	s.state.SetState(s.ValidatorSetAddress, state.GetLocStructElement(loc, validatorCandidateOffset.ConsensusAddrOffset), value.ConsensusAddr)
	s.state.SetState(s.ValidatorSetAddress, state.GetLocStructElement(loc, validatorCandidateOffset.TreasuryAddrOffset), value.TreasuryAddr)
	s.state.SetState(s.ValidatorSetAddress, state.GetLocStructElement(loc, validatorCandidateOffset.CommissionRateOffset), value.CommissionRate)
}

func (s *shadowSwitch) updateCandidateProfile(loc common.Hash, value *candidateProfile) {
	candidateProfileOffset := s.ProfileStoragesSlot.CandidateProfileOffset
	s.state.SetState(s.ProfileAddress, state.GetLocStructElement(loc, candidateProfileOffset.IdOffset), value.Id)
	s.state.SetState(s.ProfileAddress, state.GetLocStructElement(loc, candidateProfileOffset.ConsensusOffset), value.Consensus)
	s.state.SetState(s.ProfileAddress, state.GetLocStructElement(loc, candidateProfileOffset.AdminOffset), value.Admin)
	s.state.SetState(s.ProfileAddress, state.GetLocStructElement(loc, candidateProfileOffset.TreasuryOffset), value.Treasury)
	s.state.SetState(s.ProfileAddress, state.GetLocStructElement(loc, candidateProfileOffset.GovernorOffset), value.Governor)
	// since bls pubkey is 48 bytes long, it requires 2 storage slot
	// store length * 2 + 1 (or "0x61") at storage slot and store each half of pubkey at keccak(storage_slot) + 0/1.
	// Ref: https://docs.soliditylang.org/en/latest/internals/layout_in_storage.html#bytes-and-string
	pubkeyLoc := state.GetLocStructElement(loc, candidateProfileOffset.PubkeyOffset)
	s.state.SetState(s.ProfileAddress, pubkeyLoc, common.HexToHash("0x61"))
	s.state.SetState(s.ProfileAddress, state.GetLocDynamicArrAtElement(pubkeyLoc, uint64(0), 1), common.HexToHash(value.Pubkey[:64]))
	s.state.SetState(s.ProfileAddress, state.GetLocDynamicArrAtElement(pubkeyLoc, uint64(1), 1), common.HexToHash(value.Pubkey[64:]+"00000000000000000000000000000000"))
}

func (s *shadowSwitch) updatePoolDetail(loc common.Hash, val *newValidator, value *poolDetail) {
	poolDetailOffset := s.StakingStorageSlot.PoolDetailOffset
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, poolDetailOffset.PidOffset), value.Pid)
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, poolDetailOffset.ShadowedPoolAdminOffset), value.ShadowedPoolAdmin)
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, poolDetailOffset.StakingAmountOffset), value.StakingAmount)
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, poolDetailOffset.StakingTotalOffset), value.StakingTotal)

	// update delegatingAmount: mapping[address] => uint256
	dloc := state.GetLocStructElement(loc, poolDetailOffset.DelegatingAmountOffset)
	s.state.SetState(s.StakingAddress, getLocMappingAtKey(val.AdminAddr, dloc), value.DelegatingAmount)

	// update lastDelegatingTimestamp: mapping[address] => uint256
	lloc := state.GetLocStructElement(loc, poolDetailOffset.LastDelegatingTimestampOffset)
	s.state.SetState(s.StakingAddress, getLocMappingAtKey(val.AdminAddr, lloc), value.LastDelegatingTimestamp)
}

func (s *shadowSwitch) updatePoolFields(loc common.Hash, value *poolFields) {
	poolFieldsOffset := s.StakingStorageSlot.PoolFieldsOffset
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, poolFieldsOffset.SharesInnerOffset), value.SharesInner)
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, poolFieldsOffset.SharesLastPeriodOffset), value.SharesLastPeriod)
}

func (s *shadowSwitch) updateUserRewardFields(loc common.Hash, value *userRewardFields) {
	userRewardOffset := s.StakingStorageSlot.UserRewardFieldsOffset
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, userRewardOffset.LowestAmountOffset), value.LowestAmount)
	s.state.SetState(s.StakingAddress, state.GetLocStructElement(loc, userRewardOffset.LastPeriodOffset), value.LastPeriod)
}

func getLocMappingAtKey(key common.Hash, slot common.Hash) common.Hash {
	retByte := crypto.Keccak256(key.Bytes(), slot.Bytes())
	ret := new(big.Int)
	ret.SetBytes(retByte)
	return common.BigToHash(ret)
}
