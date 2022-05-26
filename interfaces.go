package figmod

import "fmt"

type figurageo interface {
	area() float64
	perimetro() float64
}

func Medidas(fg figurageo) {
	fmt.Println("Figura: ", fg)
	fmt.Println("Area:", fg.area())
	fmt.Println("Perimetro:", fg.perimetro())
}
