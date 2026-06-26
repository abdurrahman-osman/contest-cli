package cli

import (
	"flag"
	"fmt"
	"os"
)

// Execute fonksiyonu uygulamamızın başlangıç noktasıdır.
func Execute() {
	// 1. Kendi özel yardım menümüzü tanımlıyoruz.
	flag.Usage = func() {
		fmt.Println("Contest CLI - Network Connection Testing")
		fmt.Println("Usage: contest <command> [flags]")
		fmt.Println("\nCommands:")
		fmt.Println("  version    Show application version")
		fmt.Println("\nFlags:")
		flag.PrintDefaults() // Tanımladığımız tüm bayrakları (--help vs) ekrana basar
	}

	// 2. İşletim sisteminden gelen bayrakları okur ve işler
	flag.Parse()

	// 3. Kullanıcı hiçbir alt komut (subcommand) girmediyse:
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1) // Hata ile çıkış yap
	}

	// 4. Kendi komut yönlendiricimizi (Router) yazıyoruz:
	command := flag.Args()[0]

	switch command {
	case "version":
		fmt.Println("Contest CLI v1.0.0")
		
	case "local":
		// 'local' komutunu çalıştır ve sonrasındaki tüm parametreleri ona gönder
		runLocal(flag.Args()[1:])
		
	case "ssh":
		// Henüz yazılmadı, ilerleyen tasklarda dolduracağız.
		fmt.Println("SSH module will be added soon.")
		
	case "k8s":
		// Henüz yazılmadı, ilerleyen tasklarda dolduracağız.
		fmt.Println("Kubernetes module will be added soon.")
		
	default:
		fmt.Printf("Unknown command: '%s'\n", command)
		fmt.Println("For help, type 'contest --help'.")
		os.Exit(1)
	}
}