package main

import (
	"bufio"
	"fmt"
	"os"
	pb "github.com/alvarofuentesm/SD-T3/proto"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	
	argsWithoutProg := os.Args[1:]
	
	var nombre = "Thrawn"
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "ahsoka" {
		nombre = "Ahsoka"
	}


	var playing = true

	if len(argsWithoutProg) == 0 {  
		fmt.Println("Bienvenido informante, elija el numero de su nombre:")
		fmt.Println("1. Ahsoka Tano")
		fmt.Println("2. Almirante Thrawn")
		
		opt_user := leer_opt()
		nombre = ""

		switch opt_user {
		case '1':
			nombre = "Ahsoka Tano"
		case '2':
			nombre = "Almirante Thrawn"
		default:
			fmt.Println("Respuesta no valida")
		}
	}

	
	fmt.Println("Saludos, ", nombre)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	//conn, err = grpc.Dial("XX.X.XX.XXX:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewBrokerServiceClient(conn)
	
	// solicitar unirse al juego del calamar 
	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello World!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Reponse: %s", response.Body)


	comando := ""
	planeta := ""
	ciudad := ""
	nuevo_valor := ""

	loop:
		for playing {

			// ---------- ELEGIR COMANDO ----------
			fmt.Println("Escoja el comando que desea enviar:")
			fmt.Println("1. AddCity")
			fmt.Println("2. UpdateName")
			fmt.Println("3. UpdateNumber")
			fmt.Println("4. DeleteCity")
			fmt.Println("5. Salir")

			opt_user := leer_opt()

			switch opt_user {
			case '1':
				comando = "AddCity"
			case '2':
				comando = "UpdateName"
			case '3':
				comando = "UpdateNumber"
			case '4':
				comando = "DeleteCity"
			case '5':
				fmt.Println("Gracias, hasta la próxima.")
				playing = false
				break loop
			default:
				fmt.Println("Respuesta no valida")
			}

			// ---------- ELEGIR PLANETA ----------
			fmt.Println("Introduzca el nombre del planeta:")
			fmt.Scanln(&planeta)

			// ---------- ELEGIR CIUDAD ----------
			fmt.Println("Introduzca el nombre de la ciudad:")
			fmt.Scanln(&ciudad)

			rpta_final := ""

			// ---------- ELEGIR NUEVO_VALOR ----------
			switch comando {
			case "AddCity":
				fmt.Println("Introduzca la cantidad de soldados rebeldes:")
				fmt.Scanln(&nuevo_valor)
			case "UpdateName":
				fmt.Println("Introduzca el nuevo nombre de la ciudad:")
				fmt.Scanln(&nuevo_valor)
			case "UpdateNumber":
				fmt.Println("Introduzca el nuevo numero de rebeldes en la ciudad:")
				fmt.Scanln(&nuevo_valor)
			}

			// ---------- JUNTAR RPTAS ----------
			if comando != "DeleteCity" {
				rpta_final = comando + " " + planeta + " " + ciudad + " " + nuevo_valor
			} else {
				rpta_final = comando + " " + planeta + " " + ciudad
			}

			fmt.Println("El comando que se enviará es:")
			fmt.Println(rpta_final)
			fmt.Println("Enviando...")

			// ENVIAR AL BROKER
			// BROKER RESPONDE CON EL SERVIDOR ALEATORIO A DONDE ENVIARLO
			// ENVIAR AL SERVIDOR QUE EL BROKER ESCOGIO
			// LA RPTA DEL SERVIDOR ES EL RELOJ DE VECTOR

			// GUARDAR EL VECTOR, EL COMANDO FINAL Y LA DIRECCION DEL SERVER QUE SE CONECTO AL ULTIMO
			// dice mantener en memoria y no entiendo a que se refiere
			// mantenerlo en un dict? en un txt? se deberia reiniciar cada vez que se ejecuta este script?? aaaaaaa

			// ---------- SEGUIR EJECUTANDO ----------
			fmt.Println("¿Desea seguir enviando comandos?:")
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
