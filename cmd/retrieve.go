package cmd

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/go-ucanto/client/retrieval"
	"github.com/storacha/go-ucanto/core/invocation"
	"github.com/storacha/go-ucanto/transport/headercar/message"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve <url> <auth>",
	Short: "Retrieve data from the network.",
	Long: `Attempt to retrieve data from the passed URL using the provided
authorization - an X-Agent-Message header.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		url, err := url.Parse(args[0])
		if err != nil {
			panic(fmt.Errorf("parsing retrieval URL: %w", err))
		}

		var inv invocation.Invocation

		authInput := args[1]

		msg, err := message.DecodeHeader(authInput)
		if err != nil {
			panic(fmt.Errorf("decoding X-Agent-Message header: %w", err))
		}
		cmd.Printf("successfully decoded X-Agent-Message header (%d bytes)\n", len(authInput))
		if len(msg.Invocations()) != 1 {
			panic(fmt.Errorf("unexpected number of invocations in message: %d", len(msg.Invocations())))
		}
		inv, _, err = msg.Invocation(msg.Invocations()[0])
		if err != nil {
			panic(fmt.Errorf("extracting invocation %s: %w", msg.Invocations()[0], err))
		}
		if len(inv.Capabilities()) != 1 {
			panic(fmt.Errorf("unexpected number of capabilities in invocation: %d", len(inv.Capabilities())))
		}

		cmd.Println("blocks:")
		for b, err := range inv.Blocks() {
			if err != nil {
				panic(fmt.Errorf("iterating invocation blocks: %w", err))
			}
			cmd.Printf("%s\n", b.Link())
			s, err := ipldfmt.FormatDagCBOR(b.Bytes())
			if err != nil {
				panic(fmt.Errorf("formatting block %s: %w", b.Link(), err))
			}
			cmd.Println(s)
			cmd.Println("")
		}

		cmd.Printf("sending %s invocation to %s\n", inv.Capabilities()[0].Can(), inv.Audience().DID())
		cmd.Printf("GET %s\n", url.String())

		conn, err := retrieval.NewConnection(inv.Audience(), url)
		if err != nil {
			panic(fmt.Errorf("creating connection to %s: %w", inv.Audience().DID(), err))
		}

		xres, hres, err := retrieval.Execute(context.Background(), inv, conn)
		if hres != nil {
			cmd.Printf("Response status: %d\n", hres.Status())
			cmd.Printf("Response headers:\n")
			for k, v := range hres.Headers() {
				cmd.Printf("\t%s: %s\n", k, strings.Join(v, ","))
			}
		}
		if err != nil {
			if hres != nil {
				cmd.Printf("Response body:\n")
				io.Copy(cmd.OutOrStderr(), hres.Body())
				cmd.Println()
				return
			}
			panic(fmt.Errorf("executing invocation: %w", err))
		}

		cmd.Println("Response blocks:")
		for b, err := range xres.Blocks() {
			if err != nil {
				panic(fmt.Errorf("iterating execution response blocks: %w", err))
			}
			cmd.Printf("%s\n", b.Link())
			s, err := ipldfmt.FormatDagCBOR(b.Bytes())
			if err != nil {
				panic(fmt.Errorf("formatting block %s: %w", b.Link(), err))
			}
			cmd.Println(s)
			cmd.Println("")
		}

		outfile, _ := cmd.Flags().GetString("output")
		if outfile != "" {
			f, err := os.Create(outfile)
			if err != nil {
				panic(fmt.Errorf("creating output file: %w", err))
			}
			io.Copy(f, hres.Body())
			cmd.Printf("Response body written to: %s\n", outfile)
		} else {
			n, _ := io.Copy(io.Discard, hres.Body())
			cmd.Printf("Response body discarded (use --output or -o to save to file), %d bytes transferred\n", n)
		}
	},
}

func init() {
	retrieveCmd.Flags().StringP("output", "o", "", "Save output to the specified file")
	rootCmd.AddCommand(retrieveCmd)
}
