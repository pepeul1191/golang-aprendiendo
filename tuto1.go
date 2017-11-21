package main

import "fmt"

func main(){
	fmt.Println("hola mundo")
	var edad int = 29
	var peso float32 = 88.3  
	cualquier := "1"
	fmt.Println(cualquier)
	fmt.Println("Edad : ", edad, " - Peso : ", peso)
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	fmt.Println("6 - 4 = ", 6 - 4)
	fmt.Println("6 + 4 = ", 6 + 4)
	fmt.Println("6 * 4 = ", 6 * 4)
	fmt.Println("6 / 4 = ", 6 / 4)
	fmt.Println("6 % 4 = ", 6 % 4)
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	const pi float64 = 3.14159265
	var miNombre = "Pepe Valdivia"
	fmt.Println(miNombre, " - n = ", len(miNombre))
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	var i int = 1
	for (i <= 10) {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	for k := 0; k < 5; k++{
		fmt.Println(k)
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	myEdad := 9
	if myEdad >= 18 && myEdad < 60 {
		fmt.Println("Eres mayor de edad")
	} else if myEdad >= 60 {
		fmt.Println("Estas viejo!!!")
	} else {
		fmt.Println("Eres menor de edad")
	}
}