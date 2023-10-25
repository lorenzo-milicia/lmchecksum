# Examples

Let's verify the checksum of the compressed source code of the [%version% release](https://github.com/lorenzo-milicia/lmchecksum/archive/refs/tags/%version%.tar.gz)
```Bash
$ ls
%version%.tar.gz
$ sha256sum %version%.tar.gz
9ec0bed7e084cdabd5c1d49566622e463dbdd3492fe869fff960667c70751aa6 lmchecksum %version%.tar.gz
$ lmchecksum %version%.tar.gz 9ec0bed7e084cdabd5c1d49566622e463dbdd3492fe869fff960667c70751aa6
It checks out ✅
$ lmchecksum %version%.tar.gz a-wrong-checksum                                                
Checksums don't match ❌
```