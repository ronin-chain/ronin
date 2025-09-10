package vm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func TestInternalTransactionParentOrder_NestedCalls(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}

	publishEvent := PublishEvent{
		OpCodes: []OpCode{
			CALL,
			DELEGATECALL,
			CREATE,
			CREATE2,
		},
		Event: &InternalTransactionEvent{},
	}

	internalTransactions := make([]*types.InternalTransaction, 0)
	evm := NewEVM(
		BlockContext{
			BlockNumber:          common.Big0,
			InternalTransactions: &internalTransactions,
			Transfer:             func(_ StateDB, _, _ common.Address, _ *big.Int) {},
			PublishEvents:        make(PublishEventsMap),
			CurrentTransaction:   types.NewTx(&types.LegacyTx{}),
		},
		TxContext{},
		statedb,
		&params.ChainConfig{
			IstanbulBlock: common.Big0,
		},
		Config{},
	)

	for _, opcode := range publishEvent.OpCodes {
		evm.Context.PublishEvents[opcode] = publishEvent.Event
	}

	/*
	   pragma solidity ^0.8.18;

	   contract Test1 {
	       fallback() external {
	           address test2 = address(0x202);
	           test2.call{gas: 200000}("");
	       }
	   }
	*/
	test1Contract := common.Hex2Bytes("608060405234801561001057600080fd5b5060f58061001f6000396000f3fe6080604052348015600f57600080fd5b50600061020290508073ffffffffffffffffffffffffffffffffffffffff1662030d40604051603c9060ac565b60006040518083038160008787f1925050503d80600081146078576040519150601f19603f3d011682016040523d82523d6000602084013e607d565b606091505b005b600081905092915050565b50565b60006098600083607f565b915060a182608a565b600082019050919050565b600060b582608d565b915081905091905056fea26469706673582212208ba7f80d6a96ebc77318dbf151ea79ff647cf772581f7c4acb3fd1d2babfdfb464736f6c63430008120033")
	test1Address := common.BigToAddress(big.NewInt(0x201))

	/*
	   pragma solidity ^0.8.18;

	   contract Test2 {
	       fallback() external {
	           address test3 = address(0x203);
	           test3.call{gas: 100000}("");
	       }
	   }
	*/
	test2Contract := common.Hex2Bytes("608060405234801561001057600080fd5b5060f58061001f6000396000f3fe6080604052348015600f57600080fd5b50600061020390508073ffffffffffffffffffffffffffffffffffffffff16620186a0604051603c9060ac565b60006040518083038160008787f1925050503d80600081146078576040519150601f19603f3d011682016040523d82523d6000602084013e607d565b606091505b005b600081905092915050565b50565b60006098600083607f565b915060a182608a565b600082019050919050565b600060b582608d565b915081905091905056fea2646970667358221220dc64986e2e2b84dfc27218360968609146254d5606999d4de5187b18fc1ca7a664736f6c63430008120033")
	test2Address := common.BigToAddress(big.NewInt(0x202))

	/*
	   pragma solidity ^0.8.18;

	   contract Test3 {
	       fallback() external {}
	   }
	*/
	test3Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220e499609e32f510161ba4b0c5e900173a55b3457a85347d20c83099f6d20b2ace64736f6c63430008120033")
	test3Address := common.BigToAddress(big.NewInt(0x203))

	// Deploy bytecode to addresses
	statedb.SetCode(test1Address, test1Contract)
	statedb.SetCode(test2Address, test2Contract)
	statedb.SetCode(test3Address, test3Contract)

	// "Deploy" by calling each once to get the runtime bytecode
	deployedTest1, _, err := evm.Call(AccountRef(test1Address), test1Address, []byte{}, 100_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest2, _, err := evm.Call(AccountRef(test2Address), test2Address, []byte{}, 100_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest3, _, err := evm.Call(AccountRef(test3Address), test3Address, []byte{}, 100_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	statedb.SetCode(test1Address, deployedTest1)
	statedb.SetCode(test2Address, deployedTest2)
	statedb.SetCode(test3Address, deployedTest3)

	// Execute: Test1() -> CALL Test2() -> CALL Test3()
	_, _, err = evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	internalTxs := *evm.Context.InternalTransactions
	if len(internalTxs) != 2 {
		t.Fatalf("internal txs length mismatch: got %d, want %d", len(internalTxs), 2)
	}
	// Sanity: A->B, then B->C
	if internalTxs[0].From != test1Address || internalTxs[0].To != test2Address {
		t.Fatalf("unexpected internal tx[0]: %+v", internalTxs[0])
	}
	if internalTxs[1].From != test2Address || internalTxs[1].To != test3Address {
		t.Fatalf("unexpected internal tx[1]: %+v", internalTxs[1])
	}

	// The parent of (B->C) should be the order of (A->B)
	gotParent := internalTxs[1].InternalTransactionBody.ParentOrder
	wantParent := internalTxs[0].InternalTransactionBody.Order
	if gotParent != wantParent {
		t.Fatalf("parent order mismatch for nested call; got %d, want %d", gotParent, wantParent)
	}
}
