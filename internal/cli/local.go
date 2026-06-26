package cli

import (
	"flag"
	"fmt"
	"net"
	"time"
)

// runLocal "local" alt komutu çalıştırıldığında devreye girer.
func runLocal(args []string) {
	// Sadece bu komuta özel bayrak kümesi
	localCmd := flag.NewFlagSet("local", flag.ExitOnError)

	// İhtiyacımız olan parametreler
	targetFlag := localCmd.String("target", "", "Hedef IP ve Port (Zorunlu) (örn: 192.168.1.10:9092)")
	protoFlag := localCmd.String("proto", "tcp", "Kullanılacak protokol (tcp veya udp)")
	timeoutFlag := localCmd.Duration("timeout", 3*time.Second, "Zaman aşımı")

	// Kullanıcının girdiği verileri ayrıştırıyoruz
	localCmd.Parse(args)

	// Validasyon (Doğrulama): Hedef girilmemişse hata verip kullanımı gösteriyoruz.
	if *targetFlag == "" {
		fmt.Println("Hata: --target bayrağı zorunludur!")
		localCmd.Usage()
		return
	}

	// Task 1'in sonu. Task 2'de buraya bağlantı kodlarını yazacağız.
	fmt.Printf("[LOCAL TEST] %s üzerinden %s hedefine bağlanılacak...\n", *protoFlag, *targetFlag)

	conn, err := net.DialTimeout(*protoFlag, *targetFlag, *timeoutFlag)
	if err != nil {
		fmt.Printf("Bağlantı Başarısız: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Bağlantı Başarılı!")
}