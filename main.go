package main

import (
	"fmt"
)

func main() {

	slackMessenger := slack{}
	teamsMessenger := teams{}
	jeff := user{Name: "jeff"}
	morty := user{Name: "morty"}

	_ = send(slackMessenger, morty, "hi this is jeff")
	_ = send(teamsMessenger, jeff, "hi this is morty")

	//var aIsGreaterThanb, xIsGreaterThany bool
	//a := 1
	//b := 2
	////aIsGreaterThanb = isGreaterThanI(a, b)
	//aIsGreaterThanb = isGreaterThan(a, b)
	//
	//fmt.Printf("aIsGreaterThanb: %t", aIsGreaterThanb)
	//x := int32(a)
	//y := int32(b)
	//xIsGreaterThany = isGreaterThanI(x, y)
	//xIsGreaterThany = isGreaterThan(x, y)

}

func sendItTheUglyWay(messenger interface{}, receiver user, text string) error {
	if len(text) > 20 {
		return fmt.Errorf("message too long")
	}
	var err error
	switch messenger.(type) {
	case slack:
		err = messenger.(*slack).deliver(receiver.Name, text)
	case teams:
		err = messenger.(*teams).deliver(receiver.Name, text)
	default:
		return fmt.Errorf("unknown messenger")
	}
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	Name string
}

// This is the function of interest.
// send provides functionality to send a text via some messenger to a user.
func send(messenger messenger, receiver user, text string) error {
	// Some Business Logic
	if len(text) > 20 {
		return fmt.Errorf("message too long")
	}

	// Call some method I relay on. I know nothing about its implementation nor what messenger it could be.
	err := messenger.deliver(receiver.Name, text)

	// Some Business Logic
	if err != nil {
		return err
	}
	return nil
}

// The interface the send function relays on.
type messenger interface {
	deliver(receiver, text string) error
}

// Example implementation of a messenger
type slack struct {
}

func (s slack) deliver(receiver, text string) error {
	fmt.Printf("Sent: '%s' via SLACK to '%s'\n", text, receiver)
	return nil
}

// Example implementation of a messenger
type teams struct {
}

func (t teams) deliver(receiver, text string) error {
	fmt.Printf("Sent: '%s' via TEAMS to '%s\n", text, receiver)
	return nil
}

func isGreaterThanI(a, b int) bool {
	if a > b {
		return true
	}
	return false
}

func isGreaterThanF(a, b float64) bool {
	if a > b {
		return true
	}
	return false
}

type NumberConstraint interface {
	int | float64
}

func isGreaterThan[T ~int](a, b T) bool {
	if a > b {
		return true
	}
	return false
}

func isEqual[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}
