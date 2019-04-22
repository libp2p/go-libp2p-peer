// Deprecated: Use github.com/libp2p/go-libp2p-core/peer instead.
package peer

import (
	core "github.com/libp2p/go-libp2p-core/peer"
	ic "github.com/libp2p/go-libp2p-crypto"
)

var (
	// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ErrEmptyPeerID instead.
	ErrEmptyPeerID = core.ErrEmptyPeerID
	// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ErrNoPublicKey instead.
	ErrNoPublicKey = core.ErrNoPublicKey
)

// Deprecated: use github.com/libp2p/go-libp2p-core/peer.AdvanceEnableInlining instead.
// Warning: this variable's type makes it impossible to alias by reference.
// Reads and writes from/to this variable may be inaccurate or not have the intended effect.
var AdvancedEnableInlining = core.AdvancedEnableInlining

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ID instead.
type ID = core.ID

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDSlice instead.
type IDSlice = core.IDSlice

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromString instead.
func IDFromString(s string) (core.ID, error) {
	return core.IDFromString(s)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromBytes instead.
func IDFromBytes(b []byte) (core.ID, error) {
	return core.IDFromBytes(b)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDB58Decode instead.
func IDB58Decode(s string) (core.ID, error) {
	return core.IDB58Decode(s)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDB58Encode instead.
func IDB58Encode(id ID) string {
	return core.IDB58Encode(id)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDHexDecode instead.
func IDHexDecode(s string) (core.ID, error) {
	return core.IDHexDecode(s)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDHexEncode instead.
func IDHexEncode(id ID) string {
	return core.IDHexEncode(id)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromPublicKey instead.
func IDFromPublicKey(pk ic.PubKey) (core.ID, error) {
	return core.IDFromPublicKey(pk)
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromPrivateKey instead.
func IDFromPrivateKey(sk ic.PrivKey) (core.ID, error) {
	return core.IDFromPrivateKey(sk)
}
