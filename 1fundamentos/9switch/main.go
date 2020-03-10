package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O numero", numero, "se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("voce e da familia do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Nao")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Pode dormir")
	default:
		fmt.Println("Acorda")

	}

	numero = 100
	fmt.Println("Este numero cabe num digito?")
	switch {
	case numero < 10:
		fmt.Println("sim")
	case numero >= 10 && numero <= 99:
		fmt.Println("Serve")
	case numero >= 100:
		fmt.Println("Nao da o numero é muito grande")
	}
}
