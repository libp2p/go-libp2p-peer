// package peer implements an object used to represent peers in the ipfs network.
package peer

import (
	"encoding/hex"
	"fmt"
	"strings"

	logging "github.com/ipfs/go-log" // ID represents the identity of a peer.
	b58 "github.com/jbenet/go-base58"
	ic "github.com/libp2p/go-libp2p-crypto"
	mc "github.com/multiformats/go-multicodec-packed"
	mh "github.com/multiformats/go-multihash"
)

var log = logging.Logger("peer")

type ID string

// Pretty returns a b58-encoded string of the ID
func (id ID) Pretty() string {
	return IDB58Encode(id)
}

func (id ID) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"peerID": id.Pretty(),
	}
}

// String prints out the peer.
//
// TODO(brian): ensure correctness at ID generation and
// enforce this by only exposing functions that generate
// IDs safely. Then any peer.ID type found in the
// codebase is known to be correct.
func (id ID) String() string {
	pid := id.Pretty()

	//All sha256 nodes start with Qm
	//We can skip the Qm to make the peer.ID more useful
	if strings.HasPrefix(pid, "Qm") {
		pid = pid[2:]
	}

	maxRunes := 6
	if len(pid) < maxRunes {
		maxRunes = len(pid)
	}
	return fmt.Sprintf("<peer.ID %s>", pid[:maxRunes])
}

// MatchesPrivateKey tests whether this ID was derived from sk
func (id ID) MatchesPrivateKey(sk ic.PrivKey) bool {
	return id.MatchesPublicKey(sk.GetPublic())
}

// MatchesPublicKey tests whether this ID was derived from pk
func (id ID) MatchesPublicKey(pk ic.PubKey) bool {
	oid, err := IDFromPublicKey(pk)
	if err != nil {
		return false
	}
	return oid == id
}

func (id ID) ExtractEd25519PublicKey() ic.PubKey {
	// ed25519 pubkey identity format
	// <identity mc><length (2 + 32 = 34)><ed25519-pub mc><ed25519 pubkey>
	// <0x00       ><0x22                ><0xed01        ><ed25519 pubkey>

	var nilPubKey ic.PubKey

	// Decode multihash
	decoded, err := mh.Decode([]byte(id))
	if err != nil {
		return nilPubKey
	}

	// Check ID multihash codec
	if decoded.Code != mh.ID {
		return nilPubKey
	}

	// Check multihash length
	if decoded.Length != 2+32 {
		return nilPubKey
	}

	// Split prefix
	code, pubKeyBytes := mc.SplitPrefix(decoded.Digest)

	// Check ed25519 code
	if code != mc.Ed25519Pub {
		return nilPubKey
	}

	// Unmarshall public key
	pubKey, err := ic.UnmarshalEd25519PublicKey(pubKeyBytes)
	if err != nil {
		return nilPubKey
	}

	return pubKey
}

func (id ID) ExtractPublicKey() ic.PubKey {
	var pk ic.PubKey

	// Try extract ed25519 pubkey
	pk = id.ExtractEd25519PublicKey()
	if pk != nil {
		return pk
	}

	// Try extract other type of pubkey
	/*pk = id.Extract...PublicKey()
	if pk != nil {
		return pk
	}*/

	return pk
}

// IDFromString cast a string to ID type, and validate
// the id to make sure it is a multihash.
func IDFromString(s string) (ID, error) {
	if _, err := mh.Cast([]byte(s)); err != nil {
		return ID(""), err
	}
	return ID(s), nil
}

// IDFromBytes cast a string to ID type, and validate
// the id to make sure it is a multihash.
func IDFromBytes(b []byte) (ID, error) {
	if _, err := mh.Cast(b); err != nil {
		return ID(""), err
	}
	return ID(b), nil
}

// IDB58Decode returns a b58-decoded Peer
func IDB58Decode(s string) (ID, error) {
	m, err := mh.FromB58String(s)
	if err != nil {
		return "", err
	}
	return ID(m), err
}

// IDB58Encode returns b58-encoded string
func IDB58Encode(id ID) string {
	return b58.Encode([]byte(id))
}

// IDHexDecode returns a b58-decoded Peer
func IDHexDecode(s string) (ID, error) {
	m, err := mh.FromHexString(s)
	if err != nil {
		return "", err
	}
	return ID(m), err
}

// IDHexEncode returns b58-encoded string
func IDHexEncode(id ID) string {
	return hex.EncodeToString([]byte(id))
}

// IDFromPublicKey returns the Peer ID corresponding to pk
func IDFromPublicKey(pk ic.PubKey) (ID, error) {
	b, err := pk.Bytes()
	if err != nil {
		return "", err
	}
	hash, _ := mh.Sum(b, mh.SHA2_256, -1)
	return ID(hash), nil
}

// IDFromPrivateKey returns the Peer ID corresponding to sk
func IDFromPrivateKey(sk ic.PrivKey) (ID, error) {
	return IDFromPublicKey(sk.GetPublic())
}

// IDSlice for sorting peers
type IDSlice []ID

func (es IDSlice) Len() int           { return len(es) }
func (es IDSlice) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es IDSlice) Less(i, j int) bool { return string(es[i]) < string(es[j]) }
