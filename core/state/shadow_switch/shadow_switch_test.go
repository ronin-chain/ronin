package shadow

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestUnmarshalShadowConfig(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cfgPath := filepath.Join(cwd, "config/config.test.yaml")
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		fmt.Println("err", err)
		t.Fatal(err)
	}

	assert.Equal(t, common.HexToAddress("0x9687e8c41fa369ad08fd278a43114c4207856a61"), cfg.ValidatorSetAddress)
	assert.Equal(t, uint64(1), cfg.ProfileStorageSlot.CandidateProfileOffset.ConsensusOffset)
	assert.Equal(t, "0xE5635B0172e7488072348a0395C983DfDA68755A", cfg.NewValidatorConfigs[0].CandidateAdmin)
	assert.Equal(t, "a49891083c05e85a9023d20097367d5c203bfaa4b0f02753dbb85580508fbacee0fb7e7235c7fc032176d97e3e3b101c", cfg.NewValidatorConfigs[0].Pubkey)
	assert.Equal(t, 1, len(cfg.NewValidatorConfigs))
	assert.Equal(t, uint64(220), cfg.SlashIndicatorSlot.CreditScoreSlot)
	assert.Equal(t, common.HexToAddress("0xF7837778b6E180Df6696C8Fa986d62f8b6186752"), cfg.SlashIndicatorAddress)

	assert.Equal(t, 10, len(cfg.NewBlankWallets))
	assert.Equal(t, new(big.Int).SetUint64(1000), cfg.BlankWalletBalance)

	if len(cfg.NewBlankWallets) > 0 {
		assert.Equal(t, "0x4a6c63a216711A4ec59938b784223967A2206aeE", cfg.NewBlankWallets[0].WalletAddr.Hex())
	}
}

func TestNewValidator(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cfgPath := filepath.Join(cwd, "config/config.test.yaml")
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		fmt.Println("err", err)
		t.Fatal(err)
	}

	newVals := NewValidators(cfg.NewValidatorConfigs, 9999999, 9999999)

	val := newVals[0]
	assert.Equal(t, common.HexToHash("0xE5635B0172e7488072348a0395C983DfDA68755A"), val.AdminAddr)
	assert.Equal(t, common.HexToHash("0x13377331DEAD"), val.ConsensusAddr)
	assert.Equal(t, common.HexToHash("0x3e8"), val.ValidatorCandidate.CommissionRate)
	assert.Equal(t, "a49891083c05e85a9023d20097367d5c203bfaa4b0f02753dbb85580508fbacee0fb7e7235c7fc032176d97e3e3b101c", val.PubKey)
}

func TestShadowSwitchSlotCalculation(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cfgPath := filepath.Join(cwd, "config/config.test.yaml")
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		fmt.Println("err", err)
		t.Fatal(err)
	}

	newVals := NewValidators(cfg.NewValidatorConfigs, 9999999, 9999999)

	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)

	s := &shadowSwitch{
		NewValidators:       newVals,
		ValidatorSetAddress: cfg.ValidatorSetAddress,
		ProfileAddress:      cfg.ProfileAddress,
		StakingAddress:      cfg.StakingAddress,
		TrustedOrgAddress:   cfg.TrustedOrgAddress,
		state:               statedb,

		ValidatorSetStorageSlot: cfg.ValidatorSetStorageSlot,
		ProfileStoragesSlot:     cfg.ProfileStorageSlot,
		StakingStorageSlot:      cfg.StakingStorageSlot,
		TrustedOrgSlot:          cfg.TrustedOrgSlot,
	}

	// Test Validator Set
	candidateIdsLoc := state.GetLocSimpleVariable(s.ValidatorSetStorageSlot.CandidateIdsSlot)
	assert.Equal(t, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000073"), candidateIdsLoc)

	val := s.NewValidators[0]
	assert.Equal(t, common.HexToHash("0xf79bde9ddd17963ebce6f7d021d60de7c2bd0db944d23c900c0c0e775f530061"), state.GetLocDynamicArrAtElement(candidateIdsLoc, uint64(15), 1))

	// _candidateIndex: mapping[address] => uint256
	assert.Equal(t, common.HexToHash("0xc9bd304cb5845f29246e1b7d2178c6a415abcc9f15bfc842413f5438a61a2ade"), state.GetLocMappingAtKey(val.ValidatorCandidate.ConsensusAddr, s.ValidatorSetStorageSlot.CandidateIndexSlot))

	// _validatorIds: uint256
	assert.Equal(t, common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000aa"), state.GetLocSimpleVariable(s.ValidatorSetStorageSlot.ValidatorCountSlot))

	// _candidateInfo: mapping[address] => ValidatorCandidate
	infoLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ValidatorSetStorageSlot.CandidateInfoSlot)

	// _validatorIds: mapping[uint256] => address
	assert.Equal(t, common.HexToHash("0xc5d43c4a2513442e32494a8634f103d73dcf3a2851465bd0d02221eac0820c02"), state.GetLocMappingAtKey(common.BigToHash(big.NewInt(int64(0))), s.ValidatorSetStorageSlot.ValidatorIdsSlot))

	validatorCandidateOffset := s.ValidatorSetStorageSlot.ValidatorCandidateOffset
	assert.Equal(t, common.HexToHash("0x6a40c1429db6c94359904880e6af4847b1b19f379ccfa2fe1aa91e1ff8c5cdde"), state.GetLocStructElement(infoLoc, validatorCandidateOffset.AdminOffset))
	assert.Equal(t, common.HexToHash("0x6a40c1429db6c94359904880e6af4847b1b19f379ccfa2fe1aa91e1ff8c5cddf"), state.GetLocStructElement(infoLoc, validatorCandidateOffset.ConsensusAddrOffset))
	assert.Equal(t, common.HexToHash("0x6a40c1429db6c94359904880e6af4847b1b19f379ccfa2fe1aa91e1ff8c5cde0"), state.GetLocStructElement(infoLoc, validatorCandidateOffset.TreasuryAddrOffset))
	assert.Equal(t, common.HexToHash("0x6a40c1429db6c94359904880e6af4847b1b19f379ccfa2fe1aa91e1ff8c5cde2"), state.GetLocStructElement(infoLoc, validatorCandidateOffset.CommissionRateOffset))

	// Test Profile
	// _id2Profile: mapping[address] => Profile
	profileLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.ProfileStoragesSlot.Id2ProfileSlot)
	candidateProfileOffset := s.ProfileStoragesSlot.CandidateProfileOffset
	assert.Equal(t, common.HexToHash("0x6db76b6e1f224552af634485b0df39e324e3146992a8ab6f9618acb87b0fc755"), state.GetLocStructElement(profileLoc, candidateProfileOffset.IdOffset))
	assert.Equal(t, common.HexToHash("0x6db76b6e1f224552af634485b0df39e324e3146992a8ab6f9618acb87b0fc756"), state.GetLocStructElement(profileLoc, candidateProfileOffset.ConsensusOffset))
	assert.Equal(t, common.HexToHash("0x6db76b6e1f224552af634485b0df39e324e3146992a8ab6f9618acb87b0fc757"), state.GetLocStructElement(profileLoc, candidateProfileOffset.AdminOffset))
	assert.Equal(t, common.HexToHash("0x6db76b6e1f224552af634485b0df39e324e3146992a8ab6f9618acb87b0fc758"), state.GetLocStructElement(profileLoc, candidateProfileOffset.TreasuryOffset))
	assert.Equal(t, common.HexToHash("0x6db76b6e1f224552af634485b0df39e324e3146992a8ab6f9618acb87b0fc759"), state.GetLocStructElement(profileLoc, candidateProfileOffset.GovernorOffset))
	pubkeyLoc := state.GetLocStructElement(profileLoc, candidateProfileOffset.PubkeyOffset)
	assert.Equal(t, common.HexToHash("0x6db76b6e1f224552af634485b0df39e324e3146992a8ab6f9618acb87b0fc75a"), pubkeyLoc)
	assert.Equal(t, common.HexToHash("0xa6fd47d619becf29ea4bec5ab864598044ae23ba3e0588e5e2abdac18de6ae5d"), state.GetLocDynamicArrAtElement(pubkeyLoc, uint64(0), 1))
	assert.Equal(t, common.HexToHash("0xa6fd47d619becf29ea4bec5ab864598044ae23ba3e0588e5e2abdac18de6ae5e"), state.GetLocDynamicArrAtElement(pubkeyLoc, uint64(1), 1))

	// _consensus2Id: mapping[TConsensus] => address
	assert.Equal(t, common.HexToHash("0x4b108bfc2de51d6e1b678a350190fe1b895ed120ebdf718c5886c40f636268a0"), state.GetLocMappingAtKey(val.ConsensusAddr, s.ProfileStoragesSlot.Consensus2IdSlot))

	// _registry: mapping[address] => bool
	assert.Equal(t, common.HexToHash("0x2ebeec366ec293eedc12b1ff2578606a4d7b41fe1e662f275ccf1a304dd12d54"), state.GetLocMappingAtKey(val.AdminAddr, s.ProfileStoragesSlot.RegistrySlot))
	assert.Equal(t, common.HexToHash("0x935f895d4dfe6da992b49dc1d0a22fd570f8ce06471946488f37f0c4b5b19903"), state.GetLocMappingAtKey(val.ConsensusAddr, s.ProfileStoragesSlot.RegistrySlot))

	// Test Staking
	// _poolDetail: mapping[address] => PoolDetail
	poolDetailLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.StakingStorageSlot.PoolDetailSlot)
	poolDetailOffset := s.StakingStorageSlot.PoolDetailOffset
	assert.Equal(t, common.HexToHash("0xc2173dae99d1ead7b236d87790c699953d911adfcf7e901b96bd95f0cefccc4c"), state.GetLocStructElement(poolDetailLoc, poolDetailOffset.PidOffset))
	assert.Equal(t, common.HexToHash("0xc2173dae99d1ead7b236d87790c699953d911adfcf7e901b96bd95f0cefccc4d"), state.GetLocStructElement(poolDetailLoc, poolDetailOffset.ShadowedPoolAdminOffset))
	assert.Equal(t, common.HexToHash("0xc2173dae99d1ead7b236d87790c699953d911adfcf7e901b96bd95f0cefccc4e"), state.GetLocStructElement(poolDetailLoc, poolDetailOffset.StakingAmountOffset))
	assert.Equal(t, common.HexToHash("0xc2173dae99d1ead7b236d87790c699953d911adfcf7e901b96bd95f0cefccc4f"), state.GetLocStructElement(poolDetailLoc, poolDetailOffset.StakingTotalOffset))

	// delegatingAmount: mapping[address] => uint256
	dloc := state.GetLocStructElement(poolDetailLoc, poolDetailOffset.DelegatingAmountOffset)
	assert.Equal(t, common.HexToHash("0xf18d6cb971f7c2b2f6ad3923215e353d8c1e56670572aa2e75a11a8166b9c535"), getLocMappingAtKey(val.AdminAddr, dloc))

	// lastDelegatingTimestamp: mapping[address] => uint256
	lloc := state.GetLocStructElement(poolDetailLoc, poolDetailOffset.LastDelegatingTimestampOffset)
	assert.Equal(t, common.HexToHash("0x21deda3fda1b7a6abbb4276d6b424b2da905c78cde2d714bc6531d8db0539051"), getLocMappingAtKey(val.AdminAddr, lloc))

	// _stakingPool: mapping[address] => PoolFields
	poolFieldsLoc := state.GetLocMappingAtKey(val.ConsensusAddr, s.StakingStorageSlot.StakingPoolSlot)
	poolFieldsOffset := s.StakingStorageSlot.PoolFieldsOffset
	assert.Equal(t, common.HexToHash("0x9d0cb7ed7b2894116569eb0df0ef37e88b36d2f3d4ca20e28723ee3c6e7196d1"), state.GetLocStructElement(poolFieldsLoc, poolFieldsOffset.SharesInnerOffset))
	assert.Equal(t, common.HexToHash("0x9d0cb7ed7b2894116569eb0df0ef37e88b36d2f3d4ca20e28723ee3c6e7196d2"), state.GetLocStructElement(poolFieldsLoc, poolFieldsOffset.SharesLastPeriodOffset))

	// _userReward: mapping[address]mapping[address] => UserRewardFields
	userLoc := getLocMappingAtKey(val.AdminAddr, state.GetLocMappingAtKey(val.ConsensusAddr, s.StakingStorageSlot.UserRewardSlot))
	userRewardFieldsOffset := s.StakingStorageSlot.UserRewardFieldsOffset
	assert.Equal(t, common.HexToHash("0xc8a608418ed0adc246dc035fedf0532814df6c1fb1ae549afed0a6eb5827a32f"), state.GetLocStructElement(userLoc, userRewardFieldsOffset.LowestAmountOffset))
	assert.Equal(t, common.HexToHash("0xc8a608418ed0adc246dc035fedf0532814df6c1fb1ae549afed0a6eb5827a330"), state.GetLocStructElement(userLoc, userRewardFieldsOffset.LastPeriodOffset))

	// _adminOfActivePoolMapping: mapping[address] => address
	assert.Equal(t, common.HexToHash("0x9dfd02398427702df814d336fa2c0ca54f258444ccac6d6c7b70920e6d7fc060"), state.GetLocMappingAtKey(val.AdminAddr, s.StakingStorageSlot.AdminOfActivePoolMappingSlot))

	// Test RoninTrustedOrganization
	// _consensusList: []TConsensus
	consensusLoc := state.GetLocSimpleVariable(s.TrustedOrgSlot.ConsensusListSlot)
	assert.Equal(t, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000009"), consensusLoc)
	// _governorList: []address
	governorLoc := state.GetLocSimpleVariable(s.TrustedOrgSlot.GovernorListSlot)
	assert.Equal(t, common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000000a"), governorLoc)
	// _totalWeight: uint256
	assert.Equal(t, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"), state.GetLocSimpleVariable(s.TrustedOrgSlot.TotalWeightSlot))
	// _consensusList:
	assert.Equal(t, common.HexToHash("0x6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af"), state.GetLocDynamicArrAtElement(consensusLoc, uint64(0), 1))
	assert.Equal(t, common.HexToHash("0x6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7b0"), state.GetLocDynamicArrAtElement(consensusLoc, uint64(1), 1))
	// _governorList
	assert.Equal(t, common.HexToHash("0xc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8"), state.GetLocDynamicArrAtElement(governorLoc, uint64(0), 1))
	assert.Equal(t, common.HexToHash("0xc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a9"), state.GetLocDynamicArrAtElement(governorLoc, uint64(1), 1))
	// _addedBlock
	assert.Equal(t, common.HexToHash("0x494858555715775ee11a252997bcf8c561bdcf8d161dd7b58cf696c5635e03d5"), state.GetLocMappingAtKey(common.HexToHash("0xBEEFBEEFBEEF0002"), s.TrustedOrgSlot.AddedBlockSlot))
	assert.Equal(t, common.HexToHash("0x81f61bd1e6d54acd6926fce0426779f8e35a0484e1accd56f4d5e8340fda5fba"), state.GetLocMappingAtKey(common.HexToHash("0xDEADDEADDEAD00AB"), s.TrustedOrgSlot.AddedBlockSlot))
	// _consensusWeight
	assert.Equal(t, common.HexToHash("0xd165eac8b99e8a03d12e9250056a3151a2fa9f4a5b50e39dcdeb1874d90c4bf1"), state.GetLocMappingAtKey(common.HexToHash("0xBEEFBEEFBEEF0002"), s.TrustedOrgSlot.ConsensusWeightSlot))
	assert.Equal(t, common.HexToHash("0x40f11a548734dff4657261ddc229586f9bdc14d9e3716e6e9964b331808c2e1e"), state.GetLocMappingAtKey(common.HexToHash("0xDEADDEADDEAD00AB"), s.TrustedOrgSlot.ConsensusWeightSlot))
	// _governorWeight
	assert.Equal(t, common.HexToHash("0x64812ea843d7149bce1fb7ff451f6047dc05156ed859df7af6ad2831f0f571f4"), state.GetLocMappingAtKey(common.HexToHash("0xBEEFBEEFBEEF0001"), s.TrustedOrgSlot.GovernorWeightSlot))
	assert.Equal(t, common.HexToHash("0x564949134181ec11bc75ff179d70b92f617f4f154905ebd42b2057f547755c4b"), state.GetLocMappingAtKey(common.HexToHash("0xDEADDEADDEAD00AA"), s.TrustedOrgSlot.GovernorWeightSlot))

	// Test SlashIndicator
	// _creditScore
}

func TestShadowSwitchValuePreprocessing(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cfgPath := filepath.Join(cwd, "config/config.test.yaml")
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		fmt.Println("err", err)
		t.Fatal(err)
	}

	newVals := NewValidators(cfg.NewValidatorConfigs, 9999999, 9999999)
	val := newVals[0]

	// pubkey
	pubkey := "a49891083c05e85a9023d20097367d5c203bfaa4b0f02753dbb85580508fbacee0fb7e7235c7fc032176d97e3e3b101c"
	assert.Equal(t, common.HexToHash("0xa49891083c05e85a9023d20097367d5c203bfaa4b0f02753dbb85580508fbace"), common.HexToHash(pubkey[:64]))
	assert.Equal(t, common.HexToHash("0xe0fb7e7235c7fc032176d97e3e3b101c00000000000000000000000000000000"), common.HexToHash(pubkey[64:]+"00000000000000000000000000000000"))

	// Validator Set
	assert.Equal(t, common.HexToHash("0xE5635B0172e7488072348a0395C983DfDA68755A"), val.AdminAddr)
	assert.Equal(t, common.HexToHash("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), val.Index)

	// Default value
	assert.Equal(t, common.HexToHash("0x0"), common.Hash{})
	assert.Equal(t, common.HexToHash("0x0"), validatorCandidate{}.Admin)
	assert.Equal(t, common.HexToHash("0x0"), validatorCandidate{}.ConsensusAddr)
	assert.Equal(t, common.HexToHash("0x0"), validatorCandidate{}.CommissionRate)
	assert.Equal(t, common.HexToHash("0x0"), validatorCandidate{}.TreasuryAddr)
}

func TestAddBalanceToBlankWallet(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cfgPath := filepath.Join(cwd, "config/config.mainnet10.yaml")
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		fmt.Println("err", err)
		t.Fatal(err)
	}

	newVals := NewValidators(cfg.NewValidatorConfigs, 9999999, 9999999)

	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)

	s := &shadowSwitch{
		NewValidators:       newVals,
		ValidatorSetAddress: cfg.ValidatorSetAddress,
		ProfileAddress:      cfg.ProfileAddress,
		StakingAddress:      cfg.StakingAddress,
		TrustedOrgAddress:   cfg.TrustedOrgAddress,
		state:               statedb,

		ValidatorSetStorageSlot: cfg.ValidatorSetStorageSlot,
		ProfileStoragesSlot:     cfg.ProfileStorageSlot,
		StakingStorageSlot:      cfg.StakingStorageSlot,
		TrustedOrgSlot:          cfg.TrustedOrgSlot,

		NewBlankWallets:    cfg.NewBlankWallets,
		BlankWalletBalance: cfg.BlankWalletBalance,
	}

	assert.Equal(t, 10, len(s.NewBlankWallets))

	for _, w := range s.NewBlankWallets {
		s.state.AddBalance(w.WalletAddr, s.BlankWalletBalance)
	}

	for i := range s.NewBlankWallets {
		assert.Equal(t, new(big.Int).SetUint64(1000), statedb.GetBalance(s.NewBlankWallets[i].WalletAddr))
	}
}
