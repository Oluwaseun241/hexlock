# Hexlock

HexLock is a command-line tool for file encryption and
compression using AES encryption and gzip compression.

## Features

- Encrypt files using AES encryption algorithm
- Decrypt encrypted files
- Compress files using gzip compression

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

## Todo

- Encrypting and Decrypting multiple files at once(if possible a dir)
- Progress bar
