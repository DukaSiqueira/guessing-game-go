package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Boas vindas ao jogo da advinhação!")
	fmt.Println("Você deve tentar adivinhar o número secreto!")
	fmt.Println("Você tem 10 tentativas para acertar o número secreto!")
	fmt.Println("Boa sorte!")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Printf("Tentativa %d: ", i+1)
		fmt.Println("Digite um número:")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("Você deve digitar um número válido!")
			return
		}

		switch {
			case chuteInt < x:
				fmt.Println("Seu chute foi menor que o número secreto!")
			case chuteInt > x:
				fmt.Println("Seu chute foi maior que o número secreto!")
			case chuteInt == x:
				fmt.Println("Parabéns! Você acertou o número secreto!")
				fmt.Printf("O número secreto era %d\n", x)
				fmt.Printf("Seus chutes foram: %v\n", chutes[:i])
				fmt.Printf("Você acertou na tentativa %d\n", i+1)
				return
		}

		chutes[i] = chuteInt
	}

	fmt.Println("Você perdeu!")
	fmt.Printf("O número secreto era %d\n", x)
	fmt.Printf("Seus chutes foram: %v\n", chutes)
}