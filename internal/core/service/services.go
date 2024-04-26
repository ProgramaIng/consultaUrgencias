package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/core/domain"
	"github.com/consultaUrgencias/internal/repositories/database"
)

func Consulta(a domain.Paciente) string {

	var identificacion int
	fmt.Println("Ingrese el numero de documento del paciente:")
	fmt.Scanln(&identificacion)
	fmt.Println("El numero de identificacion del paciente es", identificacion)

	for _, dataPaciente := range database.Getusuarios() {
		if a.Documento == identificacion {
			fmt.Println(dataPaciente)
		}
	}
	return a.Nombre

}

func RequiereProcedimiento() {
	var pacienteRequiereProcedimiento string
	fmt.Println("El paciente requiere un procedimiento? 1. Si 2. No")
	fmt.Scanln(&pacienteRequiereProcedimiento)
	fmt.Printf("El paciente: %v, requiere un procedimiento", pacienteRequiereProcedimiento)

	if pacienteRequiereProcedimiento == "1" {
		var tipoDeProcedimiento string
		fmt.Println("Cuál es el tipo de procedimiento qe requiere el paciente: 1. inmediato, 2. ambulatorio")
		fmt.Scanln(&tipoDeProcedimiento)
		fmt.Println("El paciente requiere un procedimiento", tipoDeProcedimiento)

		switch tipoDeProcedimiento {
		case "1":
			ProcedimientoInmediato()
		case "2":
			ProcedimientoAmbuatorio()

		default:
			println("Valor no valido")
		}
	}

}
