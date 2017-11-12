/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

type Target_TargetID string

// Unique identifier of attached debugging session.
type Target_SessionID string

type Target_BrowserContextID string

type Target_TargetInfo struct {
	TargetId Target_TargetID `json:"targetId"`
	Type     string          `json:"type"`
	Title    string          `json:"title"`
	Url      string          `json:"url"`
	// Whether the target has an attached client.
	Attached bool `json:"attached"`
	// Opener target Id
	OpenerId *Target_TargetID `json:"openerId,omitempty"`
}

type Target_RemoteLocation struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
