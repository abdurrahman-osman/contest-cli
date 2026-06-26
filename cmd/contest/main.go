package main

import (
	"fmt"
	"os"
	"contest-cli/internal/cli"
)


func main() {
	fmt.Println("Merhaba! Contest CLI çalışıyor.")
	fmt.Println("Terminalden gelen veriler:", os.Args)
	cli.Execute()
}