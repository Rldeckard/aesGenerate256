# aesGenerate32
Small script to generate a private key and encrypt a string with AES. Written in Go.

You can view the code in the .go file or just download the .exe and run it. It'll spit out your private key that you need to access from your code in some way and the encrypted string


```

PS C:> .\aesGenerate.exe
Would you like to provide a key?[y/n]: n
Enter string to encrypt: potatoIcecream
Checking cipher
Encrypting...
Encrypted String: 38628970ec99bfbb22ee585838993e589dc9de28be0be57a52eacb64f1d364e85049bf987dc7f641ed1c
Decryption Key (do not lose): Z9gErvjnRpil9U09LWOczMzm9mZkHA==

Testing Secret.
Enter Decryption Key (copy from above): Z9gErvjnRpil9U09LWOczMzm9mZkHA==
Is this your card? potatoIcecream

```
