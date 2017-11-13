/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Screen orientation.
type Emulation_ScreenOrientation struct {
	// Orientation type.
	Type string `json:"type"`
	// Orientation angle.
	Angle int `json:"angle"`
}

// advance: If the scheduler runs out of immediate work, the virtual time base may fast forward to allow the next delayed task (if any) to run; pause: The virtual time base may not advance; pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending resource fetches.
type Emulation_VirtualTimePolicy string
