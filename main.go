package main

import (
	"fmt"
	"net"
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
		}
	}
}

func main() {
	networkInfo := new(NetworkInfo)
	networkInfo.Port = "500"
	networkInfo.Address = "211.251.238.39"
	networkInfo.Command = "run:GUIDER top -J -a -e dn"

	//addr 와 통신을 시도.
	conn, err := net.Dial("tcp", getTcpAddress(networkInfo))
	if nil != err {
		fmt.Println(err)
	}

	runCommand(conn, networkInfo.Command)
}