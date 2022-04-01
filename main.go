package main

import (
	"fmt"
)

func main() {

	slackMessenger := slack{}
	teamsMessenger := teams{}
	//rick := admin{Name: "rick", FavoriteMessenger: slackMessenger}
	rick := user{Name: "rick"}
	sendItTheUglyWay(slackMessenger, rick, "I'm not stupid")

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

	fmt.Println(someGenericsF("a", "b"))

}

func sendItTheUglyWay(messenger interface{}, receiver user, text string) error {
	if len(text) > 20 {
		return fmt.Errorf("message too long")
	}
	var err error
	switch messenger.(type) {
	case slack:
		err = messenger.(slack).deliver(receiver.Name, text)
	case teams:
		err = messenger.(teams).deliver(receiver.Name, text)
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

// Interfaces are implemented implicitly
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
//
// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

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

// Task 1:
// * append the user to hold a favorite messenger instance
// * implement a function sendToBuddy in a way that it gets a buddy and a text, and calls the corresponding deliver-method
// 	 * a buddy could be a user, admin, or even a guest?! doesnt matter it should be able to do its stuff anyway!
type admin struct {
	Name              string
	FavoriteMessenger messenger
}
type guest struct {
	Name              string
	FavoriteMessenger messenger
}
type randomVisitor struct {
	Name string
}

func sendToBuddy(buddy interface{}, text string) error {
	return nil
}

// some generics stuff

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

// Task 3:
// * implement a generics function for user, admin and guest, but not a randomVisitor that says "I'm not stupid"
func (u user) greet() {
	fmt.Printf("I'm not stupid")
}

// Task 4:
// * implement a function that adds to variable of any integer based type, float32 and string, but not more
// * the function should return the result
func someNotGeneric(a, b interface{}) (interface{}, error) {
	err := fmt.Errorf("nah")
	switch a.(type) {
	case string:
		switch b.(type) {
		case string:
			return fmt.Sprint(a) + fmt.Sprint(b), nil
		default:
			return nil, err
		}
	default:
		return nil, err
	}

}

func someGenericsF(a, b string) string {
	return a + b
}
