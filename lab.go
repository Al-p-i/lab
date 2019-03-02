package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(help())
		return
	}

	switch cmd := args[0]; cmd {
	case "new":
		newCmd()
	case "checkout":
		checkoutCmd()
	case "ls":
		lsCmd()
	case "info":
		infoCmd()
	default:
		noSuchCmd(cmd)
	}
}

func noSuchCmd(cmd string) {
	println("No such command ", cmd)
}

func infoCmd() {
	println("info")
	cmd := exec.Command("git", "branch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func lsCmd() {
	println("ls")
}

func checkoutCmd() {
	println("checkout")
}

func newCmd() {
	println("new")
}

func help() string {
	return "help"
}
