package pty

import (
	"os"
	"strconv"
	"syscall"
	"unsafe"
	
	"github.com/aws/amazon-ssm-agent/agent/log"
)

func open() (pty, tty *os.File, err error) {
	
	var log log.T
	
	log.Info("pty -- aaaaaaa")
	
	p, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		return nil, nil, err
	}
	// In case of error after this point, make sure we close the ptmx fd.
	defer func() {
		if err != nil {
			_ = p.Close() // Best effort.
		}
	}()

	log.Info("pty -- bbbbbbb")
	sname, err := ptsname(p)
	if err != nil {
		return nil, nil, err
	}

	log.Info("pty -- cccccc")
	if err := unlockpt(p); err != nil {
		return nil, nil, err
	}

	log.Info("pty -- ddddddd")
	t, err := os.OpenFile(sname, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return nil, nil, err
	}
	return p, t, nil
}

func ptsname(f *os.File) (string, error) {
	var n _C_uint
	err := ioctl(f.Fd(), syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
	if err != nil {
		return "", err
	}
	return "/dev/pts/" + strconv.Itoa(int(n)), nil
}

func unlockpt(f *os.File) error {
	var u _C_int
	// use TIOCSPTLCK with a zero valued arg to clear the slave pty lock
	return ioctl(f.Fd(), syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&u)))
}
