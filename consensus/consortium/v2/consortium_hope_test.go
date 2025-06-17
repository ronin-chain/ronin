package v2

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
	"github.com/ethereum/go-ethereum/params"
)

func TestHopeTurnLength(t *testing.T) {
	// Create test config with Hope hardfork enabled
	config := &params.ChainConfig{
		ChainID:   big.NewInt(2025),
		HopeBlock: big.NewInt(0), // Enable Hope from genesis
		Consortium: &params.ConsortiumConfig{
			Period:     3,
			EpochV2:    200,
			TurnLength: 3, // Allow 3 consecutive blocks per validator
		},
	}

	// Test validators (4 validators)
	validators := []common.Address{
		common.HexToAddress("0x1000000000000000000000000000000000000001"),
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
		common.HexToAddress("0x3000000000000000000000000000000000000003"),
		common.HexToAddress("0x4000000000000000000000000000000000000004"),
	}

	// Create snapshot with validators
	snap := &Snapshot{
		Number:               0,
		Hash:                 common.Hash{},
		ValidatorsWithBlsPub: make([]finality.ValidatorWithBlsPub, len(validators)),
		TurnLength:           config.Consortium.TurnLength,
		chainConfig:          config,
	}

	// Convert validators to ValidatorsWithBlsPub format
	for i, val := range validators {
		snap.ValidatorsWithBlsPub[i] = finality.ValidatorWithBlsPub{
			Address: val,
		}
	}

	// Test consecutive block production logic
	testCases := []struct {
		blockNumber    uint64
		validator      common.Address
		expectedInTurn bool
		description    string
	}{
		// First validator (index 0) should produce blocks 1, 2, 3
		{1, validators[0], true, "Validator 0 should be in turn for block 1"},
		{2, validators[0], true, "Validator 0 should be in turn for block 2"},
		{3, validators[0], true, "Validator 0 should be in turn for block 3"},
		{4, validators[0], false, "Validator 0 should NOT be in turn for block 4"},

		// Second validator (index 1) should produce blocks 4, 5, 6
		{4, validators[1], true, "Validator 1 should be in turn for block 4"},
		{5, validators[1], true, "Validator 1 should be in turn for block 5"},
		{6, validators[1], true, "Validator 1 should be in turn for block 6"},
		{7, validators[1], false, "Validator 1 should NOT be in turn for block 7"},

		// Third validator (index 2) should produce blocks 7, 8, 9
		{7, validators[2], true, "Validator 2 should be in turn for block 7"},
		{8, validators[2], true, "Validator 2 should be in turn for block 8"},
		{9, validators[2], true, "Validator 2 should be in turn for block 9"},

		// Fourth validator (index 3) should produce blocks 10, 11, 12
		{10, validators[3], true, "Validator 3 should be in turn for block 10"},
		{11, validators[3], true, "Validator 3 should be in turn for block 11"},
		{12, validators[3], true, "Validator 3 should be in turn for block 12"},

		// Cycle repeats: First validator again for blocks 13, 14, 15
		{13, validators[0], true, "Validator 0 should be in turn for block 13 (cycle repeats)"},
		{14, validators[0], true, "Validator 0 should be in turn for block 14"},
		{15, validators[0], true, "Validator 0 should be in turn for block 15"},
	}

	for _, tc := range testCases {
		snap.Number = tc.blockNumber - 1 // Snapshot represents state before block
		result := snap.inturnWithTurnLength(tc.validator, config, big.NewInt(int64(tc.blockNumber)))
		if result != tc.expectedInTurn {
			t.Errorf("Block %d: %s - Expected %v, got %v",
				tc.blockNumber, tc.description, tc.expectedInTurn, result)
		}
	}
}

func TestHopeTurnLengthWithDifferentConfig(t *testing.T) {
	// Test with turn length of 2
	config := &params.ChainConfig{
		ChainID:   big.NewInt(2025),
		HopeBlock: big.NewInt(0),
		Consortium: &params.ConsortiumConfig{
			TurnLength: 2, // Allow 2 consecutive blocks per validator
		},
	}

	validators := []common.Address{
		common.HexToAddress("0x1000000000000000000000000000000000000001"),
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
	}

	snap := &Snapshot{
		Number:               0,
		Hash:                 common.Hash{},
		ValidatorsWithBlsPub: make([]finality.ValidatorWithBlsPub, len(validators)),
		TurnLength:           config.Consortium.TurnLength,
		chainConfig:          config,
	}

	for i, val := range validators {
		snap.ValidatorsWithBlsPub[i] = finality.ValidatorWithBlsPub{
			Address: val,
		}
	}

	testCases := []struct {
		blockNumber    uint64
		validator      common.Address
		expectedInTurn bool
	}{
		// First validator: blocks 1, 2
		{1, validators[0], true},
		{2, validators[0], true},
		{3, validators[0], false},

		// Second validator: blocks 3, 4
		{3, validators[1], true},
		{4, validators[1], true},
		{5, validators[1], false},

		// Cycle repeats: First validator: blocks 5, 6
		{5, validators[0], true},
		{6, validators[0], true},
	}

	for _, tc := range testCases {
		snap.Number = tc.blockNumber - 1
		result := snap.inturnWithTurnLength(tc.validator, config, big.NewInt(int64(tc.blockNumber)))
		if result != tc.expectedInTurn {
			t.Errorf("Block %d with validator %v: Expected %v, got %v",
				tc.blockNumber, tc.validator, tc.expectedInTurn, result)
		}
	}
}

func TestPreHopeHardfork(t *testing.T) {
	// Test that before Hope hardfork, the original logic is used
	config := &params.ChainConfig{
		ChainID:   big.NewInt(2025),
		HopeBlock: big.NewInt(100), // Hope enabled at block 100
		Consortium: &params.ConsortiumConfig{
			TurnLength: 3, // This should be ignored before Hope
		},
	}

	validators := []common.Address{
		common.HexToAddress("0x1000000000000000000000000000000000000001"),
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
	}

	snap := &Snapshot{
		Number:               49, // Block 50, before Hope hardfork
		Hash:                 common.Hash{},
		ValidatorsWithBlsPub: make([]finality.ValidatorWithBlsPub, len(validators)),
		chainConfig:          config,
	}

	for i, val := range validators {
		snap.ValidatorsWithBlsPub[i] = finality.ValidatorWithBlsPub{
			Address: val,
		}
	}

	// Before Hope, should use original round-robin logic
	// Block 50: (50) % 2 = 0, so validator[0] should be in turn
	result1 := snap.inturnWithTurnLength(validators[0], config, big.NewInt(50))
	result2 := snap.inturnWithTurnLength(validators[1], config, big.NewInt(50))

	if !result1 {
		t.Error("Before Hope hardfork, validator 0 should be in turn for block 50")
	}
	if result2 {
		t.Error("Before Hope hardfork, validator 1 should NOT be in turn for block 50")
	}
}
