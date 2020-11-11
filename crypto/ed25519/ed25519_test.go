// Copyright 2020 Wintex
// SPDX-License-Identifier: LGPL-3.0-only

package ed25519

import (
	"reflect"
	"testing"
)

func TestNewKeypair(t *testing.T) {
	kp, err := NewKeypairFromSeed("action glow era all liquid critic achieve lawsuit era anger loud slight")

	if err != nil {
		t.Fatal(err)
	}

	if kp.PublicKey() == "" {
		t.Fatalf("key is missing data: %#v", kp)
	}

	if kp.PublicKey() != "ebc77aae202a4f12237e10892f4fe0e44f8fb3dfc07008dcc12b37f8f70c1149" {
		t.Fatalf("pubkey is not correct: %#v", kp)
	}
}

func TestEncodeAndDecodeKeypair(t *testing.T) {
	kp, err := GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	enc := kp.Encode()
	res := new(Keypair)
	err = res.Decode(enc)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res, kp) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}
