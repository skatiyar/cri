package types

type Log_LogEntry struct {
	// Log entry source.
	Source string `json:"source"`
	// Log entry severity.
	Level string `json:"level"`
	// Logged text.
	Text string `json:"text"`
	// Timestamp when this entry was added.
	Timestamp Runtime_Timestamp `json:"timestamp"`
	// URL of the resource if known.
	Url *string `json:"url,omitempty"`
	// Line number in the resource.
	LineNumber *int `json:"lineNumber,omitempty"`
	// JavaScript stack trace.
	StackTrace *Runtime_StackTrace `json:"stackTrace,omitempty"`
	// Identifier of the network request associated with this entry.
	NetworkRequestId *Network_RequestId `json:"networkRequestId,omitempty"`
	// Identifier of the worker associated with this entry.
	WorkerId *string `json:"workerId,omitempty"`
	// Call arguments.
	Args []Runtime_RemoteObject `json:"args,omitempty"`
}
type Log_ViolationSetting struct {
	// Violation type.
	Name string `json:"name"`
	// Time threshold to trigger upon.
	Threshold float32 `json:"threshold"`
}
