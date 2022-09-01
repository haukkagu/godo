package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func listElements() {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	for i, line := range(lines) {
		if line == "" {
			continue
		}

		fmt.Println(strconv.Itoa(i) + ") " + line)
	}
}

func addElement(element string) {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	_, err = f.Write([]byte(element + "\n"))
	check(err)
}

func removeElement(index int) {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	if index < 0 || index >= len(lines)-1 {
		panic("Index out of bounds")
	}

	f, err :=  os.Create(path)
	check(err)
	defer f.Close()

	for i, line := range(lines) {
		if i == index || line == "" {
			continue
		}
		_, err = f.Write([]byte(line + "\n"))
		check(err)
	}
}

var usr, _ = user.Current()
var path = usr.HomeDir + "/.local/share/godo.txt"
func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		listElements()
		return
	}

	switch(args[0]) {
		case "help":
			fmt.Println(`godo is a very simple task management application
Usage:
	help - list commands
	list - list all elements in order
	add - add an element to the list
	remove - remove an element from the list by its index`)
		case "list":
			listElements()
		case "add":
			element := ""
			for i, arg := range(args[1:]) {
				element += arg
				if i < len(args)-1 {
					element += " "
				}
			}

			addElement(element)
		case "remove":
			index, err := strconv.Atoi(args[1])
			check(err)
			removeElement(index)
		default:
			fmt.Println("Unknown argument")
	}
}
