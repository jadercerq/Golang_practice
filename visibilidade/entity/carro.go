package entity

type Car struct {
	Name string `json:"carro"`
	Year int    `json:"ano"`
}

func (c *Car) Drive() {
	println(c.Name, "andou")
	c.Name = "Polo"
}

func (c Car) Cambio() string {
	return "manual"
}
