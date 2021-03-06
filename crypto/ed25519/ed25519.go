// Copyright 2020 Wintex
// SPDX-License-Identifier: LGPL-3.0-only

package ed25519

import (
	"github.com/radianceteam/ton-client-go/client"
	"github.com/volatiletech/null"
	"github.com/wintexpro/chainbridge-utils/crypto"
)

var _ crypto.Keypair = &Keypair{}

type Keypair struct {
	public  []byte
	secret  []byte
	address string
}

func newClient() (*client.Client, error) {
	return client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: null.StringFrom("")},
	})
}

func GenerateKeypair() (*Keypair, error) {
	c, err := newClient()
	defer c.Close()

	if err != nil {
		return nil, err
	}

	res, err := c.CryptoGenerateRandomSignKeys()

	kp := Keypair{
		public: []byte(res.Public),
		secret: []byte(res.Secret),
	}

	return &kp, err
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
		public: []byte(res.Public),
		secret: []byte(res.Secret),
	}

	return &kp, err
}

// Encode uses scale to encode underlying KeyringPair
func (kp *Keypair) Encode() []byte {
	return []byte(kp.secret)
}

// Decode initializes keypair by decoding input as a KeyringPair
func (kp *Keypair) Decode(in []byte) error {
	c, err := newClient()
	defer c.Close()

	if err != nil {
		return err
	}

	params := client.ParamsOfNaclSignKeyPairFromSecret{
		Secret: string(in),
	}

	keypair, err := c.CryptoNaclSignKeypairFromSecretKey(&params)

	kp.public = []byte(keypair.Public)
	kp.secret = []byte(keypair.Secret[0:len(keypair.Public)])

	return err
}

func (kp *Keypair) SetAddress(address string) {
	kp.address = address
}

// Address returns the setted address
func (kp *Keypair) Address() string {
	return kp.address
}

// PublicKey returns the publickey encoded as a string
func (kp *Keypair) PublicKey() string {
	return string(kp.public)
}

// SecretKey returns the secretkey encoded as a string
func (kp *Keypair) SecretKey() string {
	return string(kp.secret)
}
