/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// ServiceWorker registration.
type ServiceWorker_ServiceWorkerRegistration struct {
	RegistrationId string `json:"registrationId"`
	ScopeURL       string `json:"scopeURL"`
	IsDeleted      bool   `json:"isDeleted"`
}

type ServiceWorker_ServiceWorkerVersionRunningStatus string

type ServiceWorker_ServiceWorkerVersionStatus string

// ServiceWorker version.
type ServiceWorker_ServiceWorkerVersion struct {
	VersionId      string                                          `json:"versionId"`
	RegistrationId string                                          `json:"registrationId"`
	ScriptURL      string                                          `json:"scriptURL"`
	RunningStatus  ServiceWorker_ServiceWorkerVersionRunningStatus `json:"runningStatus"`
	Status         ServiceWorker_ServiceWorkerVersionStatus        `json:"status"`
	// The Last-Modified header value of the main script.
	ScriptLastModified *float32 `json:"scriptLastModified,omitempty"`
	// The time at which the response headers of the main script were received from the server. For cached script it is the last time the cache entry was validated.
	ScriptResponseTime *float32          `json:"scriptResponseTime,omitempty"`
	ControlledClients  []Target_TargetID `json:"controlledClients,omitempty"`
	TargetId           *Target_TargetID  `json:"targetId,omitempty"`
}

// ServiceWorker error message.
type ServiceWorker_ServiceWorkerErrorMessage struct {
	ErrorMessage   string `json:"errorMessage"`
	RegistrationId string `json:"registrationId"`
	VersionId      string `json:"versionId"`
	SourceURL      string `json:"sourceURL"`
	LineNumber     int    `json:"lineNumber"`
	ColumnNumber   int    `json:"columnNumber"`
}
