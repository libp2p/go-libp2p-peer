// Deprecated: Use github.com/libp2p/go-libp2p-core/peer instead.
package peer

import (
	moved "github.com/libp2p/go-libp2p-core/peer"
	ic "github.com/libp2p/go-libp2p-crypto"
)

var (
	// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ErrEmptyPeerID instead.
	ErrEmptyPeerID = moved.ErrEmptyPeerID
	// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ErrNoPublicKey instead.
	ErrNoPublicKey = moved.ErrNoPublicKey
)

// Deprecated: use github.com/libp2p/go-libp2p-core/peer.AdvanceEnableInlining instead.
// Warning: this variable's type makes it impossible to alias by reference.
// Reads and writes from/to this variable may be inaccurate or not have the intended effect.
var AdvancedEnableInlining = moved.AdvancedEnableInlining

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ID instead.
type ID = moved.ID

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDSlice instead.
type IDSlice = moved.IDSlice

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromString instead.
func IDFromString(s string) (moved.ID, error) {
	return moved.IDFromString(s)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromBytes instead.
func IDFromBytes(b []byte) (moved.ID, error) {
	return moved.IDFromBytes(b)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDB58Decode instead.
func IDB58Decode(s string) (moved.ID, error) {
	return moved.IDB58Decode(s)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDB58Encode instead.
func IDB58Encode(id ID) string {
	return moved.IDB58Encode(id)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDHexDecode instead.
func IDHexDecode(s string) (moved.ID, error) {
	return moved.IDHexDecode(s)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDHexEncode instead.
func IDHexEncode(id ID) string {
	return moved.IDHexEncode(id)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromPublicKey instead.
func IDFromPublicKey(pk ic.PubKey) (moved.ID, error) {
	return moved.IDFromPublicKey(pk)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromPrivateKey instead.
func IDFromPrivateKey(sk ic.PrivKey) (moved.ID, error) {
	return moved.IDFromPrivateKey(sk)
}
