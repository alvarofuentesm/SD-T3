package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/alvarofuentesm/SD-T3/proto"
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

// para convertir un vector en formato string en slice de int
func convertStringVector(string_vector string) []int {
	var vector []int
	aux := strings.Fields(string_vector)

	for i := 0; i < 3; i++ {
		number, err := strconv.Atoi(aux[i])
		if err != nil {
			fmt.Println(err)
		}
		vector = append(vector, number)
	}
	return vector
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
	//conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	conn, err = grpc.Dial("10.6.43.104:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewBrokerServiceClient(conn)

	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello World!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Reponse: %s", response.Body)

	comando := ""
	planeta := ""
	ciudad := ""
	nuevo_valor := ""

	last_fulcrum := "" // ultimo servidor solicitado
	fmt.Println(last_fulcrum)

	planet_dict := make(map[string]string) // key: name of planet, value: clock vector in format "0 0 0"
	//planet_dict["planetTest"] = "0 0 0"

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

		//fmt.Println("El comando que se enviará es:")
		//message:= strings.Fields(rpta_final)

		fmt.Println(rpta_final)
		fmt.Println("Enviando...")

		// agregar planeta a dictionario
		if _, ok := planet_dict[planeta]; ok {
			test := convertStringVector(planet_dict[planeta])
			fmt.Println(test)
		} else {
			planet_dict[planeta] = ""
		}

		// ENVIAR AL BROKER
		fulcrum_IP := ""

		switch comando {
		case "AddCity":
			response, err := c.AddCity(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: nuevo_valor, Vector: planet_dict[planeta], Last: last_fulcrum})
			if err != nil {
				log.Fatalf("Error when calling AddCity: %s", err)
			}
			log.Printf("Reponse: %s", response.IP)
			fulcrum_IP = response.IP

		case "UpdateName":
			response, err := c.UpdateName(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: nuevo_valor, Vector: planet_dict[planeta], Last: last_fulcrum})
			if err != nil {
				log.Fatalf("Error when calling UpdateName: %s", err)
			}
			log.Printf("Reponse: %s", response.IP)
			fulcrum_IP = response.IP

		case "UpdateNumber":
			response, err := c.UpdateNumber(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: nuevo_valor, Vector: planet_dict[planeta], Last: last_fulcrum})
			if err != nil {
				log.Fatalf("Error when calling UpdateNumber: %s", err)
			}
			log.Printf("Reponse: %s", response.IP)
			fulcrum_IP = response.IP

		case "DeleteCity":
			response, err := c.DeleteCity(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: "", Vector: planet_dict[planeta], Last: last_fulcrum})
			if err != nil {
				log.Fatalf("Error when calling DeleteCity: %s", err)
			}
			log.Printf("Reponse: %s", response.IP)
			fulcrum_IP = response.IP

		}

		// ENVIAR AL SERVIDOR QUE EL BROKER ESCOGIO
		// LA RPTA DEL SERVIDOR ES EL RELOJ DE VECTOR

		// conectar con Fulcrum X
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fulcrum_IP, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c_fulcrum := pb.NewFulcrumServiceClient(conn)

		response_vector := ""

		switch comando {
		case "AddCity":
			response, err := c_fulcrum.AddCity(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: nuevo_valor, Vector: planet_dict[planeta]})
			if err != nil {
				log.Fatalf("Error when calling AddCity: %s", err)
			}
			log.Printf("Reponse: %s", response.Vector)
			response_vector = response.Vector

		case "UpdateName":
			response, err := c_fulcrum.UpdateName(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: nuevo_valor, Vector: planet_dict[planeta]})
			if err != nil {
				log.Fatalf("Error when calling UpdateName: %s", err)
			}
			log.Printf("Reponse: %s", response.Vector)
			response_vector = response.Vector

		case "UpdateNumber":
			response, err := c_fulcrum.UpdateNumber(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: nuevo_valor, Vector: planet_dict[planeta]})
			if err != nil {
				log.Fatalf("Error when calling UpdateNumber: %s", err)
			}
			log.Printf("Reponse: %s", response.Vector)
			response_vector = response.Vector

		case "DeleteCity":
			response, err := c_fulcrum.DeleteCity(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: "", Vector: planet_dict[planeta]})
			if err != nil {
				log.Fatalf("Error when calling DeleteCity: %s", err)
			}
			log.Printf("Reponse: %s", response.Vector)
			response_vector = response.Vector

		}

		// guardar vector
		planet_dict[planeta] = response_vector
		fmt.Println(convertStringVector(planet_dict[planeta]))

		last_fulcrum = fulcrum_IP

		
	}
}
