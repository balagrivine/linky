package utils

import (
	"crypto/sha256"
	b64 "encoding/base64"
)

// Helper function to hash URL id
func HashURLId(url_id int32) (string) {
	
	// Convert id to string to enable hashing
	id := string(url_id)

	hash := sha256.New()

	hash.Write([]byte(id))
	hashed_string := hash.Sum(nil)

	return string(hashed_string)
}

// Helper function to encode the hashed URL id
func EncodeURL(hashed_url string) (string, error) {

	encoded_url := b64.URLEncoding.EncodeToString([]byte(hashed_url))

	return encoded_url, nil
}

// Helper function to decode the encoded url
func DecodeURL(encoded_url string) (string, error) {

	decoded_url, err := b64.URLEncoding.DecodeString(encoded_url)
	if err != nil {
		return "", nil
	}

	return string(decoded_url), nil
}
