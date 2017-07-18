package main

import (
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	HELP = `Usage: single port command args...
port      the port to listen to (should be greater than 1024 if not root)
command   the command to run
args      the command arguments`
)

// Run command with given arguments and return exit value.
func Execute(command string, args ...string) int {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	exit := 0
	if err != nil {
		message := err.Error()
		if !strings.HasPrefix(message, "exit status") {
			println(message)
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exit = ws.ExitStatus()
		} else {
			exit = -4
		}
	}
	return exit
}

// Run a TCP server on given port to ensure that a single instance is running
// on a machine. Fails if another instance is already running.
func Singleton(port int) (net.Listener, error) {
	listener, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			listener.Accept()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return listener, nil
}

func main() {
	if len(os.Args) < 3 {
		println("ERROR you must pass port and command on command line")
		println(HELP)
		os.Exit(-1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("ERROR port number '" + os.Args[1] + "' is invalid")
		println(HELP)
		os.Exit(-2)
	}
	command := os.Args[2]
	args := os.Args[3:]
	listener, err := Singleton(port)
	if err != nil {
		println("ERROR another instance is running")
		os.Exit(-3)
	}
	exit := Execute(command, args...)
	listener.Close()
	if exit != 0 {
		os.Exit(exit)
	}
}
