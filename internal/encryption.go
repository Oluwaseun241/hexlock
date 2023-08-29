package internal

import (
	"crypto/rand"
)

func GenerateKey() []byte {
  key := make([]byte, 32)
  _, err := rand.Read(key)
  if err != nil {
    panic("Error generating key")
  }
  return key
}

// TODO
// Grab the key generated and push to machine dirs
// Store and call for it when doing decryption
// Sort the master key too
