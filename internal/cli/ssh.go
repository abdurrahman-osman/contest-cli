package cli

import (
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
	portFlag := sshCmd.String("port", "80", "Target Port (örn: 9092)(Default: 80)")

	sshCmd.Parse(args)

	if *hostsFlag == "" || *targetFlag == "" {
		fmt.Println("Error: --hosts and --target flags are required!")
		sshCmd.Usage()
		return
	}

	rawHosts := strings.Split(*hostsFlag, ",")

	for _, host := range rawHosts {
		cleanHost := strings.TrimSpace(host)
		
		fmt.Printf("-> SSH ile %s@%s Connection\n with privkey %s and proto %s and port %s\n", *userFlag, cleanHost, *keyFlag, *protoFlag, *portFlag)
		//TODO: link ssh connection logic
		
	}
}