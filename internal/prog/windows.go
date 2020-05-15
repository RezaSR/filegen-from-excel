package prog

import "debug/pe"

const SUBSYSTEM_WINDOWS_GUI = uint16(2)

func isWindowsCli() bool {
	peFile, err := pe.Open(ExePath)

	if err != nil {
		return false
	} else {
		optionalHeader, ok := peFile.OptionalHeader.(*pe.OptionalHeader64)
		if ok {
			return optionalHeader.Subsystem != SUBSYSTEM_WINDOWS_GUI
		} else {
			optionalHeader, ok := peFile.OptionalHeader.(*pe.OptionalHeader32)
			if ok {
				return optionalHeader.Subsystem != SUBSYSTEM_WINDOWS_GUI
			}
		}

		return false
	}
}
