# MP0
This is my take on MP0

**Note: I acknowledge that this solution is meant to work when the content of "email" does not contain any new lines.**

# To Run

To run the Sean's Mail (Smail) server and client, the user must first run the server (smailS.go).
```bash
go run smailS.go 1234
```
This command allows Smail server to listen to port #1234.

Then, the user needs to run the client (smailC) by typing in a new terminal window,
```bash
go run smailC.go 127.0.0.1:1234
```
The user will be prompted to provide the sender and receiver information, as well as the title and content of the email.

When done, the server should output the exact same email as the user input on the client.

#Structure and Design

The client (which we'll call it Process A) uses struct type Email to store five data fields: To, From, Date, Title, Content.
```bash
type Email struct {
	to, from, title, content, date string
} 
```
Process B, the server, waits for a specific port number and connects to it (the client) when found.

In this program, Process A converts type Email into a series of strings that can be sent through a TCP channel established with Process B. Then Process B reconstructs type Email and displays the necessary data in a readable format.

The output should look something like this.
```
go run smailS.go 1234
...
Connected!

You got a mail!

To : prof@bc.edu
From : me@bc.edu
Title : Hello, prof!
Content : Looking forward to our powerlifting sesh real soon when things are looking better :)
Date : 2020-09-07 22:42:53

Exiting Smail
```