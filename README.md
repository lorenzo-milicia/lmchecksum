# About lmchecksum

`lmchecksum` is a small and simple CLI tool for checking the validity of a checksum.

## Installation

You can install `lmchecksum` using `go install`:
```Bash
go install go.lorenzomilicia.dev/lmchecksum/v2@latest
```

## Checksum validation
To check the validity of a checksum run the command:

```Bash
lmchecksum validate <file name> <checksum>
```

## Hash generation
To generate the hash of a given file, use:
```Bash
lmchecksum hash <file name>
```

> The default hashing function used for both `validate` and `hash` is *SHA-256*

To read more check the [full documentation](https://github.lorenzomilicia.dev/lmchecksum).
