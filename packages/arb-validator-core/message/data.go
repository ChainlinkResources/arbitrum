/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"bytes"
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func bytesToValues(val []byte) []value.Value {
	var ints []value.Value
	for i := 0; i < len(val); i += 32 {
		remaining := len(val) - i
		if remaining < 32 {
			data := [32]byte{}
			copy(data[:], val[i:])
			ints = append(
				ints,
				value.NewIntValue(new(big.Int).SetBytes(data[:])),
			)
		} else {
			ints = append(
				ints,
				value.NewIntValue(new(big.Int).SetBytes(val[i:i+32])),
			)
		}
	}
	return ints
}

func StackValueToList(val value.Value) ([]value.Value, error) {
	values := make([]value.Value, 0)
	for !val.Equal(value.NewEmptyTuple()) {
		tupVal, ok := val.(value.TupleValue)
		if !ok {
			return nil, errors.New("value was not in stack format")
		}
		if tupVal.Len() != 2 {
			return nil, errors.New("stack expected to be 2-tuple")
		}
		member, err := tupVal.GetByInt64(1)
		if err != nil {
			return nil, err
		}
		values = append(values, member)
		val, err = tupVal.GetByInt64(0)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func ListToStackValue(vals []value.Value) value.TupleValue {
	ret := value.NewEmptyTuple()
	for _, val := range vals {
		ret = value.NewTuple2(ret, val)
	}
	return ret
}

func ByteStackToHex(val value.Value) ([]byte, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("bytestack expected tuple value")
	}
	if tup.Len() != 2 {
		return nil, errors.New("bytestack expected to be 2-tuple")
	}
	lengthVal, _ := tup.GetByInt64(0)
	lengthIntVal, ok := lengthVal.(value.IntValue)
	if !ok {
		return nil, errors.New("bytestack expected length to be int value")
	}
	intLength := lengthIntVal.BigInt().Uint64()

	stackVal, _ := tup.GetByInt64(1)
	tupVal, ok := stackVal.(value.TupleValue)
	if !ok {
		return nil, errors.New("bytestack expected 2 tuple value")
	}

	byteChunks := make([][32]byte, 0)
	vals, err := StackValueToList(tupVal)
	if err != nil {
		return nil, err
	}

	for _, val := range vals {
		intVal, ok := val.(value.IntValue)
		if !ok {
			return nil, errors.New("bytestack expected chunk to be int")
		}
		byteChunks = append(byteChunks, intVal.ToBytes())
	}

	var buf bytes.Buffer
	for i := range byteChunks {
		buf.Write(byteChunks[len(byteChunks)-1-i][:])
	}
	return buf.Bytes()[:intLength], nil
}

func BytesToByteStack(val []byte) value.TupleValue {
	chunks := bytesToValues(val)
	ret := ListToStackValue(chunks)
	return value.NewTuple2(value.NewInt64Value(int64(len(val))), ret)
}
