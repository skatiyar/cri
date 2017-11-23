/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package cri

// BrowserProtocolVersion returns current major and minor version of browser protocol.
func BrowserProtocolVersion() (major, minor string) {
	return {{printf "%q" .BMajor}}, {{printf "%q" .BMinor}}
}

// JsProtocolVersion returns current major and minor version of js protocol.
func JsProtocolVersion() (major, minor string) {
	return {{printf "%q" .JMajor}}, {{printf "%q" .JMinor}}
}
