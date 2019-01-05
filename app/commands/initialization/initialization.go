package initialization

import (
	"github.com/parking-lot/app/constants"
	. "github.com/parking-lot/app/model/parking-lot"
	"github.com/pkg/errors"
	"strconv"
)

func Command(args []string) (*ParkingLot, error) {
	slots, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New(constants.SlotsMustBeNumber)
	}

	lot := &ParkingLot{}
	if err := lot.SetSlots(slots); err != nil {
		return nil, err
	}

	return lot, nil
}
