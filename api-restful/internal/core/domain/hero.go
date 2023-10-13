package domain

import (
	"fmt"
)

type Hero struct {
	ID         int
	Name       string
	ActualName string
}

func (h Hero) String() string {
	return fmt.Sprintf("Hero {ID=%d, Name=%s, ActualName=%s}", h.ID, h.Name, h.ActualName)
}
