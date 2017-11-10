package types

type Browser_WindowID int
type Browser_WindowState string
type Browser_Bounds struct {
	Left        *int                 `json:"left,omitempty"`
	Top         *int                 `json:"top,omitempty"`
	Width       *int                 `json:"width,omitempty"`
	Height      *int                 `json:"height,omitempty"`
	WindowState *Browser_WindowState `json:"windowState,omitempty"`
}
