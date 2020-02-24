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

package gobridge

import (
	"context"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var messagesBisectedID ethcommon.Hash
var messagesOneStepProofCompletedID ethcommon.Hash

type messagesChallengeWatcher struct {
	*bisectionChallengeWatcher
	challengeInfo *challengeData
	client        *GoArbClient
}

func newMessagesChallengeWatcher(address common.Address, client *GoArbClient) (*messagesChallengeWatcher, error) {
	fmt.Println("in newMessagesChallengeWatcher")
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}

	chalData := client.GoEthClient.challenges[address]
	client.GoEthClient.challengeWatchersMutex.Lock()
	if _, ok := client.GoEthClient.challengeWatcherEvents[chalData]; !ok {
		client.GoEthClient.challengeWatcherEvents[chalData] = make(map[*common.BlockId][]arbbridge.Event)
	}
	client.GoEthClient.challengeWatchersMutex.Unlock()

	return &messagesChallengeWatcher{bisectionChallengeWatcher: bisectionChallenge, challengeInfo: chalData, client: client}, nil
}

func (c *messagesChallengeWatcher) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	fmt.Println("in messagesChallengeWatcher GetEvents")

	c.client.GoEthClient.challengeWatchersMutex.Lock()
	cw := c.client.GoEthClient.challengeWatcherEvents[c.challengeInfo][blockId]
	c.client.GoEthClient.challengeWatchersMutex.Unlock()
	return cw, nil
}

func (c *messagesChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		messagesBisectedID,
		messagesOneStepProofCompletedID,
	}
	return append(tops, c.bisectionChallengeWatcher.topics()...)
}