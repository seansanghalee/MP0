# MP0
This is my take on MP0

**Note: I acknowledge that this solution is meant to work when the content of "email" does not contain any new lines.**

# To Run

To run the Sean's Mail (Smail) server and client, the user must first run the server (smailS.go).
```bash
go run smailS.go 1234
```
This command allows Smail server to listen to a port #1234.

Then, the user needs to run the client (smailC) by typing in a new terminal window,
```bash
go run smailC.go 127.0.0.1:1234
```
The user will be prompted to provide the sender and receiver information, as well as the title and content of the email.

When done, the server should output the exact same email as the user input on the client.