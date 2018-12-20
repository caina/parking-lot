package exit

import (
	"github.com/parking-lot/app"
	"github.com/parking-lot/app/constants"
)

func Command(client *app.Client) {
	client.Send <- constants.ExitMessage
	client.Stop <- true
}
