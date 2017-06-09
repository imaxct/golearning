package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	attr := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}
	fmt.Println(time.Now(), "aa")
	pid, err := syscall.ForkExec("/usr/bin/sleep", []string{"/usr/bin/sleep", "2s"}, attr)
	fmt.Println(time.Now(), "bb")
	go func(pid int) {
		fmt.Println(time.Now(), "cc")
		fmt.Println("go start.")
		time.Sleep(time.Second)
		fmt.Println("go end.")
		ps, err := os.FindProcess(pid)
		if err != nil {
			panic(err)
		}
		if ps == nil {
			return
		} else {
			fmt.Println("process not done, kill.")
			err := ps.Kill()
			fmt.Println(err)
		}
	}(pid)

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
