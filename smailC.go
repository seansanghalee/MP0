package main

// This is the source code for Sean's Mail client using TCP

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type Email struct {
	to, from, title, content, date string
} // creates a type Email that contains necessary information to be sent over

func writeEmail() Email {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("To : ")
	to, _ := reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("From : ")
	from, _ := reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Title : ")
	title, _ := reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Content : ")
	content, _ := reader.ReadString('\n')

	date := time.Now().Format("2006-01-02 15:04:05")

	email := Email{to, from, title, content, date}

	return email

} // This function asks user input to create a struct Email

func emailToString(e Email) string {

	return e.to + e.from + e.title + e.content + e.date + "|"

} // This function converts information in Email to a single string

func main () {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	} // Checks if host address and port # are provided

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	} // Connects to the server via TCP

	email := writeEmail() // creates a struct Email
	emailToSend := emailToString(email) // converts Email to a string so it can be sent
	fmt.Fprint(c, emailToSend)

	for {
		stop, _ := bufio.NewReader(c).ReadString('\n')
		if stop == "STOP\n" {
			fmt.Println("Email received by the server")
			fmt.Println("Exiting Sean's Mail Client!")
			return
		}
	}

}