package main

import "testing"

func TestValidNames(t *testing.T) {
	name := "validName"
	if !isValidName(name) {
		t.Error(name + " should be a valid name.")
	}
	name = "numb3rs/4r3/ok"
	if !isValidName(name) {
		t.Error(name + " should be a valid name.")
	}
}

func TestInvalidNames(t *testing.T) {
	name := "dangerous/../folder"
	if isValidName(name) {
		t.Error(name + " should not be a valid name.")
	}
	name = "dots/.not/allows/at/beginning"
	if isValidName(name) {
		t.Error(name + " should not be a valid name.")
	}
	name = "dots/not./allows/at/end"
	if isValidName(name) {
		t.Error(name + " should not be a valid name.")
	}
	name = "cannot/end/with/folder/slash/"
	if isValidName(name) {
		t.Error(name + " should not be a valid name.")
	}
}
