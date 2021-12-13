package main

import (
	"bufio"
	"fmt"
	"os"
)

func leer_opt() rune { // Para leer numeros
	reader := bufio.NewReader(os.Stdin)
	opt, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	return opt
}
func main() {
	var playing = true
	//var opt_user = ""
	fmt.Println("Bienvenida Leia Organa")

	planeta := ""
	ciudad := ""

	for playing {
		fmt.Println("Proceda a armar su consulta:")

		// ---------- ELEGIR PLANETA ----------
		fmt.Println("Introduzca el nombre del planeta a consultar:")
		fmt.Scanln(&planeta)

		// ---------- ELEGIR CIUDAD ----------
		fmt.Println("Introduzca el nombre de la ciudad a consultar:")
		fmt.Scanln(&ciudad)

		rpta_final := "GetNumberRebels " + planeta + " " + ciudad
		fmt.Println("El comando que se enviará es:")
		fmt.Println(rpta_final)
		fmt.Println("Enviando...")

		// ENVIAR AL BROKER
		// BROKER RESPONDE SI ES QUE HAY ACTUALIZACIONES PARA EL PLANETA
		// SINO, LEIA LEE DESDE SU MEMORIA QUE VA GUARDANDO (monotonic reads) ??????
		// GUARDA SU CONSULTA EN MEMORIA JUNTO AL RELOJ DE VECTOR Y EL SERVER QUE LE RESPONDIO POR ULTIMA VEZ

		// ---------- SEGUIR EJECUTANDO ----------
		fmt.Println("¿Desea seguir consultando?:")
		fmt.Println("1. Si")
		fmt.Println("2. No")

		opt_exit := leer_opt()

		switch opt_exit {
		case '1':
			fmt.Println("Se reiniciará la interfaz...")
		case '2':
			fmt.Println("Gracias, hasta la próxima.")
			playing = false
			break
		default:
			fmt.Println("Respuesta no valida")
		}

	}
}
