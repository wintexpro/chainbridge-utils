// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package secp256k1

import (
	"testing"
)

func TestNewKeypairFromSeed(t *testing.T) {
	kp, err := GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	if kp.PublicKey() == "" || kp.Address() == "" {
		t.Fatalf("key is missing data: %#v", kp)
	}
}

func TestEncodeAndDecodeKeypair(t *testing.T) {
	// 	kp, err := GenerateKeypair()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	t.Logf("Address: %v", kp.Address())
	// 	t.Logf("PublicKey: %v", kp.PublicKey())
	// 	t.Logf("PrivateKey: %v", kp.PrivateKey())

	// 	enc := kp.Encode()
	// 	t.Logf("enc: %v", enc)
	// 	res := new(Keypair)
	// 	err = res.Decode(enc)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	if !reflect.DeepEqual(res, kp) {
	// 		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	// 	}
}
