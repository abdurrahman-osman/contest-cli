package cli

import (
	"flag"
	"fmt"
	"net"
	"time"
	"strings"
)

// runLocal "local" alt komutu çalıştırıldığında devreye girer.
func runLocal(args []string) {
	// Sadece bu komuta özel bayrak kümesi
	localCmd := flag.NewFlagSet("local", flag.ExitOnError)

	// İhtiyacımız olan parametreler
	targetFlag := localCmd.String("target", "", "Target IP ve Port (required) (örn: 192.168.1.10)")
	portFlag := localCmd.String("port", "80", "Target Port (örn: 9092)(Default: 80)")
	protoFlag := localCmd.String("proto", "tcp", "Protocol (tcp veya udp)")
	timeoutFlag := localCmd.Duration("timeout", 3*time.Second, "Timeout Seconds")

	// Kullanıcının girdiği verileri ayrıştırıyoruz
	localCmd.Parse(args)

	// Validasyon (Doğrulama): Hedef girilmemişse hata verip kullanımı gösteriyoruz.
	if *targetFlag == "" {
		fmt.Println("Error: --target flag is required!")
		localCmd.Usage()
		return
	}

	// Task 1'in sonu. Task 2'de buraya bağlantı kodlarını yazacağız.
	fmt.Printf("[LOCAL TEST] %s üzerinden %s hedefine bağlanılacak...\n", *protoFlag, *targetFlag)

	var addr string
	if !strings.Contains(*targetFlag, ":") {
	    addr = fmt.Sprintf("%s:%s", *targetFlag, *portFlag)
	} else {
	    addr = *targetFlag
	}

	conn, err := net.DialTimeout(*protoFlag, addr, *timeoutFlag)
	if err != nil {
		fmt.Printf("Connection Failed: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connection Successful!")
}