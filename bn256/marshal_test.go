// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gurvy DO NOT EDIT

package bn256

import (
	"bytes"
	"io"
	"math/big"
	"math/rand"
	"testing"

	"github.com/consensys/gurvy/bn256/fp"
	"github.com/consensys/gurvy/bn256/fr"
)

func TestEncoder(t *testing.T) {

	// TODO need proper fuzz testing here

	var inA uint64
	var inB fr.Element
	var inC fp.Element
	var inD G1
	var inE G1
	var inF G2
	var inG []G1
	var inH []G2

	// set values of inputs
	inA = rand.Uint64()
	inB.SetRandom()
	inC.SetRandom()
	inD.ScalarMultiplication(&g1GenAff, new(big.Int).SetUint64(rand.Uint64()))
	// inE --> infinity
	inF.ScalarMultiplication(&g2GenAff, new(big.Int).SetUint64(rand.Uint64()))
	inG = make([]G1, 2)
	inH = make([]G2, 0)
	inG[1] = inD

	// encode them, compressed and raw
	var buf, bufRaw bytes.Buffer
	enc := NewEncoder(&buf)
	encRaw := NewEncoder(&bufRaw, RawEncoding())
	toEncode := []interface{}{inA, &inB, &inC, &inD, &inE, &inF, inG, inH}
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			t.Fatal(err)
		}
		if err := encRaw.Encode(v); err != nil {
			t.Fatal(err)
		}
	}

	testDecode := func(t *testing.T, r io.Reader, n int64) {
		dec := NewDecoder(r)
		var outA uint64
		var outB fr.Element
		var outC fp.Element
		var outD G1
		var outE G1
		outE.X.SetOne()
		outE.Y.SetUint64(42)
		var outF G2
		var outG []G1
		var outH []G2

		toDecode := []interface{}{&outA, &outB, &outC, &outD, &outE, &outF, &outG, &outH}
		for _, v := range toDecode {
			if err := dec.Decode(v); err != nil {
				t.Fatal(err)
			}
		}

		// compare values
		if inA != outA {
			t.Fatal("didn't encode/decode uint64 value properly")
		}

		if !inB.Equal(&outB) || !inC.Equal(&outC) {
			t.Fatal("decode(encode(Element) failed")
		}
		if !inD.Equal(&outD) || !inE.Equal(&outE) {
			t.Fatal("decode(encode(G1) failed")
		}
		if !inF.Equal(&outF) {
			t.Fatal("decode(encode(G2) failed")
		}
		if (len(inG) != len(outG)) || (len(inH) != len(outH)) {
			t.Fatal("decode(encode(slice(points))) failed")
		}
		for i := 0; i < len(inG); i++ {
			if !inG[i].Equal(&outG[i]) {
				t.Fatal("decode(encode(slice(points))) failed")
			}
		}
		if n != dec.BytesRead() {
			t.Fatal("bytes read don't match bytes written")
		}
	}

	// decode them
	testDecode(t, &buf, enc.BytesWritten())
	testDecode(t, &bufRaw, encRaw.BytesWritten())

}

func TestIsCompressed(t *testing.T) {
	var g1Inf, g1 G1
	var g2Inf, g2 G2

	g1 = g1GenAff
	g2 = g2GenAff

	{
		b := g1Inf.Bytes()
		if !isCompressed(b[0]) {
			t.Fatal("g1Inf.Bytes() should be compressed")
		}
	}

	{
		b := g1Inf.RawBytes()
		if isCompressed(b[0]) {
			t.Fatal("g1Inf.RawBytes() should be uncompressed")
		}
	}

	{
		b := g1.Bytes()
		if !isCompressed(b[0]) {
			t.Fatal("g1.Bytes() should be compressed")
		}
	}

	{
		b := g1.RawBytes()
		if isCompressed(b[0]) {
			t.Fatal("g1.RawBytes() should be uncompressed")
		}
	}

	{
		b := g2Inf.Bytes()
		if !isCompressed(b[0]) {
			t.Fatal("g2Inf.Bytes() should be compressed")
		}
	}

	{
		b := g2Inf.RawBytes()
		if isCompressed(b[0]) {
			t.Fatal("g2Inf.RawBytes() should be uncompressed")
		}
	}

	{
		b := g2.Bytes()
		if !isCompressed(b[0]) {
			t.Fatal("g2.Bytes() should be compressed")
		}
	}

	{
		b := g2.RawBytes()
		if isCompressed(b[0]) {
			t.Fatal("g2.RawBytes() should be uncompressed")
		}
	}

}
