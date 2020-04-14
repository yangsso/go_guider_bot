package main

import (
	"fmt"
	"net"
	"os"
)

type NetworkInfo struct {
	Port	string
	Address	string
	Command	string
}

//private class method
func getTcpAddress(info *NetworkInfo) string {
	return fmt.Sprintf("%s:%s", info.Address , info.Port)
}

func runCommand(conn net.Conn, command string) {
	cmd := []byte(command)
	conn.Write(cmd)

	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			fmt.Println("READ ERROR "+ err.Error())
		}

		if 0 < n {
			data := recvBuf[:n]
			fmt.Println(string(data))
		} else {
			conn.Close()
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("plz put argument")
	}

	port := os.Args[1]
	address := os.Args[2]
	command := os.Args[3]
	//"run:GUIDER top -J -a -e dn"
	networkInfo := new(NetworkInfo)
	networkInfo.Port = port
	networkInfo.Address = address
	networkInfo.Command = command

	//addr 와 통신을 시도.
	conn, err := net.Dial("tcp", getTcpAddress(networkInfo))
	if nil != err {
		fmt.Println(err)
	}

	runCommand(conn, networkInfo.Command)
}