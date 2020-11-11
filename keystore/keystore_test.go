package keystore

import (
	"testing"
)

var path string = "/Users/by-keks/workspace/projects/go/src/github.com/wintexpro/chainbridge-utils/keys"

func TestKeypairFromAddress(t *testing.T) {
	kp, err := KeypairFromAddress("0:efbbf6b8379b529c0f51a72a85bf443d536f8c71db7310d5a612746df57903af", TonChain, path, false)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", kp)

	// ring := makeTonRing()

	// t.Logf("%v", ring[AliceKey])
}
