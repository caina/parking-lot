package exit

import "github.com/parking-lot/app/client"

func Command(client *client.Client) {
	client.Send <- "Thanks for using Parking lot!"
	client.Stop <- true
}
