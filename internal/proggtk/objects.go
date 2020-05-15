package proggtk

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/rezasr/filegen-from-excel/internal/prog"
)

var (
	appWindow_Main       *gtk.ApplicationWindow
	fileBtn_DataFile     *gtk.FileChooserButton
	fileBtn_TemplateFile *gtk.FileChooserButton
	fileBtn_OutDir       *gtk.FileChooserButton
	entry_OutFileName    *gtk.Entry
	dialog_Main          *gtk.Dialog
	lbl_DialogMainHeader *gtk.Label
	lbl_DialogMainText   *gtk.Label
	aboutDialog_Main     *gtk.AboutDialog
)

func initObjects() error {
	var ok bool

	obj, err := bldr.GetObject("AppWindow_Main")
	if err != nil {
		return err
	}
	appWindow_Main, ok = obj.(*gtk.ApplicationWindow)
	if !ok {
		return err
	}
	appWindow_Main.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, err = bldr.GetObject("FileBtn_DataFile")
	if err != nil {
		return err
	}
	fileBtn_DataFile, ok = obj.(*gtk.FileChooserButton)
	if !ok {
		return err
	}

	obj, err = bldr.GetObject("FileBtn_TemplateFile")
	if err != nil {
		return err
	}
	fileBtn_TemplateFile, ok = obj.(*gtk.FileChooserButton)
	if !ok {
		return err
	}

	obj, err = bldr.GetObject("FileBtn_OutDir")
	if err != nil {
		return err
	}
	fileBtn_OutDir, ok = obj.(*gtk.FileChooserButton)
	if !ok {
		return err
	}

	obj, err = bldr.GetObject("Entry_OutFileName")
	if err != nil {
		return err
	}
	entry_OutFileName, ok = obj.(*gtk.Entry)
	if !ok {
		return err
	}

	obj, err = bldr.GetObject("Dialog_Main")
	if err != nil {
		return err
	}
	dialog_Main, ok = obj.(*gtk.Dialog)
	if !ok {
		return err
	}
	dialog_Main.HideOnDelete()
	obj, err = bldr.GetObject("Lbl_DialogMainHeader")
	if err != nil {
		return err
	}
	lbl_DialogMainHeader, ok = obj.(*gtk.Label)
	if !ok {
		return err
	}
	obj, err = bldr.GetObject("Lbl_DialogMainText")
	if err != nil {
		return err
	}
	lbl_DialogMainText, ok = obj.(*gtk.Label)
	if !ok {
		return err
	}

	obj, err = bldr.GetObject("AboutDialog_Main")
	if err != nil {
		return err
	}
	aboutDialog_Main, ok = obj.(*gtk.AboutDialog)
	if !ok {
		return err
	}
	aboutDialog_Main.SetVersion("Version: " + prog.VERSION)
	aboutDialog_Main.HideOnDelete()
	aboutDialog_Main.Connect("response", func() {
		aboutDialog_Main.Hide()
	})

	return nil
}
