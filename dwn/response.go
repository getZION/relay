package dwn

// Responses from Interface method invocations are JSON objects that MUST be constructed as follows:
// https://identity.foundation/decentralized-web-node/spec/#response-objects
type Response struct {
	// The object MUST include an requestId property, and its value MUST be the
	// [RFC4122] UUID Version 4 string from the requestId property of the Request Object it is in response to.
	RequestId string `json:"requestId,omitempty"`

	// The object MUST have a status property, and its value MUST be an object composed of the following properties:
	Status *Status `json:"status,omitempty"`

	// The object MAY have a replies property, and if present its value MUST be an
	// array of Message Result Objects, which are constructed as follows:
	Replies []*Reply `json:"replies,omitempty"`
}

type Status struct {
	// The status object MUST have a code property, and its value MUST be an integer
	// set to the HTTP Status Code appropriate for the status of the response.
	Code int `json:"code,omitempty"`

	// The status object MAY have a message property, and if present its value MUST
	// be a string that describes a terse summary of the status. It is RECOMMENDED
	// that the implementer set the message text to the standard title of the HTTP
	// Status Code, when a title/message has already been defined for that code.
	Message string `json:"message,omitempty"`
}

type Reply struct {
	// The object MUST have a messageId property, and its value MUST be the stringified
	// Version 1 CID of the associated message in the Request Object from which it was
	// received.
	MessageId string `json:"messageId,omitempty"`

	// The object MUST have a status property, and its value MUST be an object composed
	// of the following properties:
	Status *Status `json:"status,omitempty"`

	// The object MAY have a entries property if the message request was successful.
	// If present, its value MUST be the resulting message entries returned from the
	// invocation of the corresponding message.
	Entries []string `json:"entries,omitempty"`
}
