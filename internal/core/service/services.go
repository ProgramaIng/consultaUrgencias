package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/core/domain"
	"github.com/consultaUrgencias/internal/repositories/database"
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
	getProcedimientos()
	for "lavado de oido", []domain.Examenes := range map [string]database.procedimientos {
		fmt.Println("Los elementos que requiere para el lavado de oido son:",["lavado de oido"] domain.Examenes)
	}
		
}

func ProcedimientoCuracion(){
getProcedimientos()
for "curacion", []domain.Examenes := range map [string]database.procedimientos {
	fmt.Println("Los elementos que requiere para la curacion son:",["curacion"] domain.Examenes)
}

}

func ProcedimientoInmovilizacion() {

}

func ProcedimientoSutura() {

}

func RequiereLaboratorio() {
	var pacienteRequiereLaboratorio bool
	fmt.Println("El paciente requiere un laboratorio?")
	fmt.Scanln(&pacienteRequiereLaboratorio)
	fmt.Printf("El paciente: %v, requiere un laboratorio", pacienteRequiereLaboratorio)

	if pacienteRequiereLaboratorio {
		var tipoDeLaboratorio string
		fmt.Println("Cuál es el laboratorio que requiere el paciente: cuadro hematico, parcial de orina, coproscopico")
		fmt.Scanln(&tipoDeLaboratorio)
		fmt.Println("El paciente requiere un laboratorio de:", tipoDeLaboratorio)

		switch tipoDeLaboratorio {
		case "cuadro hematico":
			LaboratorioCuadroHematico()
		case "parcial de orina":
			LaboratorioParcialDeOrina()
		case "coproscopico":
			LaboratorioCoproscopico()
		}
	}

}

func LaboratorioCuadroHematico() {
	
}

func LaboratorioParcialDeOrina() {
	
}
func LaboratorioCoproscopico() {
	
}

func RequiereRadiografia() {
	var pacienteRequiereRadiografia bool
	fmt.Println("El paciente requiere una radiografia?")
	fmt.Scanln(&pacienteRequiereRadiografia)
	fmt.Printf("El paciente: %v, requiere un laboratorio", pacienteRequiereRadiografia)

	if pacienteRequiereRadiografia {
		var tipoDeRadiografia string
		fmt.Println("Cuál es la radiografia que requiere el paciente: cabeza, miembros superiores, miembros inferiores, torax")
		fmt.Scanln(&tipoDeRadiografia)
		fmt.Println("El paciente requiere una Radiografia de:", tipoDeRadiografia)

		switch tipoDeRadiografia {
		case "cabeza":
			RadiografiaCabeza()
		case "miebros superiores":
			RadiografiaMiembrosSuperiores()
		case "miembros inferiores":
			RadiografiaMiembrosInferiores()
		case "torax":
			RadiografiaTorax()
		}
	}

}

func RadiografiaCabeza() {
	
}

func RadiografiaMiembrosSuperiores() {
	
}
func RadiografiaMiembrosInferiores() {
	
}

func RadiografiaTorax() {
	
}
