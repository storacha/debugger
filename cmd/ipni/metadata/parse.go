package metadata

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	logging "github.com/ipfs/go-log/v2"
	ipnimd "github.com/ipni/go-libipni/metadata"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/spf13/cobra"
	"github.com/storacha/go-libstoracha/capabilities/assert"
	"github.com/storacha/go-libstoracha/digestutil"
	"github.com/storacha/go-libstoracha/metadata"
	"github.com/storacha/go-ucanto/core/delegation"
	"github.com/storacha/go-ucanto/did"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse IPNI metadata.",
	Long:  `Parse a base64 encoded IPNI metadata.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		metadataBytes, err := base64.StdEncoding.DecodeString(args[0])
		if err != nil {
			panic(err)
		}

		md := metadata.MetadataContext.New()
		err = md.UnmarshalBinary(metadataBytes)
		if err != nil {
			panic(err)
		}

		PrintMetadata(md)
	},
}

func PrintMetadata(md ipnimd.Metadata) {
	// the provider may list one or more protocols for this CID
	// in our case, the protocols are just differnt types of content claims
	for _, code := range md.Protocols() {
		protocol := md.Get(code)
		// make sure this is some kind of claim protocol, ignore if not
		hasClaim, ok := protocol.(metadata.HasClaim)
		if !ok {
			fmt.Println("UNKNOWN METADATA")
			continue
		}
		protoName := "unknown"
		if protocol.ID() == metadata.LocationCommitmentID {
			protoName = "location commitment"
		} else if protocol.ID() == metadata.EqualsClaimID {
			protoName = "equivalency claim"
		} else if protocol.ID() == metadata.IndexClaimID {
			protoName = "index claim"
		}
		fmt.Printf("ID: 0x%s (%s)\n", strconv.FormatUint(uint64(protocol.ID()), 16), protoName)
		fmt.Printf("Claim: %s\n", hasClaim.GetClaim())

		if hasClaim.GetClaim().Prefix().MhType == uint64(multicodec.Identity) {
			dmh, err := multihash.Decode(hasClaim.GetClaim().Hash())
			if err != nil {
				panic(err)
			}
			dlg, err := delegation.Extract(dmh.Digest)
			if err != nil {
				panic(err)
			}
			fmt.Printf("  Can: %s\n", dlg.Capabilities()[0].Can())
			fmt.Printf("  With: %s\n", dlg.Capabilities()[0].With())
			if dlg.Capabilities()[0].Can() == assert.LocationAbility {
				fmt.Println("  Nb:")
				nb, err := assert.LocationCaveatsReader.Read(dlg.Capabilities()[0].Nb())
				if err != nil {
					panic(err)
				}
				fmt.Printf("    Content: %s\n", digestutil.Format(nb.Content.Hash()))
				fmt.Println("    Location:")
				for _, l := range nb.Location {
					fmt.Printf("        %s\n", l.String())
				}
				if nb.Range != nil {
					if nb.Range.Length != nil {
						fmt.Printf("    Range: %d-%d (%d bytes)\n", nb.Range.Offset, nb.Range.Offset+*nb.Range.Length, *nb.Range.Length)
					} else {
						fmt.Printf("    Range: %d-\n", nb.Range.Offset)
					}
				}
				if nb.Space != did.Undef {
					fmt.Printf("    Space: %s\n", nb.Space.String())
				}
			} else if dlg.Capabilities()[0].Can() == assert.IndexAbility {
				fmt.Println("  Nb:")
				nb, err := assert.IndexCaveatsReader.Read(dlg.Capabilities()[0].Nb())
				if err != nil {
					panic(err)
				}
				fmt.Printf("    Content: %s\n", nb.Content)
				fmt.Printf("    Index: %s\n", nb.Index)
			}
			if dlg.Expiration() != nil {
				fmt.Printf("  Expiration: %s\n", time.Unix(int64(*dlg.Expiration()), 0).String())
			}
		}

		switch typedProtocol := protocol.(type) {
		case *metadata.EqualsClaimMetadata:
			fmt.Printf("Equals: %s\n", typedProtocol.Equals)
			fmt.Printf("Expiration: %s\n", time.Unix(typedProtocol.Expiration, 0).String())
		case *metadata.IndexClaimMetadata:
			fmt.Printf("Index: %s\n", typedProtocol.Index)
			fmt.Printf("Expiration: %s\n", time.Unix(typedProtocol.Expiration, 0).String())
		case *metadata.LocationCommitmentMetadata:
			if typedProtocol.Shard != nil {
				fmt.Printf("Shard: %s\n", typedProtocol.Shard)
			}
			if typedProtocol.Range != nil {
				if typedProtocol.Range.Length != nil {
					fmt.Printf("Range: %d-%d (%d bytes)\n", typedProtocol.Range.Offset, typedProtocol.Range.Offset+*typedProtocol.Range.Length, *typedProtocol.Range.Length)
				} else {
					fmt.Printf("Range: %d-\n", typedProtocol.Range.Offset)
				}
			}
			fmt.Printf("Expiration: %s\n", time.Unix(typedProtocol.Expiration, 0).String())
		}
	}
}

func init() {
	parseCmd.Flags().Bool("json", false, "Output DAG JSON")
	Cmd.AddCommand(parseCmd)
}
