# Mutata 

_Simple, pipelined data mutations_

Mutata makes it easy to mutate input data through a pipeline of common encoding and hashing functions.

- **Input** from a file, input parameter, or cryptographic random data.
- **Mutate** through a pipeline of common encodings and hashes.
- **Output** to a file or direct to the console.

## Quickstart

Encode plaintext to base64
```
$ mutata.exe -input="my string input" base64
bXkgc3RyaW5nIGlucHV0
```

Calculate the MD5 hash of a file
```
$ mutata.exe -infile="README.md" md5 hex
c408d4a41f05106ca299b65310ab5c7d
```

## Supported Mutations

- `base64`
- `bcrypt`
- `hex`
- `md5`
- `sha1`
- `sha256`
