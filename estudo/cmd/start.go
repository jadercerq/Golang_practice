package cmd

import (
	"time"

	"github.com/estudo/interface"
	"github.com/estudo/logging"
)

func Start() {
	logger()
	ponteiro()
	anonimo()
	colecao()
	oo.Impl()
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
