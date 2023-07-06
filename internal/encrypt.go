
func encryptFile(inputFilePath, outputFilePath string,key []byte) error {
  input, err := ioutil.ReadFile(inputFilePath)
  if err != nil {
    return err
  }

  block, err := aes.NewCipher(key)
  if err != nil {
    return err
  }

  aesGCM, err := cipher.NewGCM(block)
  if err != nil {
    return err
  }

  nonce := make([]byte, aesGCM.NonceSize())
  if _, err := rand.Read(nonce); err != nil {
    return err
  }

  encryptedData := aesGCM.Seal(nil, nonce, input, nil)
  err = ioutil.WriteFile(outputFilePath, encryptedData, 0644)
  if err != nil {
    return err
  }
  return nil
}
