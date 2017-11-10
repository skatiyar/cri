package types

type Console_ConsoleMessage struct {
	Source string  `json:"source"`
	Level  string  `json:"level"`
	Text   string  `json:"text"`
	Url    *string `json:"url,omitempty"`
	Line   *int    `json:"line,omitempty"`
	Column *int    `json:"column,omitempty"`
}
