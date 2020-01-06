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

package ethbridge

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type ClientConnection struct {
	Client *ethclient.Client
}

func (c *ClientConnection) CurrentBlockTime(ctx context.Context) (*protocol.TimeBlocks, error) {
	header, err := c.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return protocol.NewTimeBlocks(header.Number), nil
}

func (c *ClientConnection) waitForReceipt(ctx context.Context, from common.Address, tx *types.Transaction, methodName string) (*types.Receipt, error) {
	return waitForReceipt(ctx, c.Client, from, tx, methodName)
}