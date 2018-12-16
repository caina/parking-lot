package main

import (
	"bufio"
	"fmt"
	"github.com/parking-lot/app/commands/exit"
	"github.com/parking-lot/app/commands/help"
	"github.com/parking-lot/app/config"
	"os"
)

func main() {
	fmt.Println("------------------- Welcome to parking lot! ---------------")
	fmt.Println("------ Inform the number of slots in this parking lot -----")

	resultChan := config.Client{
		Send: make(chan string),
		Stop: make(chan bool),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		// should be dynamic
		switch text {
		case "help":
			go func() {
				help.Command(&resultChan)
				return
			}()
		case "exit":
			go func() {
				exit.Command(&resultChan)
				return
			}()
		}

		go func() {
			for {
				select {
				case result := <-resultChan.Send:
					fmt.Println(result)
				case <-resultChan.Stop:
					os.Exit(1)
				}
			}
		}()
	}
}
