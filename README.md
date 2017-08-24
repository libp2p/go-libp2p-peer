# go-libp2p-peer

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/ipfs/go-libp2p-peer?status.svg)](https://godoc.org/github.com/ipfs/go-libp2p-peer)
[![Coverage Status](https://coveralls.io/repos/github/ipfs/go-libp2p-peer/badge.svg?branch=master)](https://coveralls.io/github/ipfs/go-libp2p-peer?branch=master)
[![Build Status](https://travis-ci.org/ipfs/go-libp2p-peer.svg?branch=master)](https://travis-ci.org/ipfs/go-libp2p-peer)

> PKI based identities for use in go-libp2p

## Install

TODO

## Usage

Generate a peers ID:

```go
import (
        "crypto/rand"

        crypto "github.com/libp2p/go-libp2p-crypto"
        peer "github.com/libp2p/go-libp2p-peer"
)

// Generate an identity keypair using go's cryptographic randomness source
priv, pub, err := crypto.GenerateEd25519Key(rand.Reader)
if err != nil { panic(err) }

// A peers ID is the hash of its public key
pid, err := peer.IDFromPublicKey(pub)
if err != nil { panic(err) }
```

Go to https://godoc.org/github.com/libp2p/go-libp2p-peer

## Contribute

Feel free to join in. All welcome. Open an [issue](https://github.com/ipfs/go-libp2p-peer/issues)!

This repository falls under the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

### Want to hack on IPFS?

[![](https://cdn.rawgit.com/jbenet/contribute-ipfs-gif/master/img/contribute.gif)](https://github.com/ipfs/community/blob/master/contributing.md)

## License

MIT
