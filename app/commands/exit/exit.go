package exit

import "github.com/parking-lot/app/config"

func Command(client *config.Client) {
	client.Send <- "Thanks for using Parking lot!"
	client.Stop <- true
}
