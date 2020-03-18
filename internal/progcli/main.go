package progcli

import (
	"fmt"

	"github.com/rezasr/filegen-from-excel/internal/prog"
)

func Main() error {
	err := prog.Main()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
