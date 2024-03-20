package main

import "github.com/webservice/ingFinal/consulta-urgencias-api/internal/core/domain"

type atencion interface {
	execute(domain.Paciente)
	sendResoults()
	closed()
	Setnex(atencion)
}

func main() {

}
