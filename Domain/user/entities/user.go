package userentities

import (
	"errors"
	"strings"
)

type User struct {
	Name string
	Age  int
}

var U *User = &User{}

func ValidationPayload() error {
	n := strings.TrimSpace(U.Name)
	if len(n) == 0 {
		return errors.New("not contain needed property")
	}
	return nil
}
