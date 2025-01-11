package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) NotFound() string {

	return fmt.Sprintf("Entity %s with ID %s not found", e.Entity, e.ID)
}
