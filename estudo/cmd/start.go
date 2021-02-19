package cmd

import (
	"encoding/json"
	"time"

	"github.com/estudo/entity"
	"github.com/estudo/logging"
)

func Start() {
	logger()
	ponteiro()
	anonimo()
	colecao()

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

func logger() {
	logger := logging.New(time.RFC3339, true)

	logger.Log("info", "starting up service")
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")
}

func ponteiro() {
	x := 10

	var z *int = &x
	println(*z)
}

func anonimo() {

	a := func() int {
		return 500
	}

	println(a())
}

func colecao() {
	m := make(map[string]int)

	m["bahia"] = 38
	m["fortaleza"] = 41

	println(m["bahia"])
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
