package types

type Log_LogEntry struct {
	Source			string			`json:"source"`// Log entry source.
	Level			string			`json:"level"`// Log entry severity.
	Text			string			`json:"text"`// Logged text.
	Timestamp		Runtime_Timestamp	`json:"timestamp"`// Timestamp when this entry was added.
	Url			*string			`json:"url,omitempty"`// URL of the resource if known.
	LineNumber		*int			`json:"lineNumber,omitempty"`// Line number in the resource.
	StackTrace		*Runtime_StackTrace	`json:"stackTrace,omitempty"`// JavaScript stack trace.
	NetworkRequestId	*Network_RequestId	`json:"networkRequestId,omitempty"`// Identifier of the network request associated with this entry.
	WorkerId		*string			`json:"workerId,omitempty"`// Identifier of the worker associated with this entry.
	Args			[]Runtime_RemoteObject	`json:"args,omitempty"`// Call arguments.
}
type Log_ViolationSetting struct {
	Name		string	`json:"name"`// Violation type.
	Threshold	float32	`json:"threshold"`// Time threshold to trigger upon.
}
