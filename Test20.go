package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	attr := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}
	pid, err := syscall.ForkExec("/usr/bin/ls", []string{"/usr/bin/ls", "-al"}, attr)
	if err != nil {
		panic(err)
	}
	wpid, err := syscall.Wait4(pid, nil, 0, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(wpid)
	fmt.Println("sub process done.")
}
