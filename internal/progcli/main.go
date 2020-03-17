package progcli

import (
	"fmt"
	"strings"

	"github.com/rezasr/filegen-from-excel/internal/prog"
)

func validate() error {
	var errs []string

	err := prog.TemplateFile.Init()
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = prog.OutDir.Init()
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = prog.OutFileName.Init()
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = prog.DataFile.Init()
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	} else {
		return nil
	}
}

func Main() error {
	if err := validate(); err != nil {
		return err
	}

	err := prog.Main()
	if err != nil {
		return err
	}

	return nil
}
