package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/repositories/database"
)

func RequiereProcedimiento() {
	var pacienteRequiereProcedimiento string
	fmt.Println("El paciente requiere un procedimiento? 1. Si 2. No")
	fmt.Scanln(&pacienteRequiereProcedimiento)
	fmt.Printf("El paciente: %v, requiere un procedimiento", pacienteRequiereProcedimiento)

	if pacienteRequiereProcedimiento == "" {
		var tipoDeProcedimiento string
		fmt.Println("Cuál es el tipo de procedimiento qe requiere el paciente: 1. inmediato, 2. ambulatorio")
		fmt.Scanln(&tipoDeProcedimiento)
		fmt.Println("El paciente requiere un procedimiento", tipoDeProcedimiento)

		switch tipoDeProcedimiento {
		case "1":
			ProcedimientoInmediato()
		case "2":
			ProcedimientoAmbuatorio()
		}
	}

}

func ProcedimientoInmediato() {
	var nombreProcedimientoInmediato string
	fmt.Println("Cuál es el procedimiento inmediato que requiere el paciente: 1.lavado de oido, 2.curacion, 3.sutura, 4.inmovilizacion")
	fmt.Scanln(&nombreProcedimientoInmediato)
	fmt.Println("El paciente requiere un procedimiento", nombreProcedimientoInmediato)

	switch nombreProcedimientoInmediato {
	case "1":
		ProcedimientoLavadoDeOido()
	case "2":
		ProcedimientoCuracion()
	case "3":
		ProcedimientoSutura()
	case "4":
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

	dataProcedimientos := database.GetProcedimientos()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "lavado de oido" {
			fmt.Println(dataProcedimiento)
		}
	}
}

func ProcedimientoCuracion() {

	dataProcedimientos := database.GetProcedimientos()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "curacion" {
			fmt.Println(dataProcedimiento)
		}
	}
}

func ProcedimientoInmovilizacion() {
	dataProcedimientos := database.GetProcedimientos()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "inmovilizacion" {
			fmt.Println(dataProcedimiento)
		}
	}
}

func ProcedimientoSutura() {
	dataProcedimientos := database.GetProcedimientos()
	for key, dataProcedimiento := range dataProcedimientos {
		if key == "sutura" {
			fmt.Println(dataProcedimiento)
		}
	}

}

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
