package dberr

import "strings"

// Type represents a DB error type.
type Type int

const (
	Unknown Type = iota
	Unique
	Check
	Trigger
)

func Handle(err error, f func(t Type, v string) error) error {
	if err == nil {
		return nil
	}
	t, v := parse(err.Error())
	if t == Unknown {
		return err
	}
	return f(t, v)
}

func parse(msg string) (Type, string) {
	const (
		unique  = "UNIQUE constraint failed: "
		check   = "CHECK constraint failed: "
		trigger = "TRIGGER failed: "
	)
	switch {
	case strings.HasPrefix(msg, unique):
		return Unique, strings.TrimPrefix(msg, unique)
	case strings.HasPrefix(msg, check):
		return Check, strings.TrimPrefix(msg, check)
	case strings.HasPrefix(msg, trigger):
		return Trigger, strings.TrimPrefix(msg, trigger)
	default:
		return Unknown, msg
	}
}
