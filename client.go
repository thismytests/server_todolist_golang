package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")

	if err != nil {
		fmt.Println("End")
		return
	}

	defer conn.Close()

	for {
		var source string
		fmt.Println("Enter the word")

		_, err := fmt.Scanln(&source)

		if err != nil {
			fmt.Println("err", err)
			continue
		}

		fmt.Println(" source", source)
	}
	fmt.Println("End")
}
