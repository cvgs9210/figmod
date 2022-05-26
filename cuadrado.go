package figmod

type Cuadrado struct {
	Ancho, Alto float64
}

func (c *Cuadrado) area() float64 {
	return c.Alto * c.Ancho
}

func (c *Cuadrado) perimetro() float64 {
	return 2*c.Ancho + 2*c.Alto
}
