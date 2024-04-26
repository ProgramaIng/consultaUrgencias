package service

import (
	"fmt"

	"github.com/consultaUrgencias/internal/repositories/database"
)

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
