package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn){
	data := make([]byte, 4096);

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		_, err = c.Write(data[:n])
		
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main1(){
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}
