package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var channel1 = &sync.Map{}
var channel2 = &sync.Map{}
var connMapAll = &sync.Map{}

func main() {
	var protocolo = "tcp"
	var puerto = "4040"
	ln, err := net.Listen(protocolo, ":"+puerto)
	if err != nil {

	}
	log.Println("**** inicio servidor ***")
	log.Printf("servicio inicio en : (%s) %s\n", protocolo, puerto)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			if err := conn.Close(); err != nil {
				fmt.Println("fallo al cerrar la coneccion", err)
			}
			continue
		}

		id := uuid.New().String()
		connMapAll.Store(id, conn)
		go manejarConeciones(id, conn, connMapAll)
	}
}
func manejarConeciones(id string, conn net.Conn, connMap *sync.Map) {

	defer conn.Close()

	appendBytes := func(dest, src []byte) ([]byte, error) {
		for _, b := range src {
			if b == '\n' {
				return dest, io.EOF
			}
			dest = append(dest, b)
		}
		return dest, nil
	}

	for {
		var cmdLine []byte

		for {
			chunk := make([]byte, 4)
			n, err := conn.Read(chunk)
			if err != nil {
				if err == io.EOF {
					cmdLine, _ = appendBytes(cmdLine, chunk[:n])
					break
				}
				log.Println("connection read error:", err)
				return
			}
			if cmdLine, err = appendBytes(cmdLine, chunk[:n]); err == io.EOF {
				break
			}
		}

		message := string(cmdLine)
		fmt.Println("linea de comando: " + message + ",")
		cmd, param1, param2 := parseCommand(message)
		fmt.Println("cmd:" + cmd)
		fmt.Println("param1:" + param1)
		fmt.Println("param2:" + param2)
		switch strings.ToLower(cmd) {
		case "receive":
			unirceCanal(conn, id, param1)
		case "send":
			enviarCanal(param1, param2, conn)
		default:
			if _, error := conn.Write([]byte("comando invalido \n")); error != nil {
				fmt.Println("comando invalido \n", error)
				return
			}
		}
	}
}

func mensajeBienbenida(conn net.Conn, id string) {
	if _, err := conn.Write([]byte("Conectado. Bienvenido ...: " + id + "\n")); err != nil {
		log.Println("error writing:", err)
		return
	}
	fmt.Println("mensaje de bienvenida enviado")
}
func parseCommand(cmdLine string) (cmd, param1 string, param2 string) {
	parts := strings.Split(cmdLine, " ")

	cmd = strings.TrimSpace(parts[0])
	param1 = strings.TrimSpace(parts[1])
	if len(parts) < 3 {
		param2 = ""
	} else {
		param2 = strings.TrimSpace(parts[2])
	}

	return
}
func unirceCanal(conn net.Conn, id string, canal string) {
	if canal == "1" {
		channel2.Delete(id)
		channel1.Store(id, conn)
	}
	if canal == "2" {
		channel1.Delete(id)
		channel2.Store(id, conn)
	}
}
func enviarCanal(canal string, archivo string, conn net.Conn) {
	fmt.Println("enviar canal" + canal)
	archivo = "servidor/" + archivo
	canalEnviar := &sync.Map{}
	if canal == "1" {
		canalEnviar = channel1
	}
	if canal == "2" {
		canalEnviar = channel2
	}
	if file, erroFile := os.ReadFile("servidor/file1.txt"); erroFile == nil {
		canalEnviar.Range((func(key, value interface{}) bool {
			log.Println("uuid:" + key.(string))
			if conn, ok := value.(net.Conn); ok {

				if rpt, err := conn.Write(file); err != nil {
					log.Fatal("error on writing to connection")
					log.Fatal(err.Error())
				} else {
					conn.Write([]byte("\n"))
					log.Println("archivo enviado: " + archivo + string(rpt))
				}
			}
			return true
		}))
	} else {

		log.Fatalln("no se puede leer archivo: ")
		log.Println(erroFile.Error())
	}
}
