package progcli

import (
	"fmt"
	"strconv"

	"github.com/rezasr/filegen-from-excel/internal/prog"
)

func Main() error {
	n, err := prog.Main()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number of generated files: " + strconv.Itoa(n))
	}

	return nil
}
