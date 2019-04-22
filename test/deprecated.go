package testutil

import (
	"github.com/libp2p/go-libp2p-core/peer"
	core "github.com/libp2p/go-libp2p-core/peer/test"
	ci "github.com/libp2p/go-libp2p-crypto"
)

func RandPeerID() (peer.ID, error) {
	return core.RandPeerID()
}

func RandTestKeyPair(bits int) (ci.PrivKey, ci.PubKey, error) {
	return core.RandTestKeyPair(bits)
}

func SeededTestKeyPair(seed int64) (ci.PrivKey, ci.PubKey, error) {
	return core.SeededTestKeyPair(seed)
}
