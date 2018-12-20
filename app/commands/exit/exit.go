package exit

import (
	"github.com/parking-lot/app/client"
	"github.com/parking-lot/app/constants"
)

func Command(client *client.Client) {
	client.Send <- constants.ExitMessage
	client.Stop <- true
}
