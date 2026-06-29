package ssh

import (
	"contest-cli/internal/utils"
	"fmt"
	"os/exec"
	"time"
)

type TestResult struct {
	Success  bool
	Duration time.Duration
	Output   string
	Error    error
}

func RunSSH(user, host, key, target, proto, port string) TestResult {
	start := time.Now()
	var err error

	validatedUser, err := utils.GetValidUserInput(user)
	if err != nil {
		return TestResult{
			Success:  false,
			Error:    err,
			Duration: time.Since(start),
			Output:   "",
		}
	}
	targetHost, targetPort, err := utils.SplitTargetAndPort(target)
	if err != nil {
		return TestResult{
			Success:  false,
			Duration: time.Since(start),
			Output:   "",
			Error:    err,
		}
	}
	if port != "" {
		targetPort = port
	}

	remoteCmd := fmt.Sprintf("nc -vz -w 3 %s %s", targetHost, targetPort)

	args := []string{
		"-o", "StrictHostKeyChecking=no",
		"-o", "BatchMode=yes",
		"-o", "ConnectTimeout=5",
	}

	if key != "" {
		args = append(args, "-i", key)
	}

	args = append(args, fmt.Sprintf("%s@%s", validatedUser, host))

	args = append(args, remoteCmd)

	cmd := exec.Command("ssh", args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return TestResult{
			Success:  false,
			Duration: time.Since(start),
			Output:   string(output),
			Error:    err,
		}
	}

	return TestResult{
		Success:  true,
		Duration: time.Since(start),
		Output:   string(output),
		Error:    nil,
	}
}
