// Copyright 2020 Wintex
// SPDX-License-Identifier: LGPL-3.0-only

package ed25519

import (
	"github.com/radianceteam/ton-client-go/client"
	"github.com/wintexpro/chainbridge-utils/crypto"
)

var _ crypto.Keypair = &Keypair{}

type Keypair struct {
	public   []byte
	secret   []byte
	mnemonic string
}

func newClient() (*client.Client, error) {
	return client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: ""},
	})
}

func GenerateKeypair() (*Keypair, error) {
	c, err := newClient()
	defer c.Close()

	if err != nil {
		return nil, err
	}

	res, err := c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{})
	kp, err := NewKeypairFromSeed(res.Phrase)

	return kp, err
}

func NewKeypairFromSeed(seed string) (*Keypair, error) {
	c, err := newClient()
	defer c.Close()

	if err != nil {
		return nil, err
	}

	mnemonic := client.ParamsOfMnemonicDeriveSignKeys{Phrase: seed}
	res, err := c.CryptoMnemonicDeriveSignKeys(&mnemonic)

	kp := Keypair{
		public:   []byte(res.Public),
		secret:   []byte(res.Secret),
		mnemonic: seed,
	}

	return &kp, err
}

// Encode uses scale to encode underlying KeyringPair
func (kp *Keypair) Encode() []byte {
	return []byte(kp.mnemonic)
}

// Decode initializes keypair by decoding input as a KeyringPair
func (kp *Keypair) Decode(in []byte) error {
	newkp, err := NewKeypairFromSeed(string(in))

	kp.public = newkp.public
	kp.secret = newkp.secret
	kp.mnemonic = newkp.mnemonic

	return err
}

// Address returns the ss58 formated address
func (kp *Keypair) Address() string {
	return ""
}

// PublicKey returns the publickey encoded as a string
func (kp *Keypair) PublicKey() string {
	return string(kp.public)
}

// SecretKey returns the secretkey encoded as a string
func (kp *Keypair) SecretKey() string {
	return string(kp.secret)
}
