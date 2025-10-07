package sandbox

import (
	"log"
	"os/exec"
	"syscall"
	"time"
)

func RunWithTimeout(command string) int {
	cmd := exec.Command("bash", "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()

	done := make(chan interface{}, 1)
	go func() {
		cmd.Wait()
		done <- nil
	}()

	select {
	case <-done:
		return cmd.ProcessState.ExitCode()
	case <-time.After(time.Second):
		log.Printf("command %q timed out, killing process", command)
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		return -1
	}
}
