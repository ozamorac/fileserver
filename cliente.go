package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	modo := arguments[1]
	canal := arguments[2]
	protocolo := "tcp"
	servidor := "127.0.0.1:4040"
	fmt.Println(len(arguments))
	if len(arguments) > 5 || len(arguments) == 1 {
		fmt.Println("error de comandos")
		return
	}
	conn, err := net.Dial(protocolo, servidor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("canal selecionado" + canal)

	comandoInicial(conn, arguments)

	for {
		switch strings.Trim(strings.ToLower(modo), " ") {
		case "receive":
			if message, err := bufio.NewReader(conn).ReadBytes('\n'); err == nil {
				id := arguments[3]
				fmt.Println("archivo recibido")
				os.WriteFile("cliente/file"+id+".txt", message, 0644)
			} else {
				fmt.Print(err.Error())
				conn.Close()
				return
			}
		case "send":
			if message, err := bufio.NewReader(conn).ReadString('\n'); err == nil {
				fmt.Println("Message from server para send: " + message)
				if message == "cerrar" {
					fmt.Print(err.Error())
					conn.Close()
					return
				}
			} else {
				fmt.Print(err.Error())
				conn.Close()
				return
			}
		default:
			log.Println("escriba los comandos de manera correcta.\n cerrando conecccion..\n")
			conn.Close()
			return
		}

	}
}
func comandoInicial(conn net.Conn, arguments []string) {
	modo := arguments[1]
	canal := arguments[2]
	comandoEnviado := ""
	switch strings.Trim(strings.ToLower(modo), " ") {
	case "receive":
		comandoEnviado = modo + " " + canal + "\n"
		fmt.Println(comandoEnviado)
		if _, err := conn.Write([]byte(comandoEnviado)); err != nil {
			log.Fatal("error on writing to connection")
		}
	case "send":
		log.Println("send opccion")
		file := arguments[3]
		comandoEnviado = modo + " " + canal + " " + file + "\n"
		fmt.Println(comandoEnviado)
		if _, err := conn.Write([]byte(comandoEnviado)); err != nil {
			log.Fatal("error on writing to connection")
		}
	default:
		log.Println("escriba los comandos de manera correcta.\n cerrando conecccion..\n")
		conn.Close()
		return
	}
}
