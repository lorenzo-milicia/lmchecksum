# lmchecksum validate

Check the validity of a checksum

```
lmchecksum validate <file> <checksum> [flags]
```

## Options

```
      --hash-function <hash function>   set the hash function to use (default sha256)
  -h, --help                            help for validate
```

## Examples

```Bash
$ lmchecksum validate %version%.tar.gz 0b3843af8870caa97753ceddcd90bc494cd9bc07b7ab6c798db56ec31a340de2
[✓] The checksum matches
$ lmchecksum validate %version%.tar.gz "a wrong checksum"                                                
[x] The checksum doesn't match
```

### See also

* [lmchecksum](lmchecksum.md)	 - CLI tool for checking the validity of a checksum

