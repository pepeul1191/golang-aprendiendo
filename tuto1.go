package main

import "fmt"
import "strconv"

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
	favNumbers2 := [3]float64 {1.2, 5, 8}
	fmt.Println(favNumbers2)
	for i, value := range favNumbers2 {
		fmt.Println(i, " - ", value)		
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
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	var favNumbers[3] float64
	favNumbers[0] = 12 
	favNumbers[1] = 2.3
	favNumbers[2] = 6.45
	fmt.Println(favNumbers)
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	presAge := make(map[string] int)
	presAge["Pepito"] = 29
	presAge["Yacky"] = 27
	fmt.Println(presAge)
	fmt.Println(presAge["Yacky"])
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	listNum := []float64{1,2,3,4,5}
	fmt.Println("SUMA : ", sumaArreglo(listNum))
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	promedio := func() float32{
		var x1 float32 = 1
		var x2 float32 = 10
		var x3 float32 = 11
			fmt.Println("1 --------")
			fmt.Println(x1 + x2 + x3)
			fmt.Println("2 --------")
		return (x1 + x2 + x3)/3
	}
	fmt.Println(promedio())
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	fmt.Println("Factorial : ", factorial(5))
	fmt.Println("++++++++++++++++++++++++++++++++++++")
	permiso1 := Permiso{id : 1, nombre : "Crear Usuario", llave : "crear_usuario"}
	fmt.Println(permiso1)
	permiso2 := Permiso{2, "Editar Usuario", "editar_usuario"}
	fmt.Println(permiso2)
	fmt.Println(permiso2.getLlave())
	rol1 := Rol{1, "Gestor de Usuarios"}
	fmt.Println(rol1.toString())
}

func sumaArreglo(numbres []float64) float64{ 
	sum := 0.0
	for _, val := range numbres{
		sum = sum + val
	}
	return sum
}

func factorial(num int) int{
	if num == 0{
		return 1
	}
	return num * factorial(num - 1)
}
//++++++++++++++++++++++++++++++++++++++++++++++++
type demoInterface interface{
	toString() string
}

type Permiso struct {
	id int
	nombre string
	llave string
}

func (permiso *Permiso) getLlave() string{
	return "Llave : " + permiso.llave
}

type Rol struct{
	id int
	nombre string
}

func (rol *Rol) toString() string{
	return "id : " + strconv.Itoa(rol.id) + ", nombre : " + rol.nombre
}
//++++++++++++++++++++++++++++++++++++++++++++++++
// 38'44