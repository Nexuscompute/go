// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hmac

import (
	"bytes"
	"crypto/internal/fips"
	"crypto/internal/fips/sha256"
	"errors"
)

func init() {
	fips.CAST("HMAC-SHA2-256", func() error {
		input := []byte{
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
			0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
		}
		want := []byte{
			0xf0, 0x8d, 0x82, 0x8d, 0x4c, 0x9e, 0xad, 0x3d,
			0xdc, 0x12, 0x9c, 0x4e, 0x70, 0xc4, 0x19, 0x2a,
			0x4f, 0x12, 0x73, 0x23, 0x73, 0x77, 0x66, 0x05,
			0x10, 0xee, 0x57, 0x6b, 0x3a, 0xc7, 0x14, 0x41,
		}
		h := New(sha256.New, input)
		h.Write(input)
		h.Write(input)
		if got := h.Sum(nil); !bytes.Equal(got, want) {
			return errors.New("unexpected result")
		}
		return nil
	})
}
