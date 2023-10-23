# About %product%

`%product%` is a small and simple CLI tool for checking the validity of a checksum.

## Installation

You can install `%product%` using `go install`:
```Bash
go install go.lorenzomilicia.dev/%product%@%version%
```

## How to use

The command to run is:
```Bash
%product% <file name> <checksum>
```
> The default hashing function used is *SHA-256*
## Hashing functions

The default hashing function used is *SHA-256*, but you can override it with the option `-algorithm`. The current list of available algorithms is:

| Hashing Function | `-algorithm` value |
| ---------------- | ------------ |
| SHA-256          | sha256       |
| SHA-1            | sha1         |
| SHA-512          | sha512       |
| MD5              | md5          |