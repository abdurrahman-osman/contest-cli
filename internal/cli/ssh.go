package cli

import (
	"contest-cli/internal/tester/ssh"
	"flag"
	"fmt"
	"strings"
)

func runSSH(args []string) {
	sshCmd := flag.NewFlagSet("ssh", flag.ExitOnError)

	hostsFlag := sshCmd.String("hosts", "", "Target host IP list (required)Example: 10.0.0.1,10.0.0.2)")
	userFlag := sshCmd.String("user", "root", "SSH User Name (Default: root)")
	keyFlag := sshCmd.String("key", "", "Optional: SSH Private Key file path")
	targetFlag := sshCmd.String("target", "", "Target IP:Port to test connectivity (Required) (Example: 192.168.1.10:9092)")
	protoFlag := sshCmd.String("proto", "tcp", "Protocol to use (tcp or udp)")
	portFlag := sshCmd.String("port", "", "Target Port (e.x: 9092)")

	err := sshCmd.Parse(args)
	if err != nil {
		fmt.Printf("failed to parse args: %v\n", err)
		return
	}

	if *hostsFlag == "" || *targetFlag == "" {
		fmt.Println("Error: --hosts, --target flags are required!")
		sshCmd.Usage()
		return
	}

	rawHosts := strings.Split(*hostsFlag, ",")

	for _, host := range rawHosts {
		cleanHost := strings.TrimSpace(host)
		result := ssh.RunSSHGo(*userFlag, cleanHost, *keyFlag, *targetFlag, *protoFlag, *portFlag)
		if !result.Success {
			fmt.Printf("Output: %s\nError: %v\n", result.Output, result.Error)
		} else {
			fmt.Printf("Success: %s\n", result.Output)
		}
	}
}
