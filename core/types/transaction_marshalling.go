// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

// txJSON is the JSON representation of transactions.
type txJSON struct {
	Type hexutil.Uint64 `json:"type"`

	// Common transaction fields:
	Nonce                *hexutil.Uint64 `json:"nonce"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxFeePerBlobGas     *hexutil.Big    `json:"maxFeePerBlobGas,omitempty"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	Value                *hexutil.Big    `json:"value"`
	Data                 *hexutil.Bytes  `json:"input"`
	V                    *hexutil.Big    `json:"v"`
	R                    *hexutil.Big    `json:"r"`
	S                    *hexutil.Big    `json:"s"`
	To                   *common.Address `json:"to"`
	AuthorizationList    []Authorization `json:"authorizationList,omitempty"`
	YParity              *hexutil.Uint64 `json:"yParity,omitempty"`

	// Access list transaction fields:
	ChainID    *hexutil.Big `json:"chainId,omitempty"`
	AccessList *AccessList  `json:"accessList,omitempty"`

	// Sponsored transaction fields
	ExpiredTime *hexutil.Uint64 `json:"expiredTime,omitempty"`
	PayerV      *hexutil.Big    `json:"payerV,omitempty"`
	PayerR      *hexutil.Big    `json:"payerR,omitempty"`
	PayerS      *hexutil.Big    `json:"payerS,omitempty"`

	// Blob transaction fields
	BlobVersionedHashes []common.Hash        `json:"blobVersionedHashes,omitempty"`
	Blobs               []kzg4844.Blob       `json:"blobs,omitempty"`
	Commitments         []kzg4844.Commitment `json:"commitments,omitempty"`
	Proofs              []kzg4844.Proof      `json:"proofs,omitempty"`

	// Only used for encoding:
	Hash common.Hash `json:"hash"`
}

// yParityValue returns the YParity value from JSON. For backwards-compatibility reasons,
// this can be given in the 'v' field or the 'yParity' field. If both exist, they must match.
func (tx *txJSON) yParityValue() (*big.Int, error) {
	if tx.YParity != nil {
		val := uint64(*tx.YParity)
		if val != 0 && val != 1 {
			return nil, errInvalidYParity
		}
		bigval := new(big.Int).SetUint64(val)
		if tx.V != nil && tx.V.ToInt().Cmp(bigval) != 0 {
			return nil, errVYParityMismatch
		}
		return bigval, nil
	}
	if tx.V != nil {
		return tx.V.ToInt(), nil
	}
	return nil, errVYParityMissing
}

// MarshalJSON marshals as JSON with a hash.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = t.Hash()
	enc.Type = hexutil.Uint64(t.Type())
	nonce := t.Nonce()
	gas := t.Gas()
	data := t.Data()
	v, r, s := t.RawSignatureValues()
	enc.Nonce = (*hexutil.Uint64)(&nonce)
	enc.Gas = (*hexutil.Uint64)(&gas)
	enc.Value = (*hexutil.Big)(t.Value())
	enc.Data = (*hexutil.Bytes)(&data)
	enc.To = t.To()
	enc.V = (*hexutil.Big)(v)
	enc.R = (*hexutil.Big)(r)
	enc.S = (*hexutil.Big)(s)

	// Other fields are set conditionally depending on tx type.
	switch tx := t.inner.(type) {
	case *LegacyTx:
		enc.GasPrice = (*hexutil.Big)(tx.GasPrice)
	case *AccessListTx:
		enc.GasPrice = (*hexutil.Big)(tx.GasPrice)
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.AccessList = &tx.AccessList
	case *DynamicFeeTx:
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.AccessList = &tx.AccessList
		enc.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap)
		enc.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap)
	case *SponsoredTx:
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap)
		enc.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap)
		enc.ExpiredTime = (*hexutil.Uint64)(&tx.ExpiredTime)
		enc.PayerV = (*hexutil.Big)(tx.PayerV)
		enc.PayerR = (*hexutil.Big)(tx.PayerR)
		enc.PayerS = (*hexutil.Big)(tx.PayerS)
	case *BlobTx:
		enc.ChainID = (*hexutil.Big)(tx.ChainID.ToBig())
		enc.AccessList = &tx.AccessList
		enc.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap.ToBig())
		enc.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap.ToBig())
		enc.MaxFeePerBlobGas = (*hexutil.Big)(tx.BlobFeeCap.ToBig())
		enc.BlobVersionedHashes = tx.BlobHashes
		if sidecar := tx.Sidecar; sidecar != nil {
			enc.Blobs = tx.Sidecar.Blobs
			enc.Commitments = tx.Sidecar.Commitments
			enc.Proofs = tx.Sidecar.Proofs
		}
	case *SetCodeTx:
		enc.ChainID = (*hexutil.Big)(new(big.Int).SetUint64(tx.ChainID))
		enc.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap.ToBig())
		enc.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap.ToBig())
		enc.AccessList = &tx.AccessList
		enc.AuthorizationList = tx.AuthList
		yparity := tx.V.Uint64()
		enc.YParity = (*hexutil.Uint64)(&yparity)
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	var to *common.Address
	if dec.To != nil {
		to = dec.To
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' in transaction")
	}
	nonce := uint64(*dec.Nonce)
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' in transaction")
	}
	gas := uint64(*dec.Gas)
	if dec.Value == nil {
		return errors.New("missing required field 'value' in transaction")
	}
	value := (*big.Int)(dec.Value)
	if dec.Data == nil {
		return errors.New("missing required field 'input' in transaction")
	}
	data := *dec.Data
	v, err := dec.yParityValue()
	if err != nil {
		return err
	}
	if dec.R == nil {
		return errors.New("missing required field 'r' in transaction")
	}
	r := (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' in transaction")
	}
	s := (*big.Int)(dec.S)
	withSignature := v.Sign() != 0 || r.Sign() != 0 || s.Sign() != 0
	if withSignature {
		maybeProtected := false
		if dec.Type == LegacyTxType {
			maybeProtected = true
		}

		if err := sanityCheckSignature(v, r, s, maybeProtected); err != nil {
			return err
		}
	}

	// Decode / verify fields according to transaction type.
	var inner TxData
	switch dec.Type {
	case LegacyTxType:
		itx := LegacyTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		if dec.GasPrice == nil {
			return errors.New("missing required field 'gasPrice' in transaction")
		}
		itx.GasPrice = (*big.Int)(dec.GasPrice)

	case AccessListTxType:
		itx := AccessListTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		// Access list is optional for now.
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.GasPrice == nil {
			return errors.New("missing required field 'gasPrice' in transaction")
		}
		itx.GasPrice = (*big.Int)(dec.GasPrice)

	case DynamicFeeTxType:
		itx := DynamicFeeTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		// Access list is optional for now.
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = (*big.Int)(dec.MaxPriorityFeePerGas)
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = (*big.Int)(dec.MaxFeePerGas)

	case SponsoredTxType:
		itx := SponsoredTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = (*big.Int)(dec.MaxPriorityFeePerGas)
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = (*big.Int)(dec.MaxFeePerGas)
		if dec.ExpiredTime == nil {
			return errors.New("missing required field 'expiredTime' in transaction")
		}
		itx.ExpiredTime = uint64(*dec.ExpiredTime)
		if dec.PayerV == nil {
			return errors.New("missing required field 'payerV' in transaction")
		}
		itx.PayerV = (*big.Int)(dec.PayerV)
		if dec.PayerR == nil {
			return errors.New("missing required field 'payerR' in transaction")
		}
		itx.PayerR = (*big.Int)(dec.PayerR)
		if dec.PayerS == nil {
			return errors.New("missing required field 'payerS' in transaction")
		}
		itx.PayerS = (*big.Int)(dec.PayerS)
		if err := sanityCheckSignature(itx.PayerV, itx.PayerR, itx.PayerS, false); err != nil {
			return err
		}
	case BlobTxType:
		itx := BlobTx{
			Nonce: nonce,
			Gas:   gas,
			Value: uint256.MustFromBig(value),
			Data:  data,
		}
		inner = &itx
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = uint256.MustFromBig((*big.Int)(dec.ChainID))
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = uint256.MustFromBig((*big.Int)(dec.MaxPriorityFeePerGas))
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = uint256.MustFromBig((*big.Int)(dec.MaxFeePerGas))
		if *to == (common.Address{}) {
			return errors.New("missing required field 'to' in transaction")
		}
		itx.To = *to

		if dec.MaxFeePerBlobGas == nil {
			return errors.New("missing required field 'maxFeePerBlobGas' in transaction")
		}
		itx.BlobFeeCap = uint256.MustFromBig((*big.Int)(dec.MaxFeePerBlobGas))
		if dec.BlobVersionedHashes == nil {
			return errors.New("missing required field 'blobVersionedHashes' in transaction")
		}
		itx.BlobHashes = dec.BlobVersionedHashes
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		// signature R
		var overflow bool
		itx.R, overflow = uint256.FromBig(r)
		if overflow {
			return errors.New("'r' value overflows uint256")
		}
		// signature S
		itx.S, overflow = uint256.FromBig(s)
		if overflow {
			return errors.New("'s' value overflows uint256")
		}
		itx.V, overflow = uint256.FromBig(v)
		if overflow {
			return errors.New("'v' value overflows uint256")
		}
	case SetCodeTxType:
		itx := SetCodeTx{
			Nonce: nonce,
			Gas:   gas,
			Value: uint256.MustFromBig(value),
			Data:  data,
		}
		inner = &itx
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = dec.ChainID.ToInt().Uint64()
		if dec.To == nil {
			return errors.New("missing required field 'to' in transaction")
		}
		itx.To = *dec.To
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = uint256.MustFromBig((*big.Int)(dec.MaxPriorityFeePerGas))
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = uint256.MustFromBig((*big.Int)(dec.MaxFeePerGas))
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.AuthorizationList == nil {
			return errors.New("missing required field 'authorizationList' in transaction")
		}
		itx.AuthList = dec.AuthorizationList

		// signature R
		var overflow bool
		itx.R, overflow = uint256.FromBig(r)
		if overflow {
			return errors.New("'r' value overflows uint256")
		}
		// signature S
		itx.S, overflow = uint256.FromBig(s)
		if overflow {
			return errors.New("'s' value overflows uint256")
		}
		itx.V, overflow = uint256.FromBig(v)
		if overflow {
			return errors.New("'v' value overflows uint256")
		}
		if itx.V.Sign() != 0 || itx.R.Sign() != 0 || itx.S.Sign() != 0 {
			if err := sanityCheckSignature(v, itx.R.ToBig(), itx.S.ToBig(), false); err != nil {
				return err
			}
		}
	default:
		return ErrTxTypeNotSupported
	}

	// Now set the inner transaction.
	t.setDecoded(inner, 0)

	// TODO: check hash here?
	return nil
}
