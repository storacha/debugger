package ucanfmt

import (
	"fmt"

	"github.com/storacha/go-ucanto/core/message"
)

func PrintMessage(m message.AgentMessage) {
	fmt.Printf("%s\n", m.Root().Link())

	invocations := m.Invocations()
	if len(invocations) > 0 {
		fmt.Println("  Invocations:")
		for _, invLink := range invocations {
			inv, ok, err := m.Invocation(invLink)
			if err != nil {
				fmt.Printf("    Error getting invocation %s: %v\n", invLink, err)
				continue
			}
			if !ok {
				fmt.Printf("    Invocation not found: %s\n", invLink)
				continue
			}
			doPrintDelegation(inv, 1)
		}
	}

	receipts := m.Receipts()
	if len(receipts) > 0 {
		fmt.Println("  Receipts:")
		for _, rcptLink := range receipts {
			rcpt, ok, err := m.Receipt(rcptLink)
			if err != nil {
				fmt.Printf("    Error getting receipt %s: %v\n", rcptLink, err)
				continue
			}
			if !ok {
				fmt.Printf("    Receipt not found: %s\n", rcptLink)
				continue
			}
			doPrintReceipt(rcpt, 1)
		}
	}
}
