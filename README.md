# debugger

Various tools and commands that can help debugging the Storacha Network.

## Install

```sh
go install github.com/storacha/debugger
```

## Usage

### `debugger delegation extract [car-file]`

Extract a delegation that has been archived to a CAR. You can pipe directly to this command.

e.g.

```sh
$ debugger delegation extract ucan.car
bafyreib5ygdak2sc6fd3coryjql6u4gcmjg7co5w2rbpvb6lqkqbnzehti
  Issuer: did:web:staging.up.storacha.network
  Audience: did:web:staging.indexer.storacha.network
  Capabilities:
    Can: assert/index
    With: did:web:staging.indexer.storacha.network
    Nb: &{map[content:0x140003497e0 index:0x140003497c0] [{index 0x140003497c0} {content 0x140003497e0}]}
  Proofs:
    bafyreifb4tz7tup4b5s6kie4adqbhi2oisqx3t5d6jxlgpxnpcdykoklpu
      Issuer: did:web:staging.indexer.storacha.network
      Audience: did:web:staging.up.storacha.network
      Capabilities:
        Can: assert/equals
        With: did:web:staging.indexer.storacha.network
        Nb: &{map[] []}
        Can: assert/index
        With: did:web:staging.indexer.storacha.network
        Nb: &{map[] []}
```

### `debugger delegation parse <value>`

Parse and print a UCAN delegation.

### `debugger ipni metadata parse <value>`

Parse a base64 encoded IPNI metadata.

e.g.

```sh
$ debugger ipni metadata parse gID4AaNhY9gqWCUAAXESID3BhgVqQvFHsTo4TBfqcMJiTfE7ttRC+ofLgqAW5IeaYWUAYWnYKlgmAAGCBBIgUX8wwCPTHNNAl3QqAdrM+j5wl47D87QWT+Ps3ZlKQ4Y=
ID: 0x3e0000 (index claim)
Claim: bafyreib5ygdak2sc6fd3coryjql6u4gcmjg7co5w2rbpvb6lqkqbnzehti
Index: bagbaierakf7tbqbd2mongqexoqvadwwm7i7hbf4oypz3ifsp4pwn3gkkioda
Expiration: 1970-01-01 01:00:00 +0100 BST
```

### `debugger retrieve <url> <auth>`

Attempt to retrieve data from the passed URL using the provided authorization, which is expected to be an `X-Agent-Message` header.

Note: when providing auth as a `X-Agent-Message`, it is expected to be gzipped and multibase encoded.

### `debugger xagentmessage parse <value>`

Parse a gzipped and multibase encoded `X-Agent-Message` header.
