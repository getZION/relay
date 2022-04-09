package handler

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/base58"
	"github.com/getzion/relay/api/dwn"
	"github.com/getzion/relay/api/errors"
	. "github.com/getzion/relay/utils"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/sirupsen/logrus"
)

type RequestContext struct {
	// SchemaManager *schema.SchemaManager
	Message *dwn.Message
}

// Get the public key from the attestation DID.
func (c *RequestContext) GetPublicKey() (*ecdsa.PublicKey, *errors.MessageLevelError) {
	// Validate the attestation object
	if c.Message.Attestation == nil {
		return nil, errors.NewMessageLevelError(400, "attestation cannot be null or empty", nil)
	} else if c.Message.Attestation.Protected == nil {
		return nil, errors.NewMessageLevelError(400, "attestation protected cannot be null", nil)
	} else if c.Message.Attestation.Protected.Alg != "ES256K" {
		return nil, errors.NewMessageLevelError(400, "Unsupported signing algorithm", nil)
	} else if strings.Trim(c.Message.Attestation.Protected.Kid, " ") == "" {
		return nil, errors.NewMessageLevelError(400, "Unsupported DID method", nil)
	} else if strings.HasPrefix(c.Message.Attestation.Protected.Kid, "did:key:") == false {
		return nil, errors.NewMessageLevelError(400, "Unsupported DID method", nil)
	}

	// Convert the multibase fingerprint to bytes.
	// This reverses the client-side method `getMultibaseFingerprintFromPublicKeyBytes`.
	fingerprintWithoutPrefix := strings.TrimPrefix(c.Message.Attestation.Protected.Kid, "did:key:z")
	didBytes := base58.Decode(fingerprintWithoutPrefix)
	pubKeyBytes := make([]byte, 33)
	pubKeyBytes = didBytes[2:]

	// Parse the pubkey bytes to verify that it corresponds to a valid public key on the secp256k1 curve.
	pubKey, err := btcec.ParsePubKey(pubKeyBytes)
	if err != nil {
		return nil, errors.NewMessageLevelError(400, "Invalid pubkey", nil)
	} else {
		Log.Info().
			Bool("Compressed", btcec.IsCompressedPubKey(pubKeyBytes)).
			Msg("Received valid pubkey")
	}

	// Convert the secp256k1 key to an ECDSA key.
	ecdsaKey := pubKey.ToECDSA()

	return ecdsaKey, nil
}

func (c *RequestContext) RecreateStringToSign() (string, *errors.MessageLevelError) {
	payload := c.Message.Attestation.Payload
	stringifiedProtected, err := json.Marshal(&c.Message.Attestation.Protected)
	if err != nil {
		logrus.Errorf("stringifiedProtected parse failed: %v", err)
		return "", errors.NewMessageLevelError(400, "stringifiedProtected parse failed", err)
	}
	base64Protected := base64.RawURLEncoding.EncodeToString([]byte(string(stringifiedProtected)))
	signedString := fmt.Sprintf("%s.%s.", base64Protected, payload)
	return signedString, nil
}

// Validate that this signature was signed by the given key.
func (c *RequestContext) VerifyRequest(publicKey *ecdsa.PublicKey) (bool, *errors.MessageLevelError) {

	sigTest := c.Message.Attestation.Signature

	verified, err3 := jws.Verify([]byte(sigTest), jwa.ES256K, publicKey)
	fmt.Printf("VERIFIED PAYLOAD: %s \n   ERR: %s", verified, err3)
	return false, nil
	// payload := "bagaaieraxcfupl7um4tcsbbckzynk2gnd7y6sjwhubhtzlqi6tpkiqpsnhaa"

	// testInput := []byte(`{
	// 	"protected": "eyJhbGciOiJFUzI1NksiLCJraWQiOiJkaWQ6a2V5OnpRM3NoWVpoUVo4dnpaZDM2Um1LRGpGa29ydExVOXNNZXNhUmpTNjJQVDdjTVpzWFAifQ",
	// 	"payload": "ImJhZ2FhaWVyYXhjZnVwbDd1bTR0Y3NiYmNrenluazJnbmQ3eTZzandodWJodHpscWk2dHBraXFwc25oYWEi",
	// 	"signature": "IjgxNzliYTVjMGEwNThlMTg1ZjM5YmYxMGZkYjFjZjBhNTZmMDc5MjRmNmM4OGRmMzc2ZmY0ZTU2MTk0YjMwMTA2NWRjNTRiZjZhNzNlZjVlMzYzZWFkZmE4NmVkOTkzYTRlMGJmOWI3MjFlM2FhNzU0NTA0NmQ4OTIyMWZjYzJiIg"
	//  }`)

	// testo, _ := jws.Parse(testInput)
	// siggy := testo.LookupSignature(c.Message.Attestation.Protected.Kid)[0]

	// hdrs := jws.NewHeaders()

	// hdrs.Set("alg", c.Message.Attestation.Protected.Alg)
	// hdrs.Set("kid", c.Message.Attestation.Protected.Kid)

	// siggy.SetProtectedHeaders(hdrs)
	// testo.ClearSignatures()
	// testo.AppendSignature(siggy)
	// // testo.
	// fmt.Printf("testo payload: %s \n\n", testo.Payload())
	// fmt.Printf("testo signature: %s \n\n", siggy.Signature())
	// fmt.Printf("testo payload: %s \n\n", testo.)

	// buf, err2 := json.Marshal(testo)
	// buf, err2 := json.MarshalIndent(testo, "", "  ")
	// if err2 != nil {
	// 	fmt.Printf("%s\n", err2)
	// 	return false, nil
	// }

	// fmt.Printf("buf: %s \n\n", buf)
	// fmt.Printf("[]byte(buf): %s \n\n", []byte(buf))

	// testo.UnmarshalJSON()

	// payload, err := jws.Verify(), jwa.ES256K, publicKey)
	// payload, err := jws.Verify([]byte(buf), jwa.ES256K, publicKey)
	// fmt.Printf("DUMMY VERIFIED PAYLOAD: %s \n   ERR: %s", payload, err)

	// return false, nil

	// const keysrc = `{"kty":"oct",
	// 	"k":"AyM1SysPpbyDfgZld3umj1qzKObwVMkoqQ-EstJQLr_T-1qS0gZH75aKtMN3Yj0iPS4hcgUuTwjAzZr1Z9CAow"
	// 	}`

	// key, _ := jwk.ParseKey([]byte(keysrc))
	// // inputa := []byte(`eyJhbGciOiJIUzI1NiIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19..A5dxf2s96_n5FLueVuW1Z_vh161FwXZC4YLPff6dmDY`)
	// inputa := []byte(`{
	// 	"protected": "eyJhbGciOiJIUzI1NiIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19",
	// 	"payload": "$.02",
	// 	"signature": "A5dxf2s96_n5FLueVuW1Z_vh161FwXZC4YLPff6dmDY"
	//  }`)

	// payload, err := jws.Verify(inputa, jwa.HS256, key)
	// fmt.Printf("DUMMY VERIFIED PAYLOAD: %s \n   ERR: %s", payload, err)

	// return false, nil

	// var public1, protected1 jws.Headers
	// var protected1 jws.Headers

	fmt.Printf("c.Message.Attestation.Payload: %s \n", c.Message.Attestation.Payload)

	decodedPayload := c.Message.Attestation.Payload
	// decodedPayload, _ := base64.StdEncoding.DecodeString(c.Message.Attestation.Payload)
	sigBytes := []byte(c.Message.Attestation.Signature)
	// First build the proper JWS format from context values.

	fmt.Printf("decodedPayload: %s \n", decodedPayload)

	fmt.Printf("c.Message.Attestation.Protected.Alg: %s \n", c.Message.Attestation.Protected.Alg)
	fmt.Printf("c.Message.Attestation.Protected.Kid: %s \n", c.Message.Attestation.Protected.Kid)

	hdrs := jws.NewHeaders()

	hdrs.Set("alg", c.Message.Attestation.Protected.Alg)
	hdrs.Set("kid", c.Message.Attestation.Protected.Kid)

	fmt.Printf("\n hdrs: %s \n", hdrs)
	fmt.Printf("\n Algorithm: %s \n\n", hdrs.Algorithm())
	fmt.Printf("\n KeyID: %s \n\n", hdrs.KeyID())

	s := jws.NewSignature().
		SetSignature(sigBytes).
		SetProtectedHeaders(hdrs)
		//.
		// SetPublicHeaders(public1)

	fmt.Printf("\n signature: %s \n\n", s)

	m := jws.NewMessage().
		SetPayload([]byte(decodedPayload)).
		AppendSignature(s)

	fmt.Printf("\n m.Payload(): %s \n\n", m.Payload())

	// return false, nil

	buf, err2 := json.Marshal(m) // buf
	// buf, err2 := json.MarshalIndent(m, "", "  ") // buf
	if err2 != nil {
		fmt.Printf("%s\n", err2)
		return false, nil
	}

	fmt.Printf("buf: %v \n", buf)
	fmt.Printf("jwa.ES256K: %v \n", jwa.ES256K)
	fmt.Printf("publicKey: %v \n", publicKey)

	// what := jws.Parse()
	// []byte("{ \"asdfasdf\": [] }")

	// verified, err3 := jws.Verify(buf, jwa.ES256K, publicKey)
	// fmt.Printf("\n verified: %v, error:%v \n", verified, err3)

	// return false, nil

	// fmt.Printf("m: %v", m)

	// verified, err3 := jws.Verify(buf, jwa.ES256K, publicKey)

	return false, nil
}

func (c *RequestContext) VerifyRequestOld(signedString string, publicKey *ecdsa.PublicKey) (bool, *errors.MessageLevelError) {
	verified, err := jws.Verify([]byte(signedString), jwa.ES256K, publicKey)
	fmt.Printf("verified: %v, error:%v", verified, err)
	return false, nil
}
