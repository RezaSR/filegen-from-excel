package main

import (
	"flag"
	"fmt"

	"github.com/rezasr/filegen-from-excel/internal/prog"
	"github.com/rezasr/filegen-from-excel/internal/progcli"
)

func init() {
	prog.Mode.Set(flag.Bool("c", false, "Run in CLI mode and do not open GUI"))
	prog.TemplateFile.Set(flag.String("t", "", "Template file"))
	prog.OutDir.Set(flag.String("o", "out", "Output directory"))
	prog.OutFileName.Set(flag.String("f", "[0000].txt", "Output file names"))
	prog.DataFile.Set(flag.String("d", "", "Excel data file"))
	flag.Parse()
}

func main() {
	if prog.Mode.IsCli() {
		err := progcli.Main()
		if err != nil {
			fmt.Println(err)
		}
	}
}
