package database

import "github.com/webservice/ingFinal/consulta-urgencias-api/internal/core/domain"

var procedimientos map[string][]domain.Examenes
var laboratorios map[string][]domain.Examenes
var radiografias map[string][]domain.Examenes

func actividades() {
	procedimientos = make(map[string][]domain.Examenes)
	procedimientos["lavado de oido"] = []domain.Examenes{{Elemento: "SSN", Cantidad: 125}, {Elemento: "jeringa de 10cc", Cantidad: 1}}
	procedimientos["curacion"] = []domain.Examenes{{Elemento: "gasas", Cantidad: 3}, {Elemento: "micropore", Cantidad: 1}, {Elemento: "ssn", Cantidad: 250}, {Elemento: "acido fusidico", Cantidad: 1}}
	procedimientos["sutura"] = []domain.Examenes{{Elemento: "equipo de sutura", Cantidad: 1}, {Elemento: "gasas", Cantidad: 4}, {Elemento: "prolene", Cantidad: 1}, {Elemento: "micropore", Cantidad: 1}}
	procedimientos["inmovilizacion"] = []domain.Examenes{{Elemento: "algdon laminado", Cantidad: 2}, {Elemento: "vendas elasticas", Cantidad: 2}, {Elemento: "vendaje de yeso", Cantidad: 2}}

	laboratorios = make(map[string][]domain.Examenes)
	laboratorios["cuadro hematico"] = []domain.Examenes{{Elemento: "tubo tapa roja", Cantidad: 1}}
	laboratorios["parcial de orina"] = []domain.Examenes{{Elemento: "frasco para parcial de orina", Cantidad: 1}}
	laboratorios["coproscopico"] = []domain.Examenes{{Elemento: "frasco para coproscopico", Cantidad: 1}}

	radiografias = make(map[string][]domain.Examenes)
	radiografias["cabeza"] = []domain.Examenes{{Elemento: "lamina para rx", Cantidad: 1}}
	radiografias["miembros Superiores"] = []domain.Examenes{{Elemento: "lamina para rx", Cantidad: 1}}
	radiografias["miembros inferiores"] = []domain.Examenes{{Elemento: "lamina para rx", Cantidad: 1}}
	radiografias["torax"] = []domain.Examenes{{Elemento: "lamina para rx", Cantidad: 1}}

}
