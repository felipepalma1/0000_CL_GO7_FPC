package main

import "fmt"

func main() {
	var estudiantes [3]string
	/*Definimos un arreglo de tres elementos. Si quisieramos crear un arreglo sin definir la cantidad máxima de elementos, debemos declarar `var estudiantes [3]string`*/
	fmt.Printf("Estudiantes: %v\n", estudiantes)

	estudiantes[0] = "Lisa Lisa"
	fmt.Printf("Estudiantes %v \n", estudiantes[0])
	estudiantes[1] = "El Meme"
	fmt.Printf("Estudiantes %v \n", estudiantes[1])
	estudiantes[2] = "Arnold"
	fmt.Printf("Estudiantes %v \n", estudiantes[2])

	var matrizIdentidad [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(matrizIdentidad)

	var otraMatrizIdentidad [3][3]int
	otraMatrizIdentidad[0] = [3]int{1, 0, 0}
	otraMatrizIdentidad[1] = [3]int{0, 1, 0}
	otraMatrizIdentidad[2] = [3]int{0, 0, 1}
	fmt.Println(otraMatrizIdentidad)

	/*A continuación, se copian los valores de un arreglo hacia otro. No es por referencia. Es por valor*/
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{3, 3, 3}
	d := &c // Como se hace una asignación POR REFERENCIA, en el siguiente paso que define d[0] = 81, cambiará c[0] a el nuevo valor 81. Este es el poder de los punteros.
	d[0] = 81

	fmt.Println("c", c)
	fmt.Println("d", d)

	/* Slice -- Todo lo que podemos hacer en un arreglo lo podemos hacer en un slice, con un par de excepciones */

}
