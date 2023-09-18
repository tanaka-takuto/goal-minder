package model

import (
	"fmt"
)

// AccountID アカウントID
type AccountID int

func (id AccountID) String() string {
	return fmt.Sprintf("%d", id)
}
