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

package arbbridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type IncomingMessageType int

const (
	CommonMessage IncomingMessageType = iota
	ChallengeMessage
)

type Event interface {
}

type VMEvent interface {
	GetIncomingMessageType() IncomingMessageType
}

type Notification struct {
	Header *types.Header
	VMID   common.Address
	Event  Event
	TxHash [32]byte
}

type StakeCreatedEvent struct {
	Staker   common.Address
	NodeHash [32]byte
}

func (e StakeCreatedEvent) RelatedToStaker(staker common.Address) bool {
	return staker == e.Staker
}

type ChallengeStartedEvent struct {
	Asserter          common.Address
	Challenger        common.Address
	ChallengeType     structures.ChildType
	ChallengeContract common.Address
}

func (e ChallengeStartedEvent) RelatedToStaker(staker common.Address) bool {
	return staker == e.Asserter || staker == e.Challenger
}

type ChallengeCompletedEvent struct {
	Winner            common.Address
	Loser             common.Address
	ChallengeContract common.Address
}

func (e ChallengeCompletedEvent) RelatedToStaker(staker common.Address) bool {
	return staker == e.Winner || staker == e.Loser
}

type StakeRefundedEvent struct {
	Staker common.Address
}

func (e StakeRefundedEvent) RelatedToStaker(staker common.Address) bool {
	return staker == e.Staker
}

type PrunedEvent struct {
	Leaf [32]byte
}

func (e PrunedEvent) RelatedToStaker(staker common.Address) bool {
	return false
}

type StakeMovedEvent struct {
	Staker   common.Address
	Location [32]byte
}

func (e StakeMovedEvent) RelatedToStaker(staker common.Address) bool {
	return staker == e.Staker
}

type AssertedEvent struct {
	PrevLeafHash  [32]byte
	Params        *structures.AssertionParams
	Claim         *structures.AssertionClaim
	MaxPendingTop [32]byte
}

func (e AssertedEvent) RelatedToStaker(staker common.Address) bool {
	return false
}

type ConfirmedEvent struct {
	NodeHash [32]byte
}

func (e ConfirmedEvent) RelatedToStaker(staker common.Address) bool {
	return false
}

type ConfirmedAssertionEvent struct {
	LogsAccHash [32]byte
}

type InitiateChallengeEvent struct {
	Deadline structures.TimeTicks
}

type AsserterTimeoutEvent struct{}

type ChallengerTimeoutEvent struct{}

type ContinueChallengeEvent struct {
	SegmentIndex *big.Int
	Deadline     structures.TimeTicks
}

type OneStepProofEvent struct{}

type PendingTopBisectionEvent struct {
	ChainHashes [][32]byte
	TotalLength *big.Int
	Deadline    structures.TimeTicks
}

type MessagesBisectionEvent struct {
	ChainHashes   [][32]byte
	SegmentHashes [][32]byte
	TotalLength   *big.Int
	Deadline      structures.TimeTicks
}

type ExecutionBisectionEvent struct {
	Assertions []*protocol.ExecutionAssertionStub
	TotalSteps uint32
	Deadline   structures.TimeTicks
}

type MessageDeliveredEvent struct {
	Msg protocol.Message
}

type NewTimeEvent struct{}