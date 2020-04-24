// +build ignore

package main

import (
	"io"
	"os"
)

func main() {
	out, _ := os.Create("internal/proggtk/gladefile.go")
	out.Write([]byte("package proggtk \n\nconst (\n"))
	out.Write([]byte("\tgladeStr = `"))
	f, _ := os.Open("internal/proggtk/glade/main.glade")
	io.Copy(out, f)
	out.Write([]byte("`\n"))
	out.Write([]byte(")\n"))
}
