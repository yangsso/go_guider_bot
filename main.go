package main

import (
	"fmt"
	"net"
	"os"
	//"github.com/nlopes/slack"
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

func getRunGuiderCommand(info *NetworkInfo) string {
	return fmt.Sprintf("run:%s",info.Command)
}

func tcpConnectByNetInfo(info *NetworkInfo) net.Conn {
	//addr 와 통신을 시도.
	conn, err := net.Dial("tcp", getTcpAddress(info))
	if nil != err {
		panic(err)
	}
	return conn
}

func runCommand(conn net.Conn, info *NetworkInfo) {
	cmd := []byte(getRunGuiderCommand(info))
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

	conn := tcpConnectByNetInfo(networkInfo)
	runCommand(conn, networkInfo)
}