/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Console message.
type Console_ConsoleMessage struct {
	// Message source.
	Source string `json:"source"`
	// Message severity.
	Level string `json:"level"`
	// Message text.
	Text string `json:"text"`
	// URL of the message origin.
	Url *string `json:"url,omitempty"`
	// Line number in the resource that generated this message (1-based).
	Line *int `json:"line,omitempty"`
	// Column number in the resource that generated this message (1-based).
	Column *int `json:"column,omitempty"`
}
