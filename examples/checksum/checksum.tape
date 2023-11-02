Require lmchecksum
Set Theme "Catppuccin Frappe"
Set TypingSpeed 70ms
Set Height 500
Set Width 1200
Sleep 1s
Type "ls"
Sleep 0.3
Enter
Sleep 1s
Type "sha256sum v2.0.0.tar.gz"
Enter
Sleep 1s
Type "lmchecksum validate v2.0.0.tar.gz "
Sleep 0.1s
Hide
Type@1ms "0b3843af8870caa97753ceddcd90bc494cd9bc07b7ab6c798db56ec31a340de2"
Show
Sleep 1s
Enter
Sleep 1.5s
Type 'lmchecksum validate v2.0.0.tar.gz "a wrong checksum"'
Sleep 0.5s
Enter
Sleep 4s

