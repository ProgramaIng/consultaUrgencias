package main

import (
	"github.com/consultaUrgencias/internal/core/domain"
	"github.com/consultaUrgencias/internal/server"
)

type atencion interface {
	execute(domain.Paciente)
	sendResoults()
	closed()
	Setnex(atencion)
}

func main() {
	server.Start()
}
