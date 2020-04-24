package main

//go:generate go run generate-gladefile.go

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rezasr/filegen-from-excel/internal/prog"
	"github.com/rezasr/filegen-from-excel/internal/progcli"
)

func init() {
	prog.ExePath, _ = os.Executable()
	prog.ExePath, _ = filepath.Abs(prog.ExePath)
	prog.ExePath, _ = filepath.EvalSymlinks(prog.ExePath)
	prog.WorkingDir = filepath.Dir(prog.ExePath)

	prog.InitUsage()
}

func main() {
	if prog.Mode.IsCli() {
		err := progcli.Main()
		if err != nil {
			fmt.Println(err)
		}
	}

	os.Exit(0)
}
