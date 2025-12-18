package ucanfmt

import (
	"fmt"

	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/go-ucanto/core/dag/blockstore"
	"github.com/storacha/go-ucanto/core/delegation"
	"github.com/storacha/go-ucanto/core/ipld"
	"github.com/storacha/go-ucanto/core/receipt"
	"github.com/storacha/go-ucanto/core/result"
)

func PrintReceipt(r receipt.AnyReceipt) {
	doPrintReceipt(r, 0)
}

func doPrintReceipt(r receipt.AnyReceipt, level int) {
	log := withIndent(level)

	log("%s", r.Root().Link())

	if r.Issuer() != nil {
		log("  Issuer: %s", r.Issuer().DID())
	}

	log("  Ran: %s", r.Ran().Link())

	// Print the result (Out)
	out := r.Out()
	result.MatchResultR0(out, func(ok ipld.Node) {
		log("  Out: Ok")
		jsonString, err := ipldfmt.FormatNode(ok, "    ")
		if err != nil {
			log("    Error formatting JSON: %v", err)
		} else {
			log("%s", jsonString)
		}
	}, func(err ipld.Node) {
		log("  Out: Error")
		jsonString, jsonErr := ipldfmt.FormatNode(err, "    ")
		if jsonErr != nil {
			log("    Error formatting JSON: %v", jsonErr)
		} else {
			log("%s", jsonString)
		}
	})

	// Print effects
	fx := r.Fx()
	if len(fx.Fork()) > 0 {
		log("  Effects:")
		log("    Fork:")
		for _, f := range fx.Fork() {
			log("      %s", f.Link())
		}
	}

	if fx.Join().Link() != nil {
		if len(fx.Fork()) == 0 {
			log("  Effects:")
		}
		log("    Join: %s", fx.Join().Link())
	}

	// Print metadata
	meta := r.Meta()
	if len(meta) > 0 {
		log("  Meta:")
		for k, v := range meta {
			log("    %s: %v", k, v)
		}
	}

	bs, err := blockstore.NewBlockReader(blockstore.WithBlocksIterator(r.Blocks()))
	if err != nil {
		panic(fmt.Errorf("creating blockstore: %w", err))
	}

	// Print proofs recursively
	if len(r.Proofs()) > 0 {
		log("  Proofs:")
		for _, p := range r.Proofs() {
			pd, err := delegation.NewDelegationView(p.Link(), bs)
			if err != nil {
				log("    %s", p.Link())
				continue
			}
			doPrintDelegation(pd, level+2)
		}
	}
}
