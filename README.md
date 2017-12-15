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

TODO

## Format

An `ID` is a [multihash](https://github.com/multiformats/multihash) with type either *SHA2_256* or *identity*.

For most key types the type is *SHA2_256* and the payload is the value of `github.com/libp2p/go-libp2p-crypto.PubKey.Bytes()`.

```
<SHA2_256 mc><length (32)><SHA256 hash>
<0x12       ><0x20       ><SHA256 hash>
```

Ed25519 public keys instead are encoded with type *identity* and as payload a [multicodec](https://github.com/multiformats/multicodec) key in packed format.

```
<identity mc><length (2 + 32 = 34)><ed25519-pub mc><ed25519 pubkey>
<0x00       ><0x22                ><0xed01        ><ed25519 pubkey>
```

## Contribute

Feel free to join in. All welcome. Open an [issue](https://github.com/ipfs/go-libp2p-peer/issues)!

This repository falls under the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

### Want to hack on IPFS?

[![](https://cdn.rawgit.com/jbenet/contribute-ipfs-gif/master/img/contribute.gif)](https://github.com/ipfs/community/blob/master/contributing.md)

## License

MIT
