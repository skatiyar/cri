package types

type Browser_WindowID int
type Browser_WindowState string
type Browser_Bounds struct {
	Left		*int			`json:"left,omitempty"`// The offset from the left edge of the screen to the window in pixels.
	Top		*int			`json:"top,omitempty"`// The offset from the top edge of the screen to the window in pixels.
	Width		*int			`json:"width,omitempty"`// The window width in pixels.
	Height		*int			`json:"height,omitempty"`// The window height in pixels.
	WindowState	*Browser_WindowState	`json:"windowState,omitempty"`// The window state. Default to normal.
}
