# Examples

Let's verify the checksum of the compressed source code of the [%version% release](https://github.com/lorenzo-milicia/lmchecksum/archive/refs/tags/%version%.tar.gz)

<img src="checksum.gif" alt="Checksum" border-effect="rounded"/>

```Bash
$ ls
%version%.tar.gz
$ sha256sum %version%.tar.gz
a1a2cc03f7aa392fe03109aeceeaec396c4fb0dd1f58c0e5b5f70ebaa0e988a8 lmchecksum %version%.tar.gz
$ lmchecksum %version%.tar.gz a1a2cc03f7aa392fe03109aeceeaec396c4fb0dd1f58c0e5b5f70ebaa0e988a8
[✓] The checksum matches
$ lmchecksum %version%.tar.gz "a wrong checksum"                                                
[x] The checksum doesn't match
```
{collapsible="true"}