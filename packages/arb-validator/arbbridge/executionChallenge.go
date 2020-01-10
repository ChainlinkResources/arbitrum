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
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type ExecutionChallenge interface {
	BisectionChallenge
	BisectAssertion(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertions []*protocol.ExecutionAssertionStub,
		totalSteps uint32,
	) error
	OneStepProof(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertionStub,
		proof []byte,
	) error
	ExecutionChallengeChooseSegment(
		ctx context.Context,
		assertionToChallenge uint16,
		preconditions []*protocol.Precondition,
		assertions []*protocol.ExecutionAssertionStub,
	) error
}