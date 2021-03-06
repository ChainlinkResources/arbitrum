/*
* Copyright 2020, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package message

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ERC721 struct {
	To           common.Address
	From         common.Address
	TokenAddress common.Address
	Id           *big.Int
}

func (m ERC721) String() string {
	return fmt.Sprintf("ERC721(to: %v, from: %v, token: %v, id: %v)",
		m.To,
		m.From,
		m.TokenAddress,
		m.Id,
	)
}

func (m ERC721) Equals(other Message) bool {
	o, ok := other.(ERC721)
	if !ok {
		return false
	}
	return m.To == o.To &&
		m.From == o.From &&
		m.TokenAddress == o.TokenAddress &&
		m.Id.Cmp(o.Id) == 0
}

func (m ERC721) Type() MessageType {
	return ERC721Type
}

func (m ERC721) GetFuncName() string {
	return "ERC721Transfer"
}

func (m ERC721) asValue() value.Value {
	val1, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.TokenAddress),
		addressToIntValue(m.To),
		value.NewIntValue(new(big.Int).Set(m.Id)),
	})
	val2, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(big.NewInt(int64(m.Type()))),
		addressToIntValue(m.From),
		val1,
	})
	return val2
}

func UnmarshalERC721(val value.Value) (ERC721, error) {
	from, token, to, amount, err := unmarshalToken(val, ERC721Type)
	if err != nil {
		return ERC721{}, err
	}

	return ERC721{
		To:           to,
		From:         from,
		TokenAddress: token,
		Id:           amount,
	}, nil
}

type DeliveredERC721 struct {
	ERC721
	BlockNum   *common.TimeBlocks
	Timestamp  *big.Int
	MessageNum *big.Int
}

// Equals check for equality between this object and any other message by
// checking for full equality of all members
func (m DeliveredERC721) Equals(other Message) bool {
	o, ok := other.(DeliveredERC721)
	if !ok {
		return false
	}
	return m.ERC721.Equals(o.ERC721) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.Timestamp.Cmp(o.Timestamp) == 0 &&
		m.MessageNum.Cmp(o.MessageNum) == 0
}

func (m DeliveredERC721) deliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m DeliveredERC721) deliveredTimestamp() *big.Int {
	return m.Timestamp
}

func (m DeliveredERC721) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Address(m.TokenAddress),
		hashing.Uint256(m.Id),
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.Timestamp),
		hashing.Uint256(m.MessageNum),
	)
}

func (m DeliveredERC721) ReceiptHash() common.Hash {
	return value.NewIntValue(m.MessageNum).ToBytes()
}

func (m DeliveredERC721) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		addressToIntValue(m.TokenAddress),
		value.NewIntValue(new(big.Int).Set(m.Id)),
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).Set(m.Timestamp)),
		value.NewIntValue(new(big.Int).Set(m.MessageNum)),
	})
	return val
}

func UnmarshalERC721FromCheckpoint(v value.Value) (DeliveredERC721, error) {
	tup, ok := v.(value.TupleValue)
	failRet := DeliveredERC721{}
	if !ok || tup.Len() != 7 {
		return failRet, errors.New("tx val must be 7-tuple")
	}
	to, _ := tup.GetByInt64(0)
	toInt, ok := to.(value.IntValue)
	if !ok {
		return failRet, errors.New("to must be int")
	}
	from, _ := tup.GetByInt64(1)
	fromInt, ok := from.(value.IntValue)
	if !ok {
		return failRet, errors.New("from must be int")
	}
	tokenAddress, _ := tup.GetByInt64(2)
	tokenAddressInt, ok := tokenAddress.(value.IntValue)
	if !ok {
		return failRet, errors.New("tokenAddress must be int")
	}
	val, _ := tup.GetByInt64(3)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}
	blockNum, _ := tup.GetByInt64(4)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("blockNum must be int")
	}
	timestamp, _ := tup.GetByInt64(5)
	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("timestamp must be int")
	}
	messageNum, _ := tup.GetByInt64(6)
	messageNumInt, ok := messageNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("messageNum must be int")
	}

	return DeliveredERC721{
		ERC721: ERC721{
			To:           intValueToAddress(toInt),
			From:         intValueToAddress(fromInt),
			TokenAddress: intValueToAddress(tokenAddressInt),
			Id:           valInt.BigInt(),
		},
		BlockNum:   common.NewTimeBlocks(blockNumInt.BigInt()),
		Timestamp:  timestampInt.BigInt(),
		MessageNum: messageNumInt.BigInt(),
	}, nil
}
