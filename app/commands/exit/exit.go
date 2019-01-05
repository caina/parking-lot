package exit

import (
	"github.com/parking-lot/app/constants"
)

func Command() (string, bool) {
	return constants.ExitMessage, true
}
