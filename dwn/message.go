package dwn

// All Decentralized Web Node messaging is transacted via Messages JSON objects.
// These objects contain message execution parameters, authorization material,
// authorization signatures, and signing/encryption information.
// For various purposes Messages rely on IPFS CIDs and DAG APIs.
// https://identity.foundation/decentralized-web-node/spec/#messages
type Message struct {
	// Message objects MAY contain a data property, and if present its value MUST
	// be a JSON value of the Message’s data.
	Data string `json:"data,omitempty"`

	// Message objects MUST contain a descriptor property, and its value MUST be an
	// object, as defined by the Message Descriptors section of this specification.
	Descriptor_ *MessageDescriptor `json:"descriptor,omitempty"`

	// Message objects MAY contain an attestation property, and if present its value
	// MUST be an object, as defined by the Signed Data section of this specification.
	Attestation *Attestation `json:"attestation,omitempty"`
}
