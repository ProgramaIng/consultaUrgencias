package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/repositories/database"
)

func RequiereLaboratorio() {
	var pacienteRequiereLaboratorio string
	fmt.Println("El paciente requiere un laboratorio? 1. Si, 2. No")
	fmt.Scanln(&pacienteRequiereLaboratorio)
	fmt.Printf("El paciente: %v, requiere un laboratorio", pacienteRequiereLaboratorio)

	if pacienteRequiereLaboratorio == "" {
		var tipoDeLaboratorio string
		fmt.Println("Cuál es el laboratorio que requiere el paciente: 1.cuadro hematico, 2.parcial de orina, 3.coproscopico")
		fmt.Scanln(&tipoDeLaboratorio)
		fmt.Println("El paciente requiere un laboratorio de:", tipoDeLaboratorio)

		switch tipoDeLaboratorio {
		case "1":
			LaboratorioCuadroHematico()
		case "2":
			LaboratorioParcialDeOrina()
		case "3":
			LaboratorioCoproscopico()
		default:
			println("Valor no valido")
		}
	}

}

func LaboratorioCuadroHematico() {
	dataProcedimientos := database.Getlaboratorios()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "cuadro hematico" {
			fmt.Println(dataProcedimiento)
		}
	}

}

func LaboratorioParcialDeOrina() {
	dataProcedimientos := database.Getlaboratorios()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "parcial de orina" {
			fmt.Println(dataProcedimiento)
		}
	}

}
func LaboratorioCoproscopico() {
	dataProcedimientos := database.Getlaboratorios()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "coproscopico" {
			fmt.Println(dataProcedimiento)
		}
	}

}
