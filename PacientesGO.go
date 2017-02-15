package main

import (
	"bufio"
	"fmt"
	"os"
	//"io/ioutil"
)

type Paciente struct {
	nombre      string
	id          string
	sintomas    string
	horaLlegada string
	eps         string
}

func (p *Paciente) String() string {
	return fmt.Sprint("Nombre: ", p.nombre, "\n",
		"Identificacion: ", p.id, "\n",
		"Sintomas: ", p.sintomas, "\n",
		"Hora de llegada: ", p.horaLlegada, "\n",
		"EPS: ", p.eps, "\n \n")
	//return fmt.Sprint("Nombre: ", p.nombre)
}

func nuevaCola(tamaño int) *Cola {
	return &Cola{
		pacientes: make([]*Paciente, tamaño),
		tamaño:    tamaño,
	}
}

type Cola struct {
	pacientes                      []*Paciente
	tamaño, cabeza, cola, contador int
}

func (c *Cola) agregar(p *Paciente) {
	if c.cabeza == c.cola && c.contador > 0 {
		pacientes := make([]*Paciente, len(c.pacientes)+c.tamaño)
		copy(pacientes, c.pacientes[c.cabeza:])
		copy(pacientes[len(c.pacientes)-c.cabeza:], c.pacientes[:c.cabeza])
		c.cabeza = 0
		c.cola = len(c.pacientes)
		c.pacientes = pacientes
	}
	c.pacientes[c.cola] = p
	c.cola = (c.cola + 1) % len(c.pacientes)
	c.contador++
}

func (c *Cola) atender() *Paciente {
	if c.contador == 0 {
		return nil
	}

	paciente := c.pacientes[c.cabeza]
	c.cabeza = (c.cabeza + 1) % len(c.pacientes)
	c.contador--
	return paciente
}

func (c *Cola) agregarPaciente(nombre string, id string, sintomas string, horaLlegada string, eps string) {
	fmt.Println("Agregando Paciente...")
	c.agregar(&Paciente{nombre, id, sintomas, horaLlegada, eps})

}

func main() {
	//fmt.Println("Listado pacientes: ")
	cola := nuevaCola(1)
	//eleccion := "agregar"
	//cola.agregar(&Paciente{"Pepe Perez", "AB123", "GRIPE", "08:10", "FAMISANAR"})
	//cola.agregar(&Paciente{"Pedro Perez", "AB124", "GRIPE", "08:12", "SALUDCOOP"})
	//cola.agregar(&Paciente{"Manuel Rojas", "AB125", "INFECCION", "08:15", "FAMISANAR"})
	//cola.agregar(&Paciente{"Sandra Guzman", "AB126", "MALESTAR GENERAL", "08:23", "FAMISANAR"})
	//fmt.Println(cola.atender(), cola.atender(), cola.atender())
	//for i := 0; i < 2; i++ {
	i := 0
	parar := 1
	var input string
	for i < parar {
		fmt.Println("Ingrese la opcion elegida: \n")
		/*in := bufio.NewReader(os.Stdin)
		opcion, err := in.ReadString('\n')

		if err != nil {
			fmt.Println("Hay un error al leer la opcion.")
		}*/

		inpt, _ := fmt.Scanln(&input)

		fmt.Println("Usted digitó: " + string(input))
		//opcion2 := string(opcion)

		//switch opcion2 {
		fmt.Println(inpt)

		switch input {

		case "1":
			fmt.Println("Entró al CASE 1")
			//LEEREMOS EL NOMBRE DEL PACIENTE
			fmt.Println("Ingrese nombre del paciente: \n")
			in := bufio.NewReader(os.Stdin)
			nuevoPacienteNombre, err := in.ReadString('\n')
			if err != nil {
				fmt.Println("Hay un error al leer el nombre del paciente.")
			}
			//LEEREMOS EL DOCUMENTO DE IDENTIFICACION DEL PACIENTE
			fmt.Println("Ingrese identificación del paciente: \n")
			in1 := bufio.NewReader(os.Stdin)
			nuevoPacienteId, err := in1.ReadString('\n')
			if err != nil {
				fmt.Println("Hay un error al leer el nombre del paciente.")
			}
			//LEEREMOS LOS SINTOMAS DEL PACIENTE
			fmt.Println("Ingrese sintomas del paciente: \n")
			in2 := bufio.NewReader(os.Stdin)
			nuevoPacienteSintomas, err := in2.ReadString('\n')
			if err != nil {
				fmt.Println("Hay un error al leer el nombre del paciente.")
			}
			//LEEREMOS LA HORA DE LLEGADA DLE PACIENTE
			fmt.Println("Ingrese hora de llegada del paciente: \n")
			in3 := bufio.NewReader(os.Stdin)
			nuevoPacienteHora, err := in3.ReadString('\n')
			if err != nil {
				fmt.Println("Hay un error al leer el nombre del paciente.")
			}
			//LEEREMOS A QUE EPS PERTENECE EL PACIENTE
			fmt.Println("Ingrese EPS del paciente: \n")
			in4 := bufio.NewReader(os.Stdin)
			nuevoPacienteEPS, err := in4.ReadString('\n')
			if err != nil {
				fmt.Println("Hay un error al leer el nombre del paciente.")
			}

			cola.agregarPaciente(nuevoPacienteNombre, nuevoPacienteId, nuevoPacienteSintomas, nuevoPacienteHora, nuevoPacienteEPS)
			//i = 0
			parar++

		case "2":
			fmt.Println("Entró al CASE 2")
			contador := cola.contador
			for i := 0; i < contador; i++ {
				fmt.Println(cola.atender())
			}
			//i = 2
			parar++
			
		case "3":
			os.Exit(1)

		default:
			fmt.Println("No seleccionó una opción correcta.")
			//i = 0
			parar--

		}

		i++
	}

	contador := cola.contador

	for i := 0; i < contador; i++ {
		//fmt.Println(cola.atender())
		fmt.Print(cola.atender())
	}
}
