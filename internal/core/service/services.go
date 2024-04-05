package service

import (
	"fmt"
)

func RequiereProcedimiento() {
	var pacienteRequiereProcedimiento bool
	fmt.Println("El paciente requiere un procedimiento?")
	fmt.Scanln(&pacienteRequiereProcedimiento)
	fmt.Printf("El paciente: %v, requiere un procedimiento", pacienteRequiereProcedimiento)

	if pacienteRequiereProcedimiento {
		var tipoDeProcedimiento string
		fmt.Println("Cuál es el tipo de procedimiento qe requiere el paciente: inmediato, ambulatorio")
		fmt.Scanln(&tipoDeProcedimiento)
		fmt.Println("El paciente requiere un procedimiento", tipoDeProcedimiento)

		switch tipoDeProcedimiento {
		case "inmediato":
			ProcedimientoInmediato()
		case "ambulatorio":
			ProcedimientoAmbuatorio()
		}
	}

}

func ProcedimientoInmediato() {
	var nombreProcedimientoInmediato string
	fmt.Println("Cuál es el procedimiento inmediato que requiere el paciente: lavado de oido, curacion, sutura, inmovilizacion")
	fmt.Scanln(&nombreProcedimientoInmediato)
	fmt.Println("El paciente requiere un procedimiento", nombreProcedimientoInmediato)

	switch nombreProcedimientoInmediato {
	case "lavado de oido":
		ProcedimientoLavadoDeOido()
	case "curacion":
		ProcedimientoCuracion()
	case "sutura":
		ProcedimientoSutura()
	case "inmovilzacion":
		ProcedimientoInmovilizacion()
	}
}

func ProcedimientoAmbuatorio() {
	var ordenDeProcedimiento string
	fmt.Println("Cuál es el procedimiento ambulatorio que requiere el paciente")
	fmt.Scanln(&ordenDeProcedimiento)
	fmt.Printf("El paciente requiere %s, tramitarlo en su IPS primaria de atencion", ordenDeProcedimiento)

}

func ProcedimientoLavadoDeOido() {

}

func ProcedimientoCuracion() {

}

func ProcedimientoInmovilizacion() {

}

func ProcedimientoSutura() {

}
