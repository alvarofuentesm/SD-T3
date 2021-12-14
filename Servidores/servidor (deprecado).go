package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
	fmt.Println("Bienvenido a la inicializacion del servidor")
	fmt.Println("Escoja el número que le corresponda (1,2,3)")
	fmt.Println("dist113 -> 1")
	fmt.Println("dist114 -> 2")
	fmt.Println("dist115 -> 3")
	num_server := ""
	fmt.Scanln(&num_server)

	for listening {

		fmt.Println("Esperando peticion...")
		// ESPERAR PETICION...??? (o algo asi)
		// llega una, guardarla en "peticion"
		peticion := ""

		peticion_slice := strings.Split(peticion, " ")
		comando := peticion_slice[0]
		planeta := peticion_slice[1]
		ciudad := peticion_slice[2]

		switch comando {

		case "AddCity":
			// ---------- AÑADIR CIUDAD ----------
			fmt.Println("Se añadirá una ciudad")
			cant_rebels := peticion_slice[3]

			// Chequear si existe el archivo del planeta
			// Asumir que existe y hacer append
			file, err := os.OpenFile(planeta+".txt", os.O_APPEND, 0666)
			if err != nil {
				if os.IsNotExist(err) {
					// Archivo no existe, crear uno
					file, err = os.OpenFile(planeta+".txt", os.O_CREATE, 0666) // Ni idea por que el 0666 pero asi salia en el ejemplo xd
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

		case "UpdateName":
			// ---------- RENAME CIUDAD ----------
			fmt.Println("Se añadirá una ciudad")
			nueva_ciudad := peticion_slice[3]

			// Chequear si el planeta de la ciudad existe
			input, err := ioutil.ReadFile(planeta + ".txt")
			if err != nil {
				if os.IsNotExist(err) {
					// si no existe, error
					fmt.Println("La ciudad no existe en este servidor")
					// --------------------------------------------
					// No deberia pasar, añadirle logica al fulcrum para que te mande a donde se sabe que existe
					// --------------------------------------------
					break
				}
			}

			// trozo de codigo sacado de https://www.socketloop.com/tutorials/golang-read-a-text-file-and-replace-certain-words
			// no aseguro que funcione completamente xdd hay que probarlo

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

			// falta escribir en el log

		case "UpdateNumber":
			// ---------- RENAME NUMERO ----------
			fmt.Println("Se updateara un valor")
			nuevo_valor := peticion_slice[3]
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
		case "DeleteCity":
			// ---------- ELIMINAR CIUDAD ----------
			fmt.Println("Se eliminara una ciudad")

			// Abrir un archivo si existe
			data, err := ioutil.ReadFile(planeta + ".txt")
			if err != nil {
				log.Fatal(err)
			}

			// Buscar en todo el string data la ciudad a updatear, con una regex
			m1 := regexp.MustCompile(`(` + planeta + `) (` + ciudad + `) [0-9]*\n`) // Captura x ej "ciudad1 88" para cambiarlo por "ciudad1 nuevo_valor"
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

		case "GetNumberRebels":
			// ---------- CONSULTAR DATOS ----------
			fmt.Println("Se consultara un dato")

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

		}

		// FALTA AÑADIR AL LOG DE CADA PLANETA LAS PETICIONES QUE LLEGAN

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
}
