package crypto

import "encoding/base64"

func Base64Encode(str []byte) string {
	return base64.RawURLEncoding.EncodeToString(str)
}

func Base64Decode(str string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(str)
}
