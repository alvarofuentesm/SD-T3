hello:
	echo "Hello"

leia:
	go run leia.go

info1: # Ahsoka 
	go run informante.go ahsoka

info2: # Thrawn 
	go run informante.go thrawn

fulcrumX: # para las VM sera go run fulcrum.go VM x, para la entrega cambiar eso
	go run fulcrum.go VM x

fulcrumY: # para las VM sera go run fulcrum.go VM x, para la entrega cambiar eso
	go run fulcrum.go VM y

fulcrumZ: # para las VM sera go run fulcrum.go VM x, para la entrega cambiar eso
	go run fulcrum.go VM z

broker:
	go run broker.go

clean:
	rm *.txt