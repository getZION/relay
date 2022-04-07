package dwn

// If the object is to be attested by a signer (e.g the Node owner via signature with their DID key),
// the object MUST contain the following additional properties to produce a [RFC7515] Flattened
// JSON Web Signature (JWS) object:
type Attestation struct {
	// The object MUST include a protected property...
	Protected *AttestationProtected `json:"protected,omitempty"`

	// The object MUST include a payload property, and its value MUST be the stringified Version 1 CID
	// of the DAG CBOR encoded descriptor object, whose composition is defined in the Message Descriptor
	// section of this specification.
	Payload string `json:"payload,omitempty"`

	// The object MUST include a signature property, and its value MUST be a signature string produced
	// by signing the protected and payload values, in accordance with the [RFC7515] JSON Web Signature
	// specification.
	Signature string `json:"signature,omitempty"`
}

// ...and its value MUST be an object composed of the following values:
type AttestationProtected struct {
	// The object MUST include an alg property, and its value MUST be the string representing the
	// algorithm used to verify the signature (as defined by the [RFC7515] JSON Web Signature specification).
	Alg string `json:"alg,omitempty"`

	// The object MUST include a kid property, and its value MUST be a DID URL string identifying
	// the key to be used in verifying the signature.
	Kid string `json:"kid,omitempty"`
}
