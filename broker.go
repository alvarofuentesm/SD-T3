package main

import (
	"bufio"
	"fmt"
	//"math/rand"
	"os"
	//"strings"
	pb "github.com/alvarofuentesm/SD-T3/proto"
	"log"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

/********************************** gRPC **********************************************/
type Server struct {
	pb.UnimplementedBrokerServiceServer
}

// funcion response hello para debug
func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	// conectar con Fulcrum X
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	//conn, err := grpc.Dial("X.X.XX.XXX:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_fulcrum := pb.NewFulcrumServiceClient(conn)
	
	response, err := c_fulcrum.SayHello(context.Background(), &pb.Message{Body: "TEST"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	return &pb.Message{Body: response.Body}, nil
}

func (s *Server) AddCity(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaBroker{IP: "", Vector: "0,0,0", Valor: ""}, nil
}

func (s *Server) UpdateName(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaBroker{IP: "", Vector: "0,0,0", Valor: ""}, nil
}

func (s *Server) UpdateNumber(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaBroker{IP: "", Vector: "0,0,0", Valor: ""}, nil
}

func (s *Server) DeleteCity(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaBroker{IP: "", Vector: "0,0,0", Valor: ""}, nil
}

func (s *Server) GetNumberRebelds(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaBroker{IP: "", Vector: "0,0,0", Valor: ""}, nil
}


/*************************************************************************************/
func leer_opt() rune { // Para leer numeros
	reader := bufio.NewReader(os.Stdin)
	opt, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	return opt
}

func startServer(){
	/*  Iniciar servidor Lider */
	fmt.Println("Iniciando servidor Broker...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterBrokerServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	//var listening = true
	fmt.Println("Bienvenido al broker Mos Eisley")

	go startServer()

	/*
	server_list := []int{1, 2, 3} // Deberian ir las direcciones de los server

	for listening {

		fmt.Println("Esperando peticion...")
		// ESPERAR PETICION...???
		// llega una
		peticion := ""

		// Elegir aleatoriamente uno de los tres servidores
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
	*/
	for { 
	}


	//
}
