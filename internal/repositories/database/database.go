package database

import "github.com/consultaUrgencias/internal/core/domain"

var (
	// Procedimientos ...
	procedimientos = map[string][]domain.Examenes{
		"lavado de oido": {{Elemento: "SSN", Cantidad: 125}, {Elemento: "jeringa de 10cc", Cantidad: 1}},
		"curacion":       {{Elemento: "gasas", Cantidad: 3}, {Elemento: "micropore", Cantidad: 1}, {Elemento: "ssn", Cantidad: 250}, {Elemento: "acido fusidico", Cantidad: 1}},
		"sutura":         {{Elemento: "equipo de sutura", Cantidad: 1}, {Elemento: "gasas", Cantidad: 4}, {Elemento: "prolene", Cantidad: 1}, {Elemento: "micropore", Cantidad: 1}},
		"inmovilizacion": {{Elemento: "algdon laminado", Cantidad: 2}, {Elemento: "vendas elasticas", Cantidad: 2}, {Elemento: "vendaje de yeso", Cantidad: 2}},
	}

	// Laboratorios ...
	laboratorios = map[string][]domain.Examenes{
		"cuadro hematico":  {{Elemento: "tubo tapa roja", Cantidad: 1}},
		"parcial de orina": {{Elemento: "frasco para parcial de orina", Cantidad: 1}},
		"coproscopico":     {{Elemento: "frasco para coproscopico", Cantidad: 1}},
	}

	radiografias = map[string][]domain.Examenes{
		"cabeza":              {{Elemento: "lamina para rx", Cantidad: 1}},
		"miembros Superiores": {{Elemento: "lamina para rx", Cantidad: 1}},
		"miembros inferiores": {{Elemento: "lamina para rx", Cantidad: 1}},
		"torax":               {{Elemento: "lamina para rx", Cantidad: 1}},
	}
)

func getProcedimientos() map[string][]domain.Examenes {
	return procedimientos
}
func getlaboratorios() map[string][]domain.Examenes {
	return laboratorios
}

func getradiografias() map[string][]domain.Examenes {
	return radiografias
}
