package flatfs

import (
	"path"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
	"github.com/storacha/piri/pkg/store/objectstore/flatfs"
)

var pathCmd = &cobra.Command{
	Use:   "path <blob-cid>",
	Short: "Obtain the FlatFS datastore path for a blob CID (Piri edition).",
	Long:  "Given a blob CID, convert it to a FlatFS datastore path (Piri edition).",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")
		cid, err := cid.Parse(args[0])
		if err != nil {
			panic(err)
		}
		b32, _ := multibase.Encode(multibase.Base32, cid.Hash())
		dir := flatfs.NextToLast(2).Func()(b32[1:])
		cmd.Println(path.Join("/", dir, b32[1:]+".data"))
	},
}

func init() {
	Cmd.AddCommand(pathCmd)
}
