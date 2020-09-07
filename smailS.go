package main

// This is the source code for Sean's Mail server using TCP

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func displayEmail(e string) {

	info := strings.Split(e, "\n")

	fmt.Println("To : " + info[0])
	fmt.Println("From : " + info[1])
	fmt.Println("Title : " + info[2])
	fmt.Println("Content : " + info[3])
	fmt.Println("Date : " + info[4][:len(info[4]) - 1])

	return

} // This function displays the info of email received fro the client in an appropriate format

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	} // Checks if port # is provided

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	} // Connects to the client via TCP
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("...")
	fmt.Println("Connected!")
	fmt.Println("")

	message, err := bufio.NewReader(c).ReadString('|')
	if err != nil {
		fmt.Print(err)
	} // reads email sent from the client

	displayEmail(message) // displays received email in an appropriate format
	ACK := "STOP\n"
	c.Write([]byte(ACK))

	fmt.Println("")
	fmt.Println("Exiting Sean's Mail Server!")

}