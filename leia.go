package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

// para convertir un vector en formato string en slice de int
func convertStringVector(string_vector string) []int{
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
	var playing = true
	fmt.Println("Bienvenida Leia Organa")

	// conexcion broker
	
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.104:9000", grpc.WithInsecure())
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
	

	planeta := ""
	ciudad := ""

	last_fulcrum := "" // ultimo servidor solicitado
	fmt.Println(last_fulcrum)

	planet_dict := make(map[string]string) // key: name of planet, value: clock vector in format "0 0 0"

	for playing {
		fmt.Println("Proceda a armar su consulta:")

		// ---------- ELEGIR PLANETA ----------
		fmt.Println("Introduzca el nombre del planeta a consultar:")
		fmt.Scanln(&planeta)

		// ---------- ELEGIR CIUDAD ----------
		fmt.Println("Introduzca el nombre de la ciudad a consultar:")
		fmt.Scanln(&ciudad)

		// agregar planeta a dictionario
		if _, ok := planet_dict[planeta]; ok {
			test := convertStringVector(planet_dict[planeta])
			fmt.Println(test)
		}else{
			planet_dict[planeta] = ""
		}

		rpta_final := "GetNumberRebels " + planeta + " " + ciudad
		fmt.Println("El comando que se enviará es:")
		fmt.Println(rpta_final)
		fmt.Println("Enviando...")

		response, err := c.GetNumberRebelds(context.Background(), &pb.Comando{Planeta: planeta, Ciudad: ciudad, Valor: "", Vector: planet_dict[planeta], Last: last_fulcrum})
		if err != nil {
			log.Fatalf("Error when calling GetNumberRebelds: %s", err)
		}
		log.Printf("Reponse: %s", response.IP)
		//log.Printf("Reponse: ", response)

		// guardar vector
		planet_dict[planeta] = response.Vector
		fmt.Println(convertStringVector(planet_dict[planeta]))

		last_fulcrum =  response.IP

		fmt.Println("El número de rebeldes es: ", response.Valor)

		

		// ---------- SEGUIR EJECUTANDO ----------
		/*
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
		*/

	}
}
