package testutil

import (
	"github.com/libp2p/go-libp2p-core/peer"
	tpeer "github.com/libp2p/go-libp2p-core/peer/test"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	tcrypto "github.com/libp2p/go-libp2p-core/crypto/test"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/peer/test.RandPeerID instead.
func RandPeerID() (peer.ID, error) {
	return tpeer.RandPeerID()
}

// Deprecated: use github.com/libp2p/go-libp2p-core/crypto/test.RandTestKeyPair instead.
// Supply RSA as a key type to get an equivalent result.
func RandTestKeyPair(bits int) (ci.PrivKey, ci.PubKey, error) {
	return tcrypto.RandTestKeyPair(ci.RSA, bits)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/crypto/test.SeededTestKeyPair instead.
// Supply RSA as a key type, with 512 bits, to get an equivalent result.
func SeededTestKeyPair(seed int64) (ci.PrivKey, ci.PubKey, error) {
	return tcrypto.SeededTestKeyPair(ci.RSA, 512, seed)
}
