package main

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
)

func GeneratePassword(length int, characters string) (string, error) {
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		result[i] = characters[n.Int64()]
	}
	return string(result), nil
}

func PrintPasswords(passwordKeeper map[string]string) {
	fmt.Println("All your passwords:")

	for key, value := range passwordKeeper {
		fmt.Printf("%s:\t%s\n", key, value)
	}

}

func WriteData(file *os.File, passwordKeeper map[string]string, service string) {
	writer := csv.NewWriter(file)

	//if isFirst {
	//	if err := writer.Write([]string{"Service", "Password"}); err != nil {
	//		fmt.Println(err)
	//	}
	//}

	err := writer.Write([]string{service, passwordKeeper[service]})

	if err != nil {
		fmt.Println(err)
	}

	writer.Flush()
}

func main() {

	var passwordKeeper = make(map[string]string)
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
	file, err := os.Create("passwords.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for {
		var length int
		var service string

		fmt.Print("Enter service name for which you want to get password: ")
		fmt.Scan(&service)

		fmt.Print("Enter length of a password: ")
		fmt.Scan(&length)

		if length <= 0 {
			log.Fatal("Error: length of password must be greater than zero!")
		}

		var digits string
		var symbols string

		fmt.Print("Do you want add digits to your password? (0123456789) (y/n): ")
		fmt.Scan(&digits)

		if digits != "n" {
			characters += "0123456789"
		}

		fmt.Print("Do you want add special characters to your password? (!@#$%^&*()_+-=[]{}|;:',.<>?/) (y/n): ")
		fmt.Scan(&symbols)

		if symbols != "n" {
			characters += "!@#$%^&*()_+-=[]{}|;:',.<>?/"
		}

		passwordKeeper[service], err = GeneratePassword(length, characters)

		fmt.Printf("Password to the service %s: %s\n", service, passwordKeeper[service])

		WriteData(file, passwordKeeper, service)

		var fin string

		fmt.Print("Do you want to continue? (y/n): ")
		fmt.Scan(&fin)

		if fin == "n" {
			break
		}
	}

	PrintPasswords(passwordKeeper)

}
