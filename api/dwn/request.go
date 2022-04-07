package dwn

// Request Objects are JSON object envelopes used to pass messages to Decentralized Web Nodes.
// https://identity.foundation/decentralized-web-node/spec/#request-objects
type Request struct {
	// The Request Object MUST include a requestId property, and its value MUST be
	// an [RFC4122] UUID Version 4 string to identify the request.
	RequestId string `json:"requestId,omitempty"`

	// The Request Object MUST include a target property, and its value MUST be the
	// Decentralized Identifier base URI of the DID-relative URL.
	Target string `json:"target,omitempty"`

	// The Request Object MUST include a messages property, and its value MUST be an
	// array composed of Message objects.
	Messages []*Message `json:"messages,omitempty"`
}
