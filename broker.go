package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	var listening = true
	fmt.Println("Bienvenido al broker Mos Eisley")

	server_list := []int{1, 2, 3} // Deberian ir las direcciones de los server

	for listening {

		fmt.Println("Esperando peticion...")
		// ESPERAR PETICION...???
		// llega una
		peticion := ""

		randomIndex := rand.Intn(len(server_list))
		pick := server_list[randomIndex]
		fmt.Println("El servidor elegido para la consulta es:", pick)
		fmt.Println(pick)

		if strings.Split(peticion, " ")[0] == "GetNumberRebels" { // Es una peticion de Leia porque es la unica que puede hacer de ese tipo
			// MANDAR LA PETICION DE LEIA AL SERVER ELEGIDO
			fmt.Println("a")
			// RECIBIR LA RPTA DE LOS SERVER (y el vector)
			// MANDARSELA A LEIA
		} else {
			// ENVIARLE A LOS INFORMANTES EL SERVER QUE LE TOCA, ellos se encargan de enviar por si mismos
			fmt.Println("a")
		}

		// ---------- SEGUIR EJECUTANDO ----------
		fmt.Println("¿Desea seguir esperando comandos?:")
		fmt.Println("1. Si")
		fmt.Println("2. No")

		opt_exit := leer_opt()

		switch opt_exit {
		case '1':
			fmt.Println("Se reiniciará la interfaz...")
		case '2':
			fmt.Println("Gracias, hasta la próxima.")
			listening = false
			break
		default:
			fmt.Println("Respuesta no valida")
		}
	}

	//
}
