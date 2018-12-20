package main

import (
	"bufio"
	"fmt"
	"github.com/parking-lot/app"
	"github.com/parking-lot/app/commands/exit"
	"github.com/parking-lot/app/commands/help"
	"github.com/parking-lot/app/commands/initialization"
	"github.com/parking-lot/app/commands/leave"
	"github.com/parking-lot/app/commands/park"
	"github.com/parking-lot/app/commands/status"
	"github.com/parking-lot/app/constants"
	"os"
	"strings"
)

func main() {
	fmt.Println("------------------- Welcome to parking lot! ---------------")
	fmt.Println("------ use the command help to list all the options -------")

	resultChan := app.Client{
		Send: make(chan string),
		Stop: make(chan bool),
	}
	defer close(resultChan.Stop)
	defer close(resultChan.Send)

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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Fields(input)

		// should be dynamic
		switch args[0] {
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
		case "init":
			go func() {
				initialization.Command(args[1:], &resultChan)
				return
			}()
		case "status":
			go func() {
				status.Command(&resultChan)
				return
			}()
		case "park":
			go func() {
				park.Command(args[1:], &resultChan)
				return
			}()
		case "leave":
			go func() {
				leave.Command(args[1:], &resultChan)
				return
			}()
		default:
			go func() {
				resultChan.Send <- constants.CommandNotFound
				return
			}()
		}
	}
}
