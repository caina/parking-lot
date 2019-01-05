package status

import (
	"bytes"
	"github.com/parking-lot/app/model/parking-lot"
	"strconv"
)

func Command(lot parking_lot.ParkingLot) string {
	builder := bytes.Buffer{}
	builder.WriteString("------------ Status display ----------- \n")
	builder.WriteString("------------ Available Slots: " + strconv.Itoa(lot.GetAvailableSlots()) + " ------- \n")

	builder.WriteString("------------ Parked cars -------------- \n")
	builder.WriteString(lot.ToString())
	return builder.String()
}
