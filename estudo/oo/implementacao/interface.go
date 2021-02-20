package oo

import (
	"encoding/json"

	"github.com/estudo/oo/entity"
)

func Impl() {
	c := obterCarro()
	c.Drive()
	println(c.Name, c.Year)

	println(toJSON(c))
	c = fromJSON()
	println(c.Name, c.Year)

	checkTipoCambio(c)
}

func checkTipoCambio(car entity.Vehicle) {
	println(car.Cambio())
}

func obterCarro() entity.Car {
	c := entity.Car{
		Name: "New Fiesta",
		Year: 2013,
	}

	return c
}

func toJSON(c entity.Car) string {
	result, _ := json.Marshal(c)

	return string(result)
}

func fromJSON() entity.Car {
	var carro entity.Car
	j := []byte(`{"carro":"Porsche","ano":"2020"}`)
	json.Unmarshal(j, &carro)
	return carro
}
