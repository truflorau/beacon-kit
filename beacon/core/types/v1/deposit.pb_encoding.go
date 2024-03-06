// Code generated by fastssz. DO NOT EDIT.
// Hash: 9995182da6e0fccccfb2d6100f26c457b4d4b43b544e245f23bd9f30781739dd
package typesv1

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the Deposit object
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Deposit object to a target array
func (d *Deposit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(32)

	// Offset (0) 'ValidatorPubkey'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.ValidatorPubkey)

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 20 {
		err = ssz.ErrBytesLengthFn("--.WithdrawalCredentials", size, 20)
		return
	}
	dst = append(dst, d.WithdrawalCredentials...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, d.Amount)

	// Field (0) 'ValidatorPubkey'
	if size := len(d.ValidatorPubkey); size > 96 {
		err = ssz.ErrBytesLengthFn("--.ValidatorPubkey", size, 96)
		return
	}
	dst = append(dst, d.ValidatorPubkey...)

	return
}

// UnmarshalSSZ ssz unmarshals the Deposit object
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 32 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'ValidatorPubkey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 32 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'WithdrawalCredentials'
	if cap(d.WithdrawalCredentials) == 0 {
		d.WithdrawalCredentials = make([]byte, 0, len(buf[4:24]))
	}
	d.WithdrawalCredentials = append(d.WithdrawalCredentials, buf[4:24]...)

	// Field (2) 'Amount'
	d.Amount = ssz.UnmarshallUint64(buf[24:32])

	// Field (0) 'ValidatorPubkey'
	{
		buf = tail[o0:]
		if len(buf) > 96 {
			return ssz.ErrBytesLength
		}
		if cap(d.ValidatorPubkey) == 0 {
			d.ValidatorPubkey = make([]byte, 0, len(buf))
		}
		d.ValidatorPubkey = append(d.ValidatorPubkey, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit object
func (d *Deposit) SizeSSZ() (size int) {
	size = 32

	// Field (0) 'ValidatorPubkey'
	size += len(d.ValidatorPubkey)

	return
}

// HashTreeRoot ssz hashes the Deposit object
func (d *Deposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Deposit object with a hasher
func (d *Deposit) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ValidatorPubkey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.ValidatorPubkey))
		if byteLen > 96 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(d.ValidatorPubkey)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (96+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (96+31)/32)
		}
	}

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 20 {
		err = ssz.ErrBytesLengthFn("--.WithdrawalCredentials", size, 20)
		return
	}
	hh.PutBytes(d.WithdrawalCredentials)

	// Field (2) 'Amount'
	hh.PutUint64(d.Amount)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}