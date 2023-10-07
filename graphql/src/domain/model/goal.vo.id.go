package model

import (
	"fmt"
)

// GoalID 目標ID
type GoalID int

func (id GoalID) String() string {
	return fmt.Sprintf("%d", id)
}
