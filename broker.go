package main

import (
	"bufio"
	"fmt"

	"math/rand"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/alvarofuentesm/SD-T3/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	pick := ""
	// en realidad siempre se ejecuta ese if
	if (in.Vector == ""){ // primera consulta, se busca aleatorio
		// Elegir aleatoriamente uno de los tres servidores
		randomIndex := rand.Intn(len(server_list))
		pick = server_list[randomIndex]
		fmt.Println("El servidor elegido para la consulta es:", pick)
		fmt.Println(pick)
	}else{
		pick := ""
		// al last_fulcrum o al de mayores cambios
		last_fulcrum:= in.Last // IP
		vector := convertStringVector(in.Vector) // (1,3,0)

		//consulto a servidor last_fulcrum su vectorFulcrum
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(last_fulcrum, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c_fulcrum := pb.NewFulcrumServiceClient(conn)

		response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
		if err != nil {
			log.Fatalf("Error when calling GetClockVector: %s", err)
		}
		log.Println(response.Vector)

		// buscar posicion del fulcrum
		position := 0
		for i := 0; i < 3; i++ {
			if (server_list[i] == last_fulcrum){
				position = i
			}
		}

		flag := false
		vectorFulcrum := make([]int, 3)
		if (response.Vector == ""){ // vacio, no existe el planeta en el servidor
			flag = true
		} else{
			vectorFulcrum = convertStringVector(response.Vector)
		}

		//comparar el vectorFulcrum con vector de cliente
		if ( flag || vector[position] < vectorFulcrum[position] ){
			vector_list := make([]string, 3)

			for i := 0; i < 3; i++ {
				var conn *grpc.ClientConn
				conn, err := grpc.Dial(server_list[i], grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %s", err)
				}
				defer conn.Close()

				c_fulcrum := pb.NewFulcrumServiceClient(conn)
				
				response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
				if err != nil {
					log.Fatalf("Error when calling GetClockVector: %s", err)
				}
				vector_list[i] = response.Vector
			}

			// asignar al que tenga el mayor valor de su vector
			max_value := 0 
			max_position:= 0
			for i := 0; i < 3; i++ {

				log.Println(vector_list)
				log.Println(vector_list[i])

				if (vector_list[i] == "") {
					continue
				}
				
				aux := convertStringVector(vector_list[i])[position]
				if (aux >= max_value) {
					max_value = aux
					max_position = i
				}
			}
			pick = server_list[max_position]

		}else{
			pick = last_fulcrum
		}

		return &pb.RespuestaBroker{IP: pick, Vector: "", Valor: ""}, nil
	}
	
	return &pb.RespuestaBroker{IP: pick, Vector: "", Valor: ""}, nil
}

func (s *Server) UpdateName(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	pick := ""
	// al last_fulcrum o al de mayores cambios
	last_fulcrum:= in.Last // IP
	vector := convertStringVector(in.Vector) // (1,3,0)

	//consulto a servidor last_fulcrum su vectorFulcrum
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(last_fulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_fulcrum := pb.NewFulcrumServiceClient(conn)

	response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
	if err != nil {
		log.Fatalf("Error when calling GetClockVector: %s", err)
	}
	log.Println(response.Vector)

	// buscar posicion del fulcrum
	position := 0
	for i := 0; i < 3; i++ {
		if (server_list[i] == last_fulcrum){
			position = i
		}
	}

	flag := false
	vectorFulcrum := make([]int, 3)
	if (response.Vector == ""){ // vacio, no existe el planeta en el servidor
		flag = true
	} else{
		vectorFulcrum = convertStringVector(response.Vector)
	}

	//comparar el vectorFulcrum con vector de cliente
	if ( flag || vector[position] < vectorFulcrum[position] ){
		vector_list := make([]string, 3)

		for i := 0; i < 3; i++ {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(server_list[i], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()

			c_fulcrum := pb.NewFulcrumServiceClient(conn)
			
			response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
			if err != nil {
				log.Fatalf("Error when calling GetClockVector: %s", err)
			}
			vector_list[i] = response.Vector
		}

		// asignar al que tenga el mayor valor de su vector
		max_value := 0 
		max_position:= 0
		for i := 0; i < 3; i++ {

			log.Println(vector_list)
			log.Println(vector_list[i])

			if (vector_list[i] == "") {
				continue
			}
			
			aux := convertStringVector(vector_list[i])[position]
			if (aux >= max_value) {
				max_value = aux
				max_position = i
			}
		}
		pick = server_list[max_position]

	}else{
		pick = last_fulcrum
	}
	return &pb.RespuestaBroker{IP: pick, Vector: "", Valor: ""}, nil
}

func (s *Server) UpdateNumber(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	pick := ""
	// al last_fulcrum o al de mayores cambios
	last_fulcrum:= in.Last // IP
	vector := convertStringVector(in.Vector) // (1,3,0)

	//consulto a servidor last_fulcrum su vectorFulcrum
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(last_fulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_fulcrum := pb.NewFulcrumServiceClient(conn)

	response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
	if err != nil {
		log.Fatalf("Error when calling GetClockVector: %s", err)
	}
	log.Println(response.Vector)

	// buscar posicion del fulcrum
	position := 0
	for i := 0; i < 3; i++ {
		if (server_list[i] == last_fulcrum){
			position = i
		}
	}

	flag := false
	vectorFulcrum := make([]int, 3)
	if (response.Vector == ""){ // vacio, no existe el planeta en el servidor
		flag = true
	} else{
		vectorFulcrum = convertStringVector(response.Vector)
	}

	//comparar el vectorFulcrum con vector de cliente
	if ( flag || vector[position] < vectorFulcrum[position] ){
		vector_list := make([]string, 3)

		for i := 0; i < 3; i++ {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(server_list[i], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()

			c_fulcrum := pb.NewFulcrumServiceClient(conn)
			
			response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
			if err != nil {
				log.Fatalf("Error when calling GetClockVector: %s", err)
			}
			vector_list[i] = response.Vector
		}

		// asignar al que tenga el mayor valor de su vector
		max_value := 0 
		max_position:= 0
		for i := 0; i < 3; i++ {

			log.Println(vector_list)
			log.Println(vector_list[i])

			if (vector_list[i] == "") {
				continue
			}
			
			aux := convertStringVector(vector_list[i])[position]
			if (aux >= max_value) {
				max_value = aux
				max_position = i
			}
		}
		pick = server_list[max_position]

	}else{
		pick = last_fulcrum
	}

	return &pb.RespuestaBroker{IP: pick, Vector: "", Valor: ""}, nil
}

func (s *Server) DeleteCity(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	pick := ""
	// al last_fulcrum o al de mayores cambios
	last_fulcrum:= in.Last // IP
	vector := convertStringVector(in.Vector) // (1,3,0)

	//consulto a servidor last_fulcrum su vectorFulcrum
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(last_fulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_fulcrum := pb.NewFulcrumServiceClient(conn)

	response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
	if err != nil {
		log.Fatalf("Error when calling GetClockVector: %s", err)
	}
	log.Println(response.Vector)

	// buscar posicion del fulcrum
	position := 0
	for i := 0; i < 3; i++ {
		if (server_list[i] == last_fulcrum){
			position = i
		}
	}

	flag := false
	vectorFulcrum := make([]int, 3)
	if (response.Vector == ""){ // vacio, no existe el planeta en el servidor
		flag = true
	} else{
		vectorFulcrum = convertStringVector(response.Vector)
	}

	//comparar el vectorFulcrum con vector de cliente
	if ( flag || vector[position] < vectorFulcrum[position] ){
		vector_list := make([]string, 3)

		for i := 0; i < 3; i++ {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(server_list[i], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()

			c_fulcrum := pb.NewFulcrumServiceClient(conn)
			
			response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
			if err != nil {
				log.Fatalf("Error when calling GetClockVector: %s", err)
			}
			vector_list[i] = response.Vector
		}

		// asignar al que tenga el mayor valor de su vector
		max_value := 0 
		max_position:= 0
		for i := 0; i < 3; i++ {

			//log.Println(vector_list)
			//log.Println(vector_list[i])

			if (vector_list[i] == "") {
				continue
			}
			
			aux := convertStringVector(vector_list[i])[position]
			if (aux >= max_value) {
				max_value = aux
				max_position = i
			}
		}
		pick = server_list[max_position]

	}else{
		pick = last_fulcrum
	}

	return &pb.RespuestaBroker{IP: pick, Vector: "", Valor: ""}, nil	
	//return &pb.RespuestaBroker{IP: "", Vector: "", Valor: ""}, nil
}

func (s *Server) GetNumberRebelds(ctx context.Context, in *pb.Comando) (*pb.RespuestaBroker, error) {
	pick := ""
	if (in.Vector == ""){ 
		vector_list := make([]string, 3)

		for i := 0; i < 3; i++ {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(server_list[i], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()

			c_fulcrum := pb.NewFulcrumServiceClient(conn)
			
			response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
			if err != nil {
				log.Fatalf("Error when calling GetClockVector: %s", err)
			}
			vector_list[i] = response.Vector
		}

		max_position:= 0
		for i := 0; i < 3; i++ {
			if (vector_list[i] == "") {
				continue
			}else{
				max_position = i // el primero que encuentra
			}
		}
		pick = server_list[max_position]
	}else{		
		
		// al last_fulcrum o al de mayores cambios
		last_fulcrum:= in.Last // IP
		vector := convertStringVector(in.Vector) // (1,3,0)

		//consulto a servidor last_fulcrum su vectorFulcrum
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(last_fulcrum, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c_fulcrum := pb.NewFulcrumServiceClient(conn)

		response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
		if err != nil {
			log.Fatalf("Error when calling GetClockVector: %s", err)
		}
		log.Println(response.Vector)

		// buscar posicion del fulcrum
		position := 0
		for i := 0; i < 3; i++ {
			if (server_list[i] == last_fulcrum){
				position = i
			}
		}

		flag := false
		vectorFulcrum := make([]int, 3)
		if (response.Vector == ""){ // vacio, no existe el planeta en el servidor
			flag = true
		} else{
			vectorFulcrum = convertStringVector(response.Vector)
		}

		//comparar el vectorFulcrum con vector de cliente
		if ( flag || vector[position] < vectorFulcrum[position] ){
			vector_list := make([]string, 3)

			for i := 0; i < 3; i++ {
				var conn *grpc.ClientConn
				conn, err := grpc.Dial(server_list[i], grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %s", err)
				}
				defer conn.Close()

				c_fulcrum := pb.NewFulcrumServiceClient(conn)
				
				response, err := c_fulcrum.GetClockVector(context.Background(), &pb.Message{Body: in.Planeta})
				if err != nil {
					log.Fatalf("Error when calling GetClockVector: %s", err)
				}
				vector_list[i] = response.Vector
			}

			// asignar al que tenga el mayor valor de su vector
			max_value := 0 
			max_position:= 0
			for i := 0; i < 3; i++ {

				log.Println(vector_list)
				log.Println(vector_list[i])

				if (vector_list[i] == "") {
					continue
				}
				
				aux := convertStringVector(vector_list[i])[position]
				if (aux >= max_value) {
					max_value = aux
					max_position = i
				}
			}
			pick = server_list[max_position]

		}else{
			pick = last_fulcrum
		}
	}

	conn, err := grpc.Dial(pick, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_fulcrum_ := pb.NewFulcrumServiceClient(conn)

	response, err := c_fulcrum_.GetNumberRebelds(context.Background(), &pb.Comando{Planeta: in.Planeta, Ciudad: in.Ciudad, Valor: in.Valor, Vector: in.Vector})
	if err != nil {
		log.Fatalf("Error when calling GetNumberRebelds: %s", err)
	}
	log.Println(response.Vector)

	return &pb.RespuestaBroker{IP: pick, Vector: response.Vector, Valor: response.Valor}, nil 
	
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

func startServer() {
	/*  Iniciar servidor Lider */
	fmt.Println("Iniciando servidor Broker...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterBrokerServiceServer(grpcServer, &s)
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

// para convertir un vector de int en string con formato 0 0 0
func convertIntVector(string_vector []int) string {
	var vector string
	vector = strconv.Itoa(string_vector[0]) + " " + strconv.Itoa(string_vector[1]) + " " + strconv.Itoa(string_vector[2])
	return vector
}

var server_list = [3]string{":8000", ":8100", ":8200"}
//var server_list = [3]string{":8000", ":8000", ":8000"}

func main() {
	//var listening = true
	fmt.Println("Bienvenido al broker Mos Eisley")

	go startServer()

	for {
	}

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

}
