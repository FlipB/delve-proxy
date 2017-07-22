// +build windows

package main

import (
	"os"
	"os/exec"
	"log"
	"time"
)
type delve struct {
	proc *os.Process
}

func (d *delve) startDelve() {
	d.stopDelve()
	cmd := exec.Command("dlv", "debug", "--listen", *remoteAddr, "--headless")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	d.proc = cmd.Process;
}

func (d *delve) stopDelve() {
	d.proc.Kill()
	c, ce := chanWaitProcess(d.proc)
	timeout := time.After(time.Second*5)
	select {
		case state := <- c:
			if (state.Exited()) {
				println("exited")
			} else {
				println("wtf")
			}
		case err := <- ce:
			println("error waiting for process to end.")
		case <-timeout:
			println("Process didnt stop within timeout.")
	}
	
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

func chanWaitProcess(proc *os.Process) (chan *os.ProcessState, chan error) {
	c := make(chan *os.ProcessState, 1)
	ce := make(chan error, 1)
	go func() {
		state, err := proc.Wait()
		if (err != nil) {
			ce <- err
		}
		c <- state
	}()
}
