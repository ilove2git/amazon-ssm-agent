// +build !windows

package pty

import (
	"os"
	"os/exec"
	"syscall"
	"github.com/aws/amazon-ssm-agent/agent/log"
)

// Start assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// and c.Stderr, calls c.Start, and returns the File of the tty's
// corresponding pty.
func Start(c *exec.Cmd) (pty *os.File, err error) {
	pty, tty, err := Open()
	if err != nil {
		return nil, err
	}
	
	log.Info("pyt 11111111")
	
	defer tty.Close()
	c.Stdout = tty
	c.Stdin = tty
	c.Stderr = tty
	
	log.Info("pyt 22222222")
	
	if c.SysProcAttr == nil {
		c.SysProcAttr = &syscall.SysProcAttr{}
		log.Info("pty 3333333")
	}
	c.SysProcAttr.Setctty = true
	c.SysProcAttr.Setsid = true
	
	log.Info("pty 44444444")
	
	err = c.Start()
	if err != nil {
		log.Info("pty 55555555")
		pty.Close()
		return nil, err
	}
	return pty, err
}
