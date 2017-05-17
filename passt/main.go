package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

type cmd uint32

const (
	cmdLs cmd = 1 << iota
	cmdAdd
	cmdShow
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func baseDir() string {
	return os.Getenv("HOME") + "/.passt"
}

func parseCmd() cmd {
	add := flag.Bool("a", false, "Add new password")
	show := flag.Bool("s", false, "Show password")
	list := flag.Bool("l", true, "List passwords")
	flag.Parse()

	if *add {
		return cmdAdd
	}
	if *show {
		return cmdShow
	}
	if *list {
		return cmdLs
	}
	return cmdLs
}

func list() {
	tree, err := exec.Command("tree", baseDir()).Output()
	check(err)
	fmt.Print(string(tree))
}

func filePath() string {
	if len(flag.Args()) == 0 {
		panic("Missing filename")
	}
	name := flag.Arg(0)
	if !isValidName(name) {
		panic("Invalid file name: " + name)
	}

	return baseDir() + "/" + name
}

func add() {
	fp := filePath()
	dir := path.Dir(fp)

	err := os.MkdirAll(dir, 0700)
	check(err)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	passwordLine, err := reader.ReadString('\n')
	check(err)
	password := passwordLine[:len(passwordLine)-1]

	file, err := os.Create(fp)
	check(err)
	defer file.Close()

	_, err = file.WriteString(password)
	check(err)
	file.Sync()
}

func show() {
	file, err := ioutil.ReadFile(filePath())
	check(err)
	fmt.Print(string(file))
}

func main() {
	switch parseCmd() {
	case cmdAdd:
		add()
	case cmdShow:
		show()
	case cmdLs:
		list()
	default:
		list()
	}
}
