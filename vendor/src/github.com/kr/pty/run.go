// +build !windows

package pty

import (
	"os"
	"os/exec"
	"syscall"
	"log"
)

// Start assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// and c.Stderr, calls c.Start, and returns the File of the tty's
// corresponding pty.
func Start(c *exec.Cmd) (pty *os.File, err error) {
	pty, tty, err := Open()
	if err != nil {
		return nil, err
	}
	
	log.Println("pyt 11111111")
	
	defer tty.Close()
	c.Stdout = tty
	c.Stdin = tty
	c.Stderr = tty
	
	log.Println("pyt 22222222")
	
	if c.SysProcAttr == nil {
		c.SysProcAttr = &syscall.SysProcAttr{}
		log.Println("pty 3333333")
	}
	c.SysProcAttr.Setctty = true
	c.SysProcAttr.Setsid = true
	
	log.Println("pty 44444444")
	
	err = c.Start()
	if err != nil {
		log.Println("pty 55555555")
		pty.Close()
		return nil, err
	}
	return pty, err
}
