package cli

import (
	"flag"
	"fmt"
	"os"
)


func Execute() {
	flag.Usage = func() {
		fmt.Println("Contest CLI - Network Connection Testing")
		fmt.Println("Usage: contest <command> [flags]")
		fmt.Println("\nCommands:")
		fmt.Println("  version    Show application version")
		fmt.Println("\nFlags:")
		flag.PrintDefaults() 
	}


	flag.Parse()


	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1) 
	}


	command := flag.Args()[0]

	switch command {
	case "version":
		fmt.Println("Contest CLI v1.0.0")
		
	case "local":
		runLocal(flag.Args()[1:])
		
	case "ssh":
		fmt.Println("SSH module will be added soon.")
		
	case "k8s":
		fmt.Println("Kubernetes module will be added soon.")
		
	default:
		fmt.Printf("Unknown command: '%s'\n", command)
		fmt.Println("For help, type 'contest --help'.")
		os.Exit(1)
	}
}