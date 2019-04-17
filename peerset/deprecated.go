// Deprecated: Use github.com/libp2p/go-libp2p-core/peer/set instead.
package peerset

import moved "github.com/libp2p/go-libp2p-core/peer/set"

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer/set.PeerSet instead.
type PeerSet = moved.PeerSet

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer/set.New instead.
func New() *PeerSet {
	return moved.New()
}

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer/set.NewLimited instead.
func NewLimited(size int) *PeerSet {
	return moved.NewLimited(size)
}
