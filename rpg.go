package main 

import "fmt"

func main() {
	fmt.Println("Dare ye enter?")
	options := true
	lobbyRoom(options)
}

func lobbyRoom (options bool) {
	fmt.Println("You enter a creepy lobby")
}
