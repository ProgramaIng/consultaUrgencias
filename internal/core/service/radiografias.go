package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/core/domain"
)

func RequiereRadiografia(radiografia map[string][]domain.Examenes) {
	var pacienteRequiereRadiografia string
	fmt.Println("El paciente requiere una radiografia? 1. Si 2. No")
	fmt.Scanln(&pacienteRequiereRadiografia)
	fmt.Printf("El paciente: %v, requiere un laboratorio", pacienteRequiereRadiografia)

	if pacienteRequiereRadiografia == "1" {
		var tipoDeRadiografia string
		fmt.Println("Cuál es la radiografia que requiere el paciente: 1.cabeza, 2.miembros superiores, 3.miembros inferiores, 4.torax")
		fmt.Scanln(&tipoDeRadiografia)
		fmt.Println("El paciente requiere una Radiografia de:", tipoDeRadiografia)

		if pacienteRequiereRadiografia == "1" {
			fmt.Println("Los elementos que requiere para la radiografia de cabeza son:", radiografia["cabeza"])
		}
		if pacienteRequiereRadiografia == "2" {
			fmt.Println("Los elementos que requiere para la radiografia de miembros superiores son:", radiografia["miembros superiores"])
		}
		if pacienteRequiereRadiografia == "3" {
			fmt.Println("Los elementos que requiere para la radiografia de miembros inferiores son:", radiografia["miembros inferiores"])
		}
		if pacienteRequiereRadiografia == "4" {
			fmt.Println("Los elementos que requiere para la radiografia de torax son:", radiografia["torax"])
		} else {
			fmt.Println("El Valor Ingresado es Invalido")

		}

	}

}
