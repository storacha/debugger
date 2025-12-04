package flatfs

import (
	"path"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
	"github.com/storacha/go-libstoracha/digestutil"
	"github.com/storacha/piri/pkg/store/objectstore/flatfs"
)

var pathCmd = &cobra.Command{
	Use:   "path <blob-cid-or-multihash>",
	Short: "Obtain the FlatFS datastore path for a blob CID or multihash (Piri edition).",
	Long:  "Given a blob CID or multibase encoded multihash, convert it to a FlatFS datastore path (Piri edition).",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")
		digest, err := digestutil.Parse(args[0])
		if err != nil {
			// try CID
			cid, err := cid.Parse(args[0])
			cobra.CheckErr(err)
			digest = cid.Hash()
		}
		b32, _ := multibase.Encode(multibase.Base32, digest)
		dir := flatfs.NextToLast(2).Func()(b32[1:])
		cmd.Println(path.Join("/", dir, b32[1:]+".data"))
	},
}

func init() {
	Cmd.AddCommand(pathCmd)
}
