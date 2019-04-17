// Deprecated: Use github.com/libp2p/go-libp2p-core/peer/set instead.
package peerset

import moved "github.com/libp2p/go-libp2p-core/peer"

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.Set instead.
type PeerSet = moved.Set

// Deprecated: Use github.com/libp2p/go-libp2p-core/set.NewSet instead.
func New() *PeerSet {
	return moved.NewSet()
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer/set.NewLimitedSet instead.
func NewLimited(size int) *PeerSet {
	return moved.NewLimitedSet(size)
}
