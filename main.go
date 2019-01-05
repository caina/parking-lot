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
				result := help.Command()
				resultChan.Send <- result
				return
			}()
		case "exit":
			go func() {
				message, exitBool := exit.Command()

				resultChan.Send <- message
				resultChan.Stop <- exitBool
				return
			}()
		case "init":
			go func() {
				newLot, err := initialization.Command(args[1:])
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				resultChan.Parking = *newLot
				resultChan.Send <- strings.Replace(constants.ParkingLotGenerated, "%c", args[1:][0], 1)
				return
			}()
		case "status":
			go func() {
				result := status.Command(resultChan.Parking)
				resultChan.Send <- result

				return
			}()
		case "park":
			go func() {
				ticket, err := park.Command(args[1:], &resultChan.Parking)
				if err != nil {
					resultChan.Send <- err.Error()
					return
				}

				_ = resultChan.Parking.Park(*ticket)
				resultChan.Send <- strings.Replace(constants.TicketGenerated, "%c", ticket.ToString(), 1)
				return
			}()
		case "leave":
			go func() {
				tickets := leave.Command(args[1:], resultChan.Parking)

				resultChan.Parking.SetTickets(tickets)
				resultChan.Send <- constants.CarLeftPark
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
