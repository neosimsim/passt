package main

import (
	"regexp"
)

const validChars = "[0-9a-zA-Z-_]"

// files may contain a single dot between valid characters
const fileMatch = validChars + "+(\\." + validChars + "+)*"

// '/' are only allowed between file, to ensure the last file is in fact a file and no dir
const pathMatch = "(" + fileMatch + ")+(/" + fileMatch + ")*"

func isValidName(name string) bool {
	isValid, err := regexp.MatchString("^"+pathMatch+"$", name)
	if err != nil {
		panic(err)
	}
	return isValid
}
