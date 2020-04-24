package progcli

import (
	"fmt"
	"strconv"

	"github.com/rezasr/filegen-from-excel/internal/prog"
)

func Main() error {
	n, err := prog.Main()

	fmt.Println("Number of generated files: " + strconv.Itoa(n))

	return err
}
