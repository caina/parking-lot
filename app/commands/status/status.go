package status

import (
	"bytes"
	"github.com/parking-lot/app"
	"strconv"
)

func Command(client *app.Client) {
	builder := bytes.Buffer{}
	builder.WriteString("------------ Status display ----------- \n")
	builder.WriteString("------------ Available Slots: " + strconv.Itoa(client.Parking.GetAvailableSlots()) + " ------- \n")

	builder.WriteString("------------ Parked cars -------------- \n")
	builder.WriteString(client.Parking.ToString())
	client.Send <- builder.String()
}
