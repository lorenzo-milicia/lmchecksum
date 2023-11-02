# Examples

Let's verify the checksum of the compressed source code of the [%version% release](https://github.com/lorenzo-milicia/lmchecksum/archive/refs/tags/%version%.tar.gz)

<img src="checksum.gif" alt="Checksum" border-effect="rounded"/>

```Bash
$ ls
%version%.tar.gz
$ sha256sum %version%.tar.gz
0b3843af8870caa97753ceddcd90bc494cd9bc07b7ab6c798db56ec31a340de2 %version%.tar.gz
$ lmchecksum validate %version%.tar.gz 0b3843af8870caa97753ceddcd90bc494cd9bc07b7ab6c798db56ec31a340de2
[✓] The checksum matches
$ lmchecksum validate %version%.tar.gz "a wrong checksum"                                                
[x] The checksum doesn't match
```
{collapsible="true" collapsed-title="lmchecksum"}