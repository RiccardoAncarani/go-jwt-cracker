package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {

	var tokenString = flag.String("token", "", "The JWT token you want to crack")
	var wordlist = flag.String("wordlist", "", "The wordlist you want to use")
	var brute = flag.Bool("brute", false, "Use bruteforce mode")
	var charset = flag.String("charset", "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM", "The charset to use during bruteforce")
	var maxChar = flag.Int("max", 12, "Max chars in token")

	flag.Parse()

	if *wordlist != "" && *brute == false {
		wordlistFile, _ := os.Open(*wordlist)
		defer wordlistFile.Close()

		scanner := bufio.NewScanner(wordlistFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			currentToken := scanner.Text()
			if validateToken(*tokenString, currentToken) {
				fmt.Println("[+] Valid secret found: " + currentToken)
				return
			}
		}
	}

	if *brute == true {
		for combination := range GenerateCombinations(*charset, *maxChar) {
			if validateToken(*tokenString, combination) {
				fmt.Println("[+] Valid secret found: " + combination)
				return
			}
		}
	}

}

// I found this function in a stackoverflow answer that I'm not finding anymore
// when I'll find it I'll give appropriate credits to the author
func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c) 
		
		AddLetter(c, "", alphabet, length) 
	}(c)
	return c 
}

func AddLetter(c chan string, combo string, alphabet string, length int) {
	if length <= 0 {
		return
	}
	
	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		AddLetter(c, newCombo, alphabet, length-1)
	}	
}



func validateToken(tokenString string, secret string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte(secret)
		return hmacSampleSecret, nil
	})

	if token.Valid {
		return true
	} else {
		return false
	}
}
