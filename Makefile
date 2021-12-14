hello:
	echo "Hello"

leia:
	go run leia.go

info1: # Ahsoka 
	go run informante.go ahsoka

info2: # Thrawn 
	go run informante.go thrawn

fulcrum:
	go run fulcrum.go

broker:
	go run broker.go

clean:
	rm *.txt