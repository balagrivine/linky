package utils

import (
	"crypto/sha256"
	b62 "encoding/base62"
)

// Helper function to hash URL id
func HashURLId(url_id int) (string) {
	
	// Convert id to string to enable hashing
	id := string(url_id)

	hash := sha256.New()

	hash.Write([]byte(id))
	hashed_string := hash.Sum(nil)

	return hashed_string
}

// Helper function to encode the hashed URL id
func EncodeURL(hashed_url string) (string, error) {

	encoded_url, err := b62.URLEncoding.EncodeToString([]byte(data))
	if err != nil {
		return "", err
	}

	return encoded_url, nil
}

// Helper function to decode the encoded url
func DecodeURL(encoded_url string) (string, error) {

	decoded_string, err := b62.URLDecoding.DecodeString(encoded_url)
	if err != nil {
		return "", err
	}

	return string(decoded_url), nil
}
