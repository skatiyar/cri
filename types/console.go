package types

type Console_ConsoleMessage struct {
	Source	string	`json:"source"`// Message source.
	Level	string	`json:"level"`// Message severity.
	Text	string	`json:"text"`// Message text.
	Url	*string	`json:"url,omitempty"`// URL of the message origin.
	Line	*int	`json:"line,omitempty"`// Line number in the resource that generated this message (1-based).
	Column	*int	`json:"column,omitempty"`// Column number in the resource that generated this message (1-based).
}
