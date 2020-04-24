package proggtk

import (
	"fmt"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
	"github.com/rezasr/filegen-from-excel/internal/prog"
)

func Main() error {
	gtk.Init(nil)

	n, err := prog.Main()

	fmt.Println("Number of generated files: " + strconv.Itoa(n))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
