Output ../checksum.txt
Require lmchecksum
Set Theme "Catppuccin Frappe"
Set TypingSpeed 40ms
Set Height 500
Sleep 1s
Type "ls"
Sleep 0.3
Enter
Sleep 1s
Type "sha256sum v1.0.1.zip"
Enter
Sleep 1s
Type "lmchecksum v1.0.1.zip "
Sleep 0.1s
Hide
Type@1ms "a1a2cc03f7aa392fe03109aeceeaec396c4fb0dd1f58c0e5b5f70ebaa0e988a8"
Show
Sleep 1s
Enter
Sleep 1.5s
Type 'lmchecksum v1.0.1.zip "a wrong checksum"'
Sleep 0.5s
Enter
Sleep 2s

