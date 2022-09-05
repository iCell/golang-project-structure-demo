package main

func main() {
	server := InitializeServer()
	server.Run(":4444")
}
