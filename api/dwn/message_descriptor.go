package dwn

// Message Descriptors are JSON objects that contains the parameters, signatory proof, and
// other details about the message and any data associated with it.
// https://identity.foundation/decentralized-web-node/spec/#message-descriptors
type MessageDescriptor struct {
	Method        string `json:"method,omitempty"`
	ObjectId      string `json:"objectId,omitempty"`
	Schema        string `json:"schema,omitempty"`
	DataFormat    string `json:"dataFormat,omitempty"`
	DateCreated   string `json:"dateCreated,omitempty"`
	DatePublished string `json:"datePublished,omitempty"`
	DateSort      string `json:"dateSort,omitempty"`
	Root          string `json:"root,omitempty"`
	Parent        string `json:"parent,omitempty"`
	Cid           string `json:"cid,omitempty"`
}
