# Hexlock

HexLock is a command-line tool for file encryption and
compression using AES encryption and gzip compression.

![demo](https://github.com/Oluwaseun241/images/demo.png)

## Features

- Encrypt files using AES encryption algorithm
- Decrypt encrypted files
- Compress files using gzip compression

- [x] Encrypting and Decrypting multiple files at once(if possible a dir)
- [x] Progress bar

## Usage

`./hexlock [options]`

```
# Encrypt file
./hexlock -i [filepath] -o [filepath] -mode encrypt
# Decrypt an encrypted file
./hexlock -i [filepath] -o [filepath] -mode decrypt
# Compress file
./hexlock -i [filepath] -o [filepath] -mode compress

```

> For multiple files separate by comma

## License

This project is licensed under the [MIT Licence](https://github.com/Oluwaseun241/hexlock/blob/cobra/LICENCE).
