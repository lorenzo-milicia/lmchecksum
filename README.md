# About lmchecksum

`lmchecksum` is a small and simple CLI tool for checking the validity of a checksum.

## Installation

You can install `lmchecksum` using `go install`:
```Bash
go install go.lorenzomilicia.dev/lmchecksum/v2@v2.0.0
```

## Checksum validation
To check the validity of a checksum run the command:

```Bash
lmchecksum validate <file name> <checksum>
```
> The default hashing function used is *SHA-256*

To read more check the [full documentation](https://github.lorenzomilicia.dev/lmchecksum).
