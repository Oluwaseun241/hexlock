package internal

import (
	"crypto/rand"
	"io/ioutil"
	"os"
  "path/filepath"
)

func getAppDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(homeDir, ".hexlock")

	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}

	return appDir, nil
}

// Generate secret key
func GenerateKey() []byte { 

  appDir, err := getAppDir()
  if err != nil {
    return nil
  }

  keyFilePath := filepath.Join(appDir, "key.txt")

  if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
    key := make([]byte, 32)
    _, err := rand.Read(key)
    if err != nil {
      return nil
    }

    if writeErr := ioutil.WriteFile(keyFilePath, key, 0600); writeErr != nil {
      return nil
    }
  }

  key, readErr := ioutil.ReadFile(keyFilePath)
	if readErr != nil {
		return nil
	}
  
  return key
}

// Get the secret key
func GetKey() []byte {
	appDir, err := getAppDir()
	if err != nil {
		return nil
	}

	keyFilePath := filepath.Join(appDir, "key.txt")
	key, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return nil
	}

	return key
}
