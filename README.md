# Hexlock

HexLock is a command-line tool for file encryption and
compression using AES encryption and gzip compression.

![demo img](https://github.com/Oluwaseun241/hexlock/blob/main/images/demo)

## Features

- Encrypt files using AES encryption algorithm
- Decrypt encrypted files
- Compress files using gzip compression

- [x] Encrypting and Decrypting multiple files at once
- [x] Progress bar

## Installation

```
go install github.com/Oluwaseun241/hexlock
```

## Usage

`./hexlock [options]`

```
# Encrypt file
./hexlock encrypt -i [filepath] -o [filepath]
# Decrypt an encrypted file
./hexlock decrypt -i [filepath] -o [filepath]
# Compress file
./hexlock compress -i [filepath] -o [filepath]
```

**NOTE**

> For multiple files separate by comma

> **To be able to decrypt from another computer, please provide your own key when encrypting**

## License

This project is licensed under the [MIT Licence](https://github.com/Oluwaseun241/hexlock/blob/cobra/LICENCE).
