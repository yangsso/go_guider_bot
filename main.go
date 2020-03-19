package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "211.251.238.39:500")
	if nil != err{
		fmt.Println(err)
	}

	//addr 와 통신을 시도.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		fmt.Println(err)
	}

	cmd := []byte("run:GUIDER top -J -a -e dn")
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