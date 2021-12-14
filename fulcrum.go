package main

import (
	//"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	pb "github.com/alvarofuentesm/SD-T3/proto"
	"log"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
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
	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) UpdateName(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) UpdateNumber(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) DeleteCity(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

func (s *Server) GetNumberRebelds(ctx context.Context, in *pb.Comando) (*pb.RespuestaReplica, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaReplica{Vector: "0 0 0", Valor: ""}, nil
}

// para que broker consulte los vectores
func (s *Server) GetClockVector(ctx context.Context, in *pb.Message) (*pb.RespuestaReplica, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.RespuestaReplica{Vector: "0 0 0"}, nil
}

/*************************************************************************************/
var isLocal = false // proceso corriendo localmente (no en VM's)
var isCoordinator = false // proceso es coordinador 

func startServer(){

	argsWithoutProg := os.Args[1:]

	
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "L" {
		isLocal =  true
	}

	if len(argsWithoutProg) > 1 && argsWithoutProg[1] == "x" {
		isCoordinator =  true
	}

	log.Println(isLocal, isCoordinator)



	/*  Iniciar servidor Fulcrum */
	fmt.Println("Iniciando servidor Fulcrum...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
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

var planet_dict = make(map[string]string) // key: name of planet, value: clock vector in format "0 0 0"

func main() {

	go startServer()

	for { 
	}
}
