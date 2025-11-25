# debugger

Various tools and commands that can help debugging the Storacha Network.

## Install

```sh
go install github.com/storacha/debugger
```

## Usage

### `debugger parse xagentmessage <value>`

Parse a gzipped and multibase encoded `X-Agent-Message` header.

### `debugger retrieve <url> <auth>`

Attempt to retrieve data from the passed URL using the provided authorization, which is expected to be an `X-Agent-Message` header.

Note: when providing auth as a `X-Agent-Message`, it is expected to be gzipped and multibase encoded.
