# JWT Cracker written in Golang
This is a simple tool used to bruteforce HMAC secret keys in JWT tokens.
It heavly relies on this JWT library: `github.com/dgrijalva/jwt-go`

This tool supports both wordlist and bruteforce based attacks.

## Installation
Build packages are available for the majority of the platforms, but if you want to hack it or build it yourself:

```
git clone https://github.com/riccardoancarani/go-jwt-cracker.git
cd go-jwt-cracker
export GOPATH=$(pwd)
go get github.com/dgrijalva/jwt-go
cd src/app
go install
```
You'll find the binary under the `bin` directory.

## Usage
This is a CLI tool, this means that it is meant to be used from your shell/terminal.

### Main options
- `--token`: The token you want to crack
- `--brute`: Start the brute force attack
- `--wordlist <file>`: The file for wordlist attack
- `--charset <charset>`: Specify the charset to use in the bruteforce attack
- `--max`: The upper limit of the string's lenght for the brute force attack
