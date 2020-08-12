package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	DefaultIp   = "127.0.0.1"
	DefaultPort = 6666
)

var address string

func main() {
	reader := bufio.NewReader(os.Stdin)
	address = fmt.Sprintf("%s:%d", DefaultIp, DefaultPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	clientCheckError(err)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	clientCheckError(err)
	defer conn.Close()

	for {
		printConsole()
		cmd, err := reader.ReadString('\n')
		clientCheckError(err)
		cmd = strings.Trim(cmd, "\n")

		_, err = sendCmdToServer(cmd, conn)
		clientCheckError(err)

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		clientCheckError(err)

		if n == 0 {
			printInfo("read error: read 0 bytes")
		} else if err == nil {
			printConsoleMessage(string(buff), n)
		} else {
			printInfo("server response err")
		}
	}
}

func sendCmdToServer(cmd string, conn *net.TCPConn) (n int, err error) {
	return conn.Write([]byte(cmd))
}

func clientCheckError(err error) {
	if err != nil {
		log.Printf("err: %s", err.Error())
		os.Exit(1)
	}
}

func printConsole() {
	fmt.Print(address + "> ")
}

func printConsoleMessage(msg string, num int) {
	messages := decodeMessage(msg, num)

	for _, message := range messages {
		fmt.Println(message)
	}
}

func printInfo(msg string) {
	fmt.Println(msg)
}
