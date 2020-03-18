package main

import (
	"fmt"

	"github.com/rezasr/filegen-from-excel/internal/prog"
	"github.com/rezasr/filegen-from-excel/internal/progcli"
)

func init() {
	prog.InitUsage()
}

func main() {
	if prog.Mode.IsCli() {
		err := progcli.Main()
		if err != nil {
			fmt.Println(err)
		}
	}
}
