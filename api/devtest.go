package api

import (
	"encoding/base64"
	"encoding/json"

	. "github.com/getzion/relay/utils"
)

type ParsedData struct {
	Model    string `json:"model,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Did      string
}

func Testo() {
	// Start with { model: 'Zion.User.V1' }, base64uri-encoded
	// data := "eyJtb2RlbCI6Ilppb24uVXNlci5WMSJ9"
	// data := "eyJtb2RlbCI6Ilppb24uVXNlci5WMSIsIm5hbWUiOiJUZXN0ZXIgTWFuIn0"
	data := "eyJtb2RlbCI6Ilppb24uVXNlci5WMSIsIm5hbWUiOiJUZXN0ZXIgTWFuIiwidXNlcm5hbWUiOiJidWNrb3Rlc3RvIiwiZGlkIjoiZGlkOmtleTp6UTNzaGZSMXFqdkJoTVFmelpQNXJKdzZMWm53VnVWZkIyQTh0dHNocGJhV3dFVjFHIn0"
	decodedData, _ := base64.StdEncoding.DecodeString(data)

	var parsedData ParsedData
	if err := json.Unmarshal(decodedData, &parsedData); err != nil {
		if err.Error() == "unexpected end of JSON input" {
			json.Unmarshal([]byte(string(decodedData)+"\"}"), &parsedData)
			Log.Info().Str("wat", string(decodedData)+"\"}").Msg("Retrying with closing brace")
		} else {
			Log.Err(err).Str("error msg?", err.Error()).Msg("Error unmarshaling decodedData.")
			panic(err)
		}
	}
	// json.Unmarshal(decodedData, &parsedData)

	Log.Info().
		Str("Data", data).
		Str("decodedData", string(decodedData)).
		Str("parsed Model", parsedData.Model).
		Str("parsed Name", parsedData.Name).
		Str("parsed Username", parsedData.Username).
		Str("parsed Did", parsedData.Did).
		Msg("Parse")

}
