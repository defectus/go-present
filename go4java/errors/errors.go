package main

import (
	"errors"
)

func CanFail(shouldFail, shouldPanic bool) error {
	if shouldFail {
		return errors.New("was told to fail")
	}
	if shouldPanic {
		panic("was told to panic")
	}
	return nil
}

func main() {
	CanFail(true, false)
	CanFail(false, true)
}
