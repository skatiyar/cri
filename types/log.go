package types

type Log_LogEntry struct {
	Source           string                 `json:"source"`
	Level            string                 `json:"level"`
	Text             string                 `json:"text"`
	Timestamp        Runtime_Timestamp      `json:"timestamp"`
	Url              *string                `json:"url,omitempty"`
	LineNumber       *int                   `json:"lineNumber,omitempty"`
	StackTrace       *Runtime_StackTrace    `json:"stackTrace,omitempty"`
	NetworkRequestId *Network_RequestId     `json:"networkRequestId,omitempty"`
	WorkerId         *string                `json:"workerId,omitempty"`
	Args             []Runtime_RemoteObject `json:"args,omitempty"`
}
type Log_ViolationSetting struct {
	Name      string  `json:"name"`
	Threshold float32 `json:"threshold"`
}
