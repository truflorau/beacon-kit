// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package core

import (
	"fmt"

	"github.com/berachain/beacon-kit/mod/core/state"
	"github.com/berachain/beacon-kit/mod/core/types"
	"github.com/berachain/beacon-kit/mod/primitives"
)

// BlockVerifier is responsible for verifying incoming BeaconBlocks.
type BlockVerifier struct {
	cs primitives.ChainSpec
}

// NewBlockVerifier creates a new block validator.
func NewBlockVerifier(cs primitives.ChainSpec) *BlockVerifier {
	return &BlockVerifier{
		cs: cs,
	}
}

// ValidateBlock validates the incoming block.
func (bv *BlockVerifier) ValidateBlock(
	st state.BeaconState,
	blk primitives.ReadOnlyBeaconBlock,
) error {
	// Get the block body.
	body := blk.GetBody()
	if body == nil || body.IsNil() {
		return types.ErrNilBlkBody
	}

	// Get the current slot.
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	// Ensure the block slot matches the state slot.
	if blk.GetSlot() != slot {
		return fmt.Errorf(
			"slot does not match, expected: %d, got: %d",
			slot,
			blk.GetSlot(),
		)
	}

	// Get the latest block header.
	latestBlockHeader, err := st.GetLatestBlockHeader()
	if err != nil {
		return err
	}

	// Ensure the block is within the acceptable range.
	if blk.GetSlot() <= latestBlockHeader.GetSlot() {
		return fmt.Errorf(
			"block slot is too low, expected: > %d, got: %d",
			latestBlockHeader.Slot,
			blk.GetSlot(),
		)
	}

	// Ensure the block is within the acceptable range.
	// TODO: move this is in the wrong spot.
	if deposits := body.GetDeposits(); uint64(
		len(deposits),
	) > bv.cs.MaxDepositsPerBlock() {
		return fmt.Errorf(
			"too many deposits, expected: %d, got: %d",
			bv.cs.MaxDepositsPerBlock(), len(deposits),
		)
	}

	// Ensure the parent root matches the latest block header.
	parentBlockRoot, err := latestBlockHeader.HashTreeRoot()
	if err != nil {
		return err
	}

	// Ensure the parent root matches the latest block header.
	if parentBlockRoot != blk.GetParentBlockRoot() {
		return fmt.Errorf(
			"parent root does not match, expected: %x, got: %x",
			parentBlockRoot,
			blk.GetParentBlockRoot(),
		)
	}
	return nil
}