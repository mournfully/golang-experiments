package main

import (
	"fmt"
	"testing"
)

func Test_SayHello_ValidArgument(t *testing.T) {
	name := "Mert"
	result := sayHello(name)
	expected := fmt.Sprintf("Hello %s", name)

	if result != expected {
		t.Errorf("\"sayHello('%s'\" FAILED, expected -> %v, got -> %v)", name, expected, result)
	} else {
		t.Logf("\"sayHello('%s'\" SUCCEDED, expected -> %v, got -> %v)", name, expected, result)
	}
}

func Test_SayGoodbyte(t *testing.T) {
	name := "Mert"
	result := sayGoodbye(name)
	expected := fmt.Sprintf("Bye Bye %s!", name)

	if result != expected {
		t.Errorf("\"sayGoodbye('%s'\" FAILED, expected -> %v, got -> %v)", name, expected, result)
	} else {
		t.Logf("\"sayGoodbye('%s'\" SUCCEDED, expected -> %v, got -> %v)", name, expected, result)
	}
}
