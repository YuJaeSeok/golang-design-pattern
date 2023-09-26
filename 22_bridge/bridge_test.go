package bridge

import (
	"fmt"
	"testing"
)

func ExampleCommonSMS() {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}

func ExampleCommonEmail() {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Email
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}

func ExampleUrgencyEmail() {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}

func TestBridge(t *testing.T) {
	ExampleCommonSMS()
	fmt.Println()
	ExampleCommonEmail()
	fmt.Println()
	ExampleUrgencySMS()
	fmt.Println()
	ExampleUrgencyEmail()
}
