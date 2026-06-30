package ssh

import (
	"contest-cli/internal/utils"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func RunSSH2(user, host, key, target, proto, port string) TestResult {
	start := time.Now()

	validatedUser, err := utils.GetValidUserInput(user)
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: err}
	}

	targetHost, targetPort, err := utils.SplitTargetAndPort(target)
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: err}
	}

	// check port provided
	if port != "" {
		targetPort = port
	}
	if targetPort == "" {
		return TestResult{
			Success:  false,
			Duration: time.Since(start),
			Error:    fmt.Errorf("target port is required"),
		}
	}

	// Add default key if not entered.
	if key == "" {
		key = "~/.ssh/id_rsa"
	}

	expandedKeyPath, err := utils.ExpandHomeDir(key)
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: fmt.Errorf("invalid key path: %v", err)}
	}

	keyBytes, err := os.ReadFile(expandedKeyPath)
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: fmt.Errorf("failed to read private key: %v", err)}
	}

	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: fmt.Errorf("failed to parse private key: %v", err)}
	}

	config := &ssh.ClientConfig{
		User: validatedUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", host), config)
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: fmt.Errorf("ssh dial failed: %v", err)}
	}

	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("failed to close ssh client: %v", err)
		}
	}(client)

	session, err := client.NewSession()
	if err != nil {
		return TestResult{Success: false, Duration: time.Since(start), Error: fmt.Errorf("failed to create session: %v", err)}
	}
	defer func(session *ssh.Session) {
		if err := session.Close(); err != nil && err != io.EOF {
			fmt.Printf("session close error: %v\n", err)
		}
	}(session)

	remoteCmd := fmt.Sprintf("nc -vz -w 3 %s %s", targetHost, targetPort)

	output, err := session.CombinedOutput(remoteCmd)
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
