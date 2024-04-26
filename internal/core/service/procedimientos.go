package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/core/domain"
)

func ProcedimientoInmediato(procedimiento map[string][]domain.Examenes) {
	var nombreProcedimientoInmediato string
	fmt.Println("Cuál es el procedimiento inmediato que requiere el paciente: 1.lavado de oido, 2.curacion, 3.sutura, 4.inmovilizacion")
	fmt.Scanln(&nombreProcedimientoInmediato)
	fmt.Println("El paciente requiere un procedimiento", nombreProcedimientoInmediato)

	if nombreProcedimientoInmediato == "1" {
		fmt.Println("Los elementos que requiere para el lavado de oido son:", procedimiento["lavado de oido"])
	}
	if nombreProcedimientoInmediato == "2" {
		fmt.Println("Los elementos que requiere para el curacion son:", procedimiento["curacion"])
	}
	if nombreProcedimientoInmediato == "3" {
		fmt.Println("Los elementos que requiere para la sutura son:", procedimiento["sutura"])
	}
	if nombreProcedimientoInmediato == "4" {
		fmt.Println("Los elementos que requiere para la inmovilizacion son:", procedimiento["inmovilizacion"])
	} else {
		fmt.Println("El Valor Ingresado es Invalido")

	}

}

func ProcedimientoAmbuatorio() {
	var ordenDeProcedimiento string
	fmt.Println("Cuál es el procedimiento ambulatorio que requiere el paciente")
	fmt.Scanln(&ordenDeProcedimiento)
	fmt.Printf("El paciente requiere %s, tramitarlo en su IPS primaria de atencion", ordenDeProcedimiento)

}
