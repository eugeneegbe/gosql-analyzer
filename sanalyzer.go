// Author: Alangi Derick N
// LICENSE: MIT, Copyright (c) 2017 Alangi Derick
// Description: Entry point for the SQL analyzer tool

package main

import (
	"bufio"
	"fmt"
	color "github.com/fatih/color"
	"os"
	exec "os/exec"
	"runtime"
	"strings"
)

var clear map[string]func() // create a map for storing clear functions

// functions to perform clear screen operations in linux and windows platform
func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") // Windows systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") // Mac OS systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ClearScreen clears the screen and prepare the console for the analyzer
func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported so screen can't be cleared!")
	}
}

func main() {

	// Clear screen and prepare the analyzer console
	ClearScreen()

	// Create a new color object
	c := color.New(color.FgGreen)
	bold := color.New(color.FgGreen, color.Bold)

	bold.Println("Welcome to the Go SQL Query Analyzer (GoSQA) version 1.0")
	bold.Println("Developed by: Alangi Derick N (alangiderick@gmail.com)")
	fmt.Println()

	for {
		c.Printf("gosql-query-analyzer> ")

		reader := bufio.NewReader(os.Stdin)

		query, _ := reader.ReadString('\n')

		if strings.TrimRight(query, "\n") == "quit" {
			//fmt.Println("Exiting Go SQL Query Analyzer...")
			c.Println("Press Ctrl + C to exit the console.")
		}
		c.Println("query: ", query)
		// Enter Ctrl + C to quit.
	}

}
