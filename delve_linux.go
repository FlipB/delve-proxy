// +build linux
package main

import (
	"os"
	"os/exec"
	"syscall"
)


type delve struct {
	proc *os.Process
}

func (d *delve) startDelve() {
	d.stopDelve()
	cmd = exec.Command("dlv", "debug", "--listen", *remoteAddr, "--headless")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	d.proc = cmd.Process;
}

func (d *delve) stopDelve() {
	var s os.Signal
	s = syscall.SIGTERM
	d.proc.Signal(s)
	if cmd != nil {
		cmd.Process.Signal(os.Kill)
		time.Sleep(time.Millisecond * 100)
		cmd.Process.Release()
		cmd = nil
	}
	time.Sleep(time.Millisecond * 100)
	exec.Command("screen", "-S", "delve", "-X", "kill").Run()
	time.Sleep(time.Millisecond * 500)
	exec.Command("pkill", "-SIGKILL dlv").Run()
	time.Sleep(time.Millisecond * 500)
	exec.Command("pkill", "-SIGKILL debug").Run()
	time.Sleep(time.Millisecond * 200)
}
