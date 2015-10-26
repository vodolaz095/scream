package scream

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func existsAsExecutable(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return (fi.Mode().Perm()&0111 != 0), nil
}

// SanityCheck throws error, when the notification-daemon is not installed or started properly
func SanityCheck() error {
	doClientExists, err := existsAsExecutable("/usr/bin/notify-send")
	if err != nil {
		return err
	}
	if !doClientExists {
		return fmt.Errorf("notify-send command does not exist")
	}
	doServerExists, err := existsAsExecutable("/usr/libexec/notification-daemon")
	if err != nil {
		return err
	}
	if !doServerExists {
		return fmt.Errorf("notify-send command does not exist")
	}

	pidof := exec.Command("/usr/bin/pidof", "/usr/libexec/notification-daemon")
	var pid bytes.Buffer
	pidof.Stdout = &pid
	err = pidof.Run()
	if err != nil {
		return err
	}

	if pid.String() != "" {
		fmt.Printf("Notification-Daemon is running with PID %v\n", pid.String())
	} else {
		return fmt.Errorf("/usr/libexec/notification-daemon is not running")
	}

	err = notifySend("Notification daemon", "Started properly!")
	if err != nil {
		if err.Error() == "fork/exec /usr/bin/notify-send: no such file or directory" {
			return fmt.Errorf("unable to execute notify-send")
		}
		return err
	}
	return nil
}
