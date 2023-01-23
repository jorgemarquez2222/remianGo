package main

import (
	"fmt"

	funciones "github.com/jorgemarquez2222/remianGo/funciones"
)

type Persona struct {
	Name  string
	Email string
	Age   int
}

func main() {
	hombre := &Persona{Name: "Jorge", Email: "algo@gmail.com", Age: 33}
	mapa := make(map[string]string)
	slice_ := make([]string, 0, 0)
	hombre.Name = "Pepe"
	fmt.Println(*hombre)
	mapa["clave"] = "Jeannelis"
	mapa["clave1"] = "Jeannelis1"
	mapa["clave2"] = "Jeannelis2"
	mapa["clave3"] = "Jeannelis3"
	delete(mapa, "clave")
	for _, v := range mapa {
		fmt.Println(v)
	}
	slice_ = append(slice_, "Hoal")
	slice_ = append(slice_, "como")
	slice_ = append(slice_, "estas")
	for _, v := range slice_ {
		fmt.Println(v)
	}
	fmt.Println(funciones.Suma(3, 4))

}
