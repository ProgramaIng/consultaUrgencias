package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/repositories/database"
)

func RequiereRadiografia() {
	var pacienteRequiereRadiografia string
	fmt.Println("El paciente requiere una radiografia? 1. Si 2. No")
	fmt.Scanln(&pacienteRequiereRadiografia)
	fmt.Printf("El paciente: %v, requiere un laboratorio", pacienteRequiereRadiografia)

	if pacienteRequiereRadiografia == "" {
		var tipoDeRadiografia string
		fmt.Println("Cuál es la radiografia que requiere el paciente: 1.cabeza, 2.miembros superiores, 3.miembros inferiores, 4.torax")
		fmt.Scanln(&tipoDeRadiografia)
		fmt.Println("El paciente requiere una Radiografia de:", tipoDeRadiografia)

		switch tipoDeRadiografia {
		case "1":
			RadiografiaCabeza()
		case "2":
			RadiografiaMiembrosSuperiores()
		case "3":
			RadiografiaMiembrosInferiores()
		case "4":
			RadiografiaTorax()

		default:
			println("Valor no valido")
		}
	}

}

func RadiografiaCabeza() {
	dataProcedimientos := database.Getradiografias()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "cabeza" {
			fmt.Println(dataProcedimiento)
		}
	}

}

func RadiografiaMiembrosSuperiores() {
	dataProcedimientos := database.Getradiografias()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "miembros superiores" {
			fmt.Println(dataProcedimiento)
		}
	}

}
func RadiografiaMiembrosInferiores() {
	dataProcedimientos := database.Getradiografias()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "miembros inferiores" {
			fmt.Println(dataProcedimiento)
		}
	}

}

func RadiografiaTorax() {
	dataProcedimientos := database.Getradiografias()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "torax" {
			fmt.Println(dataProcedimiento)
		}
	}

}
