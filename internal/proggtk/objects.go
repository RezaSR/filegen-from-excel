package proggtk

import (
	"github.com/gotk3/gotk3/gtk"
)

var (
	appWindow_Main       *gtk.ApplicationWindow
	fileBtn_DataFile     *gtk.FileChooserButton
	fileBtn_TemplateFile *gtk.FileChooserButton
	fileBtn_OutDir       *gtk.FileChooserButton
	entry_OutFileName    *gtk.Entry
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

	return nil
}
