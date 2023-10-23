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

The default hashing function used is *SHA-256*, but you can override it with the option <code>&#8209;algorithm</code>. The current list of available algorithms is: `sha256`, `sha1`, `sha512`, `md5`.