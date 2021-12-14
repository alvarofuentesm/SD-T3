package main

import (
	//"bufio"
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/alvarofuentesm/SD-T3/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

/********************************** gRPC **********************************************/
type Server struct {
	pb.UnimplementedFulcrumServiceServer
}

// funcion response hello para debug
func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) AddCity(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)

	// ---------- AÑADIR CIUDAD ----------
	fmt.Println("Se añadirá una ciudad")
	planeta := in.GetPlaneta()
	ciudad := in.GetCiudad()
	cant_rebels := in.GetValor()

	// Asumir que el archivo existe y hacer append
	file, err := os.OpenFile(planeta+".txt", os.O_APPEND, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			// Archivo no existe, crear uno
			file, err = os.OpenFile(planeta+".txt", os.O_CREATE, 0666)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	// Escribir los datos de la peticion al txt
	str_to_write := planeta + " " + ciudad + " " + cant_rebels
	bufferedWriter := bufio.NewWriter(file)
	_, err = bufferedWriter.WriteString(str_to_write + "\n")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) UpdateName(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)

	// ---------- RENAME CIUDAD ----------
	fmt.Println("Se añadirá una ciudad")
	planeta := in.GetPlaneta()
	ciudad := in.GetCiudad()
	nueva_ciudad := in.GetValor()

	// Chequear si el planeta de la ciudad existe
	input, err := ioutil.ReadFile(planeta + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			// si no existe, error (no deberia pasar porque el broker evita esto, pero porsiacaso)
			fmt.Println("La ciudad no existe en este servidor")
			log.Fatal(err)
		}
	}

	// trozo de codigo sacado de https://www.socketloop.com/tutorials/golang-read-a-text-file-and-replace-certain-words

	// Si existe, leerlo y reemplazar la ciudad
	output := bytes.Replace(input, []byte(ciudad), []byte(nueva_ciudad), -1)

	// Borrar el archivo planeta.txt
	err = os.Remove(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// Crear un nuevo archivo planeta.txt con la info cambiada
	if err = ioutil.WriteFile(planeta+".txt", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// basicamente "se acabo el mundo pero inmediatamente después empezó otro mundo (casi) igual"

	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) UpdateNumber(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)

	// ---------- RENAME NUMERO ----------
	fmt.Println("Se updateara un valor")
	planeta := in.GetPlaneta()
	ciudad := in.GetCiudad()
	nuevo_valor := in.GetValor()
	new_string := ciudad + " " + nuevo_valor

	// Abrir un archivo si existe
	data, err := ioutil.ReadFile(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// Buscar en todo el string data la ciudad a updatear, con una regex
	m1 := regexp.MustCompile(`(` + ciudad + ` )[0-9]*`)       // Captura x ej "ciudad1 88" para cambiarlo por "ciudad1 nuevo_valor"
	new_data := m1.ReplaceAllString(string(data), new_string) // Aplicandole el regex

	// Borrar el archivo planeta.txt
	err = os.Remove(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// Crear un nuevo archivo planeta.txt con la info cambiada
	if err = ioutil.WriteFile(planeta+".txt", []byte(new_data), 0666); err != nil { // FIJARSE SI ESTO SIRVE Y TE RESPETA LOS newline
		fmt.Println(err)
		os.Exit(1)
	}

	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) DeleteCity(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)

	// ---------- ELIMINAR CIUDAD ----------
	fmt.Println("Se eliminara una ciudad")
	planeta := in.GetPlaneta()
	ciudad := in.GetCiudad()

	// Abrir un archivo si existe
	data, err := ioutil.ReadFile(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// Buscar en todo el string data la ciudad a updatear, con una regex
	m1 := regexp.MustCompile(`(` + planeta + `) (` + ciudad + `) [0-9]*\n`) // Captura toda la linea donde se encuentra la ciudad
	new_data := m1.ReplaceAllString(string(data), "")                       // Aplicandole el regex y al cambiarlo por "" lo elimino

	// Borrar el archivo planeta.txt
	err = os.Remove(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// Crear un nuevo archivo planeta.txt con la info cambiada
	if err = ioutil.WriteFile(planeta+".txt", []byte(new_data), 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) GetNumberRebelds(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)

	// ---------- CONSULTAR DATOS ----------
	fmt.Println("Se consultara un dato")
	planeta := in.GetPlaneta()
	ciudad := in.GetCiudad()

	// Abrir el archivo
	file, err := os.Open(planeta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // se cierra luego
	rpta := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Recorro cada linea
		line := scanner.Text()
		line_slice := strings.Split(line, " ")
		if ciudad == line_slice[1] {
			rpta = line_slice[2]
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Enviarle la rpta al broker y que se lo pase a Leia
	fmt.Println("La respuesta es:", rpta)

	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

// para que broker consulte los vectores
func (s *Server) GetClockVector(ctx context.Context, in *pb.Message) (*pb.RespuestaReplica, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaReplica{Vector: "0 0 0"}, nil
}

/*************************************************************************************/
var isLocal = false       // proceso corriendo localmente (no en VM's)
var isCoordinator = false // proceso es coordinador

func startServer() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "L" {
		isLocal = true
	}

	if len(argsWithoutProg) > 1 && argsWithoutProg[1] == "x" {
		isCoordinator = true
	}

	log.Println(isLocal, isCoordinator)

	/*  Iniciar servidor Fulcrum */
	fmt.Println("Iniciando servidor Fulcrum...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterFulcrumServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
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

var planet_dict = make(map[string]string) // key: name of planet, value: clock vector in format "0 0 0"

func main() {

	go startServer()

	for {
	}
}
