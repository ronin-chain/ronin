package v2

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
)

func CalcDifficultyHope(snap *Snapshot, signer common.Address) *big.Int {
	if snap == nil || len(snap.Validators) == 0 {
		return nil
	}

	// Use the new consecutive block logic
	n := snap.TurnLength
	if n == 0 {
		n = 1 // Default to 1 if not set
	}

	// Find the index of the current signer in the validator set
	signerIndex := -1
	for i, producer := range snap.BlockProducers {
		if producer == signer {
			signerIndex = i
			break
		}
	}

	if signerIndex == -1 {
		return diffNoTurn // Not a validator
	}

	// Determine whose turn it is
	turnIndex := (snap.Number / uint64(n)) % uint64(len(snap.BlockProducers))

	if uint64(signerIndex) == turnIndex {
		return diffInTurn
	}
	return diffNoTurn
}

// This function would be part of the `Consortium` struct in `consensus/consortium/v2/consortium.go`
func (c *Consortium) blockAvoidanceCheck(header *types.Header, snap *Snapshot, author common.Address, chain consensus.ChainHeaderReader) error {
	// Get 'n', the number of consecutive blocks for this epoch
	n := snap.TurnLength
	if n <= 1 {
		return nil // The rule doesn't apply if the feature is off.
	}

	validatorCount := uint64(len(snap.Validators))
	// Define "recent past" as the number of blocks a bare majority would produce
	historyLimit := (validatorCount/2+1)*uint64(n) - 1

	producedCount := 0
	// Look backwards through the chain's history
	for i := uint64(0); i < historyLimit && header.Number.Uint64() > i+1; i++ {
		ancestor := chain.GetHeaderByNumber(header.Number.Uint64() - 1 - i)
		if ancestor == nil {
			return consensus.ErrUnknownAncestor
		}

		// Find out who produced that old block
		ancestorAuthor, err := ecrecover(ancestor, c.signatures, c.chainConfig.ChainID)
		if err != nil {
			return err
		}

		// Was it the same validator trying to produce the current block?
		if ancestorAuthor == author {
			producedCount++
		}
	}

	// If the validator has already produced 'n' blocks recently, reject this new block.
	if producedCount >= int(n) {
		return fmt.Errorf("validator %s produced blocks too frequently", author.Hex())
	}

	return nil
}

func (c *Consortium) getTurnLength(chain consensus.ChainHeaderReader, header *types.Header) (*uint8, error) {
	if !c.chainConfig.IsHope(header.Number) {
		one := uint8(1)
		return &one, nil
	}

	if c.chainConfig.Consortium.ConsecutiveEpoch == 0 {
		return c.getTurnLengthFromContract(chain, header)
	}

	return c.getTurnLengthFromConfig(chain, header)
}

// getTurnLengthFromContract gets the turn length from the contract, allow to override the turn length from config
func (c *Consortium) getTurnLengthFromContract(chain consensus.ChainHeaderReader, header *types.Header) (*uint8, error) {
	return nil, nil
}

// getTurnLengthFromConfig is the fallback function for switching to manual turn length
func (c *Consortium) getTurnLengthFromConfig(chain consensus.ChainHeaderReader, header *types.Header) (*uint8, error) {
	return nil, nil
}
