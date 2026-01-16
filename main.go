package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
	_ "math/rand"
	"time"
)

func GeneratePassword(length int, characters string) string {
	var password string
	for i := 0; i < length; i++ {
		n := rand.Intn(len(characters))
		password += string(characters[n])
	}
	return password
}

func PrintPasswords(password_keeper map[string]string) {
	fmt.Println("Все ваши пароли:")

	for key, value := range password_keeper {
		fmt.Printf("%s:\t%s\n", key, value)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	var password_keeper = make(map[string]string)

	for {
		var length int
		var service string
		characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		fmt.Print("Введите название сервиса, для которого вы хотите получить пароль: ")
		fmt.Scan(&service)

		fmt.Print("Введите длину генериуремого пароля: ")
		fmt.Scan(&length)

		if length <= 0 {
			fmt.Println("Длина пароля должна быть больше нуля")
			continue
		}

		var digits string
		var symbols string

		fmt.Print("Хотите добавить цифры? (0123456789) (y/n): ")
		fmt.Scan(&digits)

		if digits != "n" {
			characters += "0123456789"
		}

		fmt.Print("Хотите добавить спец.символы? (!@#$%^&*()_+-=[]{}|;:',.<>?/) (y/n): ")
		fmt.Scan(&symbols)

		if symbols != "n" {
			characters += "!@#$%^&*()_+-=[]{}|;:',.<>?/"
		}

		password_keeper[service] = GeneratePassword(length, characters)

		fmt.Printf("Пароль к сервису %s: %s\n", service, password_keeper[service])

		var fin string

		fmt.Print("Хотите продолжить? (y/n): ")
		fmt.Scan(&fin)

		if fin == "n" {
			break
		}
	}

	PrintPasswords(password_keeper)

}
