package ucanfmt

import (
	"fmt"
	"strings"
	"time"

	"github.com/storacha/go-ucanto/core/dag/blockstore"
	"github.com/storacha/go-ucanto/core/delegation"
)

func withIndent(level int) func(format string, args ...any) {
	indent := strings.Repeat("  ", level)
	return func(format string, args ...any) {
		fmt.Printf(indent+format+"\n", args...)
	}
}

func PrintDelegation(d delegation.Delegation) {
	doPrintDelegation(d, 0)
}

func doPrintDelegation(d delegation.Delegation, level int) {
	log := withIndent(level)

	log("%s", d.Link())
	log("  Issuer: %s", d.Issuer().DID())
	log("  Audience: %s", d.Audience().DID())

	log("  Capabilities:")
	for _, c := range d.Capabilities() {
		log("    Can: %s", c.Can())
		log("    With: %s", c.With())
		log("    Nb: %v", c.Nb())
	}

	if d.Expiration() != nil {
		log("  Expiration: %s", time.Unix(int64(*d.Expiration()), 0).String())
	}

	bs, err := blockstore.NewBlockReader(blockstore.WithBlocksIterator(d.Blocks()))
	if err != nil {
		panic(fmt.Errorf("creating blockstore: %w", err))
	}

	if len(d.Proofs()) > 0 {
		log("  Proofs:")
		for _, p := range d.Proofs() {
			pd, err := delegation.NewDelegationView(p, bs)
			if err != nil {
				log("    %s\n", p)
				continue
			}
			doPrintDelegation(pd, level+2)
		}
	}
}
