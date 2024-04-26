package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/core/domain"
)

func RequiereLaboratorio(laboratorio map[string][]domain.Examenes) {
	var pacienteRequiereLaboratorio string
	fmt.Println("El paciente requiere un laboratorio? 1. Si, 2. No")
	fmt.Scanln(&pacienteRequiereLaboratorio)
	fmt.Printf("El paciente: %v, requiere un laboratorio", pacienteRequiereLaboratorio)

	if pacienteRequiereLaboratorio == "1" {
		var tipoDeLaboratorio string
		fmt.Println("Cuál es el laboratorio que requiere el paciente: 1.cuadro hematico, 2.parcial de orina, 3.coproscopico")
		fmt.Scanln(&tipoDeLaboratorio)
		fmt.Println("El paciente requiere un laboratorio de:", tipoDeLaboratorio)

		if tipoDeLaboratorio == "1" {
			fmt.Println("Los elementos que requiere para el cuadro hematico son:", laboratorio["cuadro hematico"])
		}
		if tipoDeLaboratorio == "2" {
			fmt.Println("Los elementos que requiere para el parcial de orina son:", laboratorio["parcial de orina"])
		}
		if tipoDeLaboratorio == "3" {
			fmt.Println("Los elementos que requiere para el coproscopico son:", laboratorio["coproscopico"])
		} else {
			fmt.Println("El Valor Ingresado es Invalido")

		}
	}

}
