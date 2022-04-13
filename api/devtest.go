package api

import (
	"encoding/base64"
	"encoding/json"

	. "github.com/getzion/relay/utils"
)

type ParsedData struct {
	Model string `json:"model,omitempty"`
	// Name     string `json:"name,omitempty"`
	// Username string `json:"username,omitempty"`
	// Did      string
	// OwnerDid string
}

func Testo() {
	// Start with { model: 'Zion.User.V1' }, base64uri-encoded
	// data := "eyJtb2RlbCI6Ilppb24uVXNlci5WMSJ9"
	// data := "eyJtb2RlbCI6Ilppb24uVXNlci5WMSIsIm5hbWUiOiJUZXN0ZXIgTWFuIn0"
	// data := "eyJtb2RlbCI6Ilppb24uVXNlci5WMSIsIm5hbWUiOiJUZXN0ZXIgTWFuIiwidXNlcm5hbWUiOiJidWNrb3Rlc3RvIiwiZGlkIjoiZGlkOmtleTp6UTNzaGZSMXFqdkJoTVFmelpQNXJKdzZMWm53VnVWZkIyQTh0dHNocGJhV3dFVjFHIn0"
	// data := "eyJuYW1lIjoiVGVzdCBDb21tdW5pdHkiLCJkZXNjcmlwdGlvbiI6IkF3ZXNvbWUgdGVzdCBjb21tdW5pdHkiLCJwcmljZVBlck1lc3NhZ2UiOjUsInByaWNlVG9Kb2luIjo1LCJvd25lckRpZCI6ImRpZDprZXk6elEzc2hmUjFxanZCaE1RZnpaUDVySnc2TFpud1Z1VmZCMkE4dHRzaHBiYVd3RVYxRyIsInppZCI6ImUxYWFhMjY2LTQ0OTctNGVhYi1iNGE2LWNkZWU5Njk4NjZiNSIsIm1vZGVsIjoiWmlvbi5Db21tdW5pdHkuVjEifQ"
	data := "eyJuYW1lIjoiVGVzdCBDb21tdW5pdHkiLCJkZXNjcmlwdGlvbiI6IkF3ZXNvbWUgdGVzdCBjb21tdW5pdHkiLCJwcmljZVBlck1lc3NhZ2UiOjUsInByaWNlVG9Kb2luIjo1LCJvd25lckRpZCI6ImRpZDprZXk6elEzc2hmUjFxanZCaE1RZnpaUDVySnc2TFpud1Z1VmZCMkE4dHRzaHBiYVd3RVYxRyIsInppZCI6ImU1MjE2MjgwLTkzYjctNDYwMS1hZDFjLWEzMGVjNjE3YTU4YSIsIm1vZGVsIjoiWmlvbi5Db21tdW5pdHkuVjEifQ"
	// decodedData, _ := base64.StdEncoding.DecodeString(data)
	// Log.Info().Str("WAT", string(decodedData)).Msg("BEFORE")
	// return
	// decodedData, _ := base64.StdEncoding.DecodeString(data)

	b64String, _ := base64.StdEncoding.DecodeString(data)
	var parsedData ParsedData
	json.Unmarshal(b64String, &parsedData)
	Log.Info().Bytes("wat", []byte(b64String)).Msg("WHAT")

	Log.Info().
		Str("Data", data).
		// Str("decodedData", string(decodedData)).
		Str("parsed Model", parsedData.Model).
		// Str("parsed Name", parsedData.Name).
		// Str("parsed Username", parsedData.Username).
		// Str("parsed Did", parsedData.Did).
		// Str("parsed OwnerDid", parsedData.OwnerDid).
		Msg("Parse")

	// json.Unmarshal(decodedData, &parsedData)
	// []byte(b64String)
	// rawIn := json.RawMessage(b64String)
	// bodyBytes, err := rawIn.MarshalJSON()

	// if err != nil {
	// panic(err)
	// }

	return

	if err := json.Unmarshal([]byte(data), &parsedData); err != nil {
		panic(err)
		// if err.Error() == "unexpected end of JSON input" {
		// 	Log.Info().Str("wat", string(decodedData)).Msg("BEFORE")
		// 	decodedData = []byte(string(decodedData) + "\"}")
		// 	json.Unmarshal([]byte(string(decodedData)), &parsedData)
		// 	Log.Info().Str("wat", string(decodedData)).Msg("Retrying with closing brace")
		// } else {
		// 	Log.Err(err).Str("error msg?", err.Error()).Msg("Error unmarshaling decodedData.")
		// 	panic(err)
		// }
	}
	//

}
