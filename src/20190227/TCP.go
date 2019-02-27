package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn){
	data := make([]byte, 4096) //4096 크기의 바이트 생성

	for {
		n, err := c.Read(data) //클라이언트에서 받은 데이터를 읽음.
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) //읽은 데이터 출력

		_, err = c.Write(data[:n]) //클라이언트로 데이터를 보냄.
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main(){
	ln, err := net.Listen("tcp", ":8000") //tcp 프로토콜에 8000포트로 연결을 받는다.
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close() //main함수 끝나기 직전에 연결대기닫음.

	for {
		conn, err := ln.Accept() //클라이언트가 연결되면 TCP 연결 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		defer conn.Close() //main함수끝나기 직전에 TCP연결닫음.
		
		go requestHandler(conn) //패킷처리할 함수를 고루틴으로 실행.
	}
}