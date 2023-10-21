# Examples

Let's verify the checksum of the compressed source code of the [v1.0.0 release](https://github.com/lorenzo-milicia/lmchecksum/archive/refs/tags/v1.0.0.tar.gz)
```Bash
$ ls
v1.0.0.tar.gz
$ sha256sum v1.0.0.tar.gz 
$ lmchecksum v1.0.0.tar.gz 9ec0bed7e084cdabd5c1d49566622e463dbdd3492fe869fff960667c70751aa6
It checks out ✅
$ lmchecksum v1.0.0.tar.gz a-wrong-checksum                                                
Checksums don't match ❌
```