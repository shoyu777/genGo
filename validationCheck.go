package genGo

import "fmt"

type yearOutOfScope struct {
	year int
}

func (e *yearOutOfScope) Error() string {
	return fmt.Sprintf("The year %v was Out of Scope. Please set 645 or above.", e.year)
}

func yearRangeCheck(y int) error {
	if y < 645 {
		return &yearOutOfScope{year: y}
	}
	return nil
}
