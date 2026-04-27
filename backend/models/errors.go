package models

type NotFoundError struct {
	Resource string
	ID string
}

func (e *NotFoundError) Error() string {
	return e.Resource + " not found: " + e.ID
}