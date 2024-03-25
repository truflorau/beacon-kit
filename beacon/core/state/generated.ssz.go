// Code generated by fastssz. DO NOT EDIT.
// Hash: f489332ce0643c2a8c59ca054d0091f678af50951607702647807e10fd83e669
// Version: 0.1.3
package state

import (
	"github.com/berachain/beacon-kit/beacon/core/types"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconStateDeneb object
func (b *BeaconStateDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconStateDeneb object to a target array
func (b *BeaconStateDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(100)

	// Field (0) 'GenesisValidatorsRoot'
	dst = append(dst, b.GenesisValidatorsRoot[:]...)

	// Field (1) 'Eth1GenesisHash'
	dst = append(dst, b.Eth1GenesisHash[:]...)

	// Offset (2) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 89

	// Field (3) 'RandaoMix'
	if size := len(b.RandaoMix); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconStateDeneb.RandaoMix", size, 32)
		return
	}
	dst = append(dst, b.RandaoMix...)

	// Field (2) 'Validators'
	if size := len(b.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconStateDeneb.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconStateDeneb object
func (b *BeaconStateDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 100 {
		return ssz.ErrSize
	}

	tail := buf
	var o2 uint64

	// Field (0) 'GenesisValidatorsRoot'
	copy(b.GenesisValidatorsRoot[:], buf[0:32])

	// Field (1) 'Eth1GenesisHash'
	copy(b.Eth1GenesisHash[:], buf[32:64])

	// Offset (2) 'Validators'
	if o2 = ssz.ReadOffset(buf[64:68]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 100 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (3) 'RandaoMix'
	if cap(b.RandaoMix) == 0 {
		b.RandaoMix = make([]byte, 0, len(buf[68:100]))
	}
	b.RandaoMix = append(b.RandaoMix, buf[68:100]...)

	// Field (2) 'Validators'
	{
		buf = tail[o2:]
		num, err := ssz.DivideInt2(len(buf), 89, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*types.Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(types.Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*89 : (ii+1)*89]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconStateDeneb object
func (b *BeaconStateDeneb) SizeSSZ() (size int) {
	size = 100

	// Field (2) 'Validators'
	size += len(b.Validators) * 89

	return
}

// HashTreeRoot ssz hashes the BeaconStateDeneb object
func (b *BeaconStateDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconStateDeneb object with a hasher
func (b *BeaconStateDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisValidatorsRoot'
	hh.PutBytes(b.GenesisValidatorsRoot[:])

	// Field (1) 'Eth1GenesisHash'
	hh.PutBytes(b.Eth1GenesisHash[:])

	// Field (2) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (3) 'RandaoMix'
	if size := len(b.RandaoMix); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconStateDeneb.RandaoMix", size, 32)
		return
	}
	hh.PutBytes(b.RandaoMix)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconStateDeneb object
func (b *BeaconStateDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}