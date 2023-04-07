## Golang JWT

This project is about JWT tokens in Go,
how to generate and verify that tokens with RS256 algorithm.

#### Commands to generate public and private keys 

`openssl genrsa -out private_key.pem 2048`

`openssl rsa -in private_key.pem -outform PEM -pubout -out public_key.pem`