package main

type personal struct {
	nombre         string
	registroMedico int
	id             int
	profesion      string
}

type paciente struct {
	nombre          string
	documento       int
	historiaClinica historiaClinica
}

type historiaClinica struct {
	nombre      string
	documento   int
	diagnostico string
}

type atencion interface {
	execute(paciente)
	sendResoults()
	closed()
	Setnex(atencion)
}

type examenes struct {
	elemento string
	cantidad int
}

func main() {

	procedimientos := make(map[string][]string)

	procedimientos["lavado de oido"] = []string{"125cc de SSN", "1 jeringa de 10cc"}
	procedimientos["curacion"] = []string{"3 gasas", "micropore", "250cc de SSN", "acido fusidico"}
	procedimientos["sutura"] = []string{"1 equipo de sutura", "4 gasas", "prolene 3 ceros", "micropore"}
	procedimientos["inmovilizacion"] = []string{"2 algodon laminado", "2 vendas elasticas", "2 vendaje yeso"}

	laboratorios := make(map[string]examenes)
	laboratorios["cuadro hematico"] = examenes{elemento: "tubo tapa roja", cantidad: 1}
	laboratorios["parcial de orina"] = examenes{elemento: "frasco para parcial de orina", cantidad: 1}
	laboratorios["coproscopico"] = examenes{elemento: "frasco para coproscopico", cantidad: 1}

	radiografias := make(map[string]examenes)
	radiografias["cabeza"] = examenes{elemento: "lamina para rx", cantidad: 1}
	radiografias["miembros Superiores"] = examenes{elemento: "lamina para rx", cantidad: 1}
	radiografias["miembros inferiores"] = examenes{elemento: "lamina para rx", cantidad: 1}
	radiografias["torax"] = examenes{elemento: "lamina para rx", cantidad: 1}
}
