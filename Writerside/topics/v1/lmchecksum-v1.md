# lmchecksum (%v1%)

Command for checking the validity of a checksum

```
lmchecksum <file> <checksum> [flags]
```

## Options

<code-block>
      --hash-function &lt;hash function&gt;   set the hash function to use (default sha256)
  -h, --help                            help for lmchecksum
  -v, --version                         version for lmchecksum
</code-block>

>The option `--algorithm` still works as an alias for the new `--hash-function`
option, but it is marked as deprecated and will be removed in the next major release.

{style="warning"}

### See also

* [lmchecksum list](lmchecksum-list-v1.md)     - Print list of available hash functions
